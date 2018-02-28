package main

import (
	"encoding/json"
	"github.com/spf13/viper"
	"net/http"
)

// IndexHandler handles the index page with serves generic version information
// It also handles all other unhandled endpoints (eg 404)
type IndexHandler struct {
	config *viper.Viper
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Here you can handle subpaths of anything that comes to the index handler
	// If this handler is bound to the root route ("/"), it's special, as it will also handle all other unhandled routes

	switch r.URL.Path {
	case "/":
		h.index(w, r)
	default:
		// Handle any route that isn't found
		// https://golang.org/pkg/net/http/#ServeMux
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *IndexHandler) index(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Version string `json:"version"`
		Name    string `json:"name"`
	}
	resp := Response{
		Version: "0.0.1",
		Name:    h.config.GetString("serviceName"),
	}
	json.NewEncoder(w).Encode(resp)
}
