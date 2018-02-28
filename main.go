package main

import (
	"net/http"

	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var log *logrus.Logger

func main() {
	config, err := readInConfig("local", "./conf")
	if err != nil {
		panic(err)
	}

	log, err = createLogger()
	if err != nil {
		log.Errorln(err)
	}

	router := createRouter(config)
	server := NewServer(":"+config.GetString("server.port"), router, []Middleware{
		PanicLogging,
	})

	log.Infof("Your service is listening on port %s", config.GetString("server.port"))

	panic(server.ListenAndServe())
}

func readInConfig(env string, paths ...string) (*viper.Viper, error) {
	if env == "" {
		env = "local"
	}

	config := viper.New()
	for i := range paths {
		config.AddConfigPath(paths[i])
	}

	config.SetConfigName(env)
	if err := config.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("configuration error: %s", err)
	}

	return config, nil
}

func createLogger() (*logrus.Logger, error) {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	return logger, nil
}

// Creates the standard root logger that handles traffic to the server
func createRouter(config *viper.Viper) http.Handler {
	router := http.NewServeMux()

	router.Handle("/", &IndexHandler{config: config})

	router.Handle("/status/", &StatusHandler{
		healthChecks: createHealthChecksHandler(config),
	})

	return router
}
