package main

import (
	"net/http"

	"fmt"
	"github.com/hootsuite/healthchecks"
	"github.com/hootsuite/healthchecks/checks/httpsc"
	"github.com/spf13/viper"
	"strconv"
)

// The StatusHandler handles all Status checking endpoints, particularly those related to healthchecks
type StatusHandler struct {
	healthChecks http.Handler
}

func (h *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.healthChecks.ServeHTTP(w, r)
}

// Initialize the health checking framework
func createHealthChecksHandler(config *viper.Viper) http.Handler {

	// The list of StatusEndpoints for your service
	statusEndpoints := []healthchecks.StatusEndpoint{}

	checks := config.GetStringMap("checks")
	for key := range checks {
		check := config.GetStringMapString(fmt.Sprintf("checks.%s", key))
		// fmt.Println(check)

		isTraversable, _ := strconv.ParseBool(check["istraversable"])
		// fmt.Println(e)
		shouldRandomlyFail, _ := strconv.ParseBool(check["randomlyfail"])
		// fmt.Println(e)

		if isTraversable {
			se := healthchecks.StatusEndpoint{
				Name:          check["name"],
				Slug:          check["slug"],
				Type:          check["type"],
				IsTraversable: isTraversable,
				StatusCheck: httpsc.HttpStatusChecker{
					BaseUrl: check["baseurl"],
				},
				TraverseCheck: httpsc.HttpStatusChecker{
					BaseUrl: check["baseurl"],
				},
			}
			statusEndpoints = append(statusEndpoints, se)
		} else {
			se := healthchecks.StatusEndpoint{
				Name:          check["name"],
				Slug:          check["slug"],
				Type:          check["type"],
				IsTraversable: isTraversable,
				StatusCheck: TestHealthChecker{
					Status: healthchecks.Status{
						Description: check["name"],
						Result:      alertLevelFromString(check["result"]),
						Details:     check["details"],
					},
					ShouldRandomlyFail: shouldRandomlyFail,
				},
				TraverseCheck: nil,
			}
			statusEndpoints = append(statusEndpoints, se)
		}
	}

	aboutFilePath := "conf/about.json"
	versionFilePath := "conf/version.txt"

	// Set up any service injected customData for /Status/about response.
	// Values can be any valid JSON conversion and will override values set in about.json.
	customData := make(map[string]interface{})

	// Register all "/Status/..." requests to use our health checking framework.
	// /Status/am-i-up   - Is the service running?
	// /Status/about     - Describes the service
	// /Status/aggregate - Is the search healthy?
	// /Status/traverse  - Traverse to another level in the service graph
	// /Status/:slug     - Check an individual dependency StatusEndpoint
	return healthchecks.Handler(statusEndpoints, aboutFilePath, versionFilePath, customData)
}

func alertLevelFromString(a string) healthchecks.AlertLevel {
	switch a {
	case "OK":
		return healthchecks.OK
	case "WARN":
		return healthchecks.WARNING
	case "CRIT":
		return healthchecks.CRITICAL
	default:
		return healthchecks.CRITICAL
	}
}
