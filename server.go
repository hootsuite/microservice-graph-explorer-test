package main

import (
	"net/http"
	"time"
)

// NewServer helps build a server with some setup steps and a given router.
// It applies all of the given middleware to the handler; thus they act as "global" middleware
func NewServer(addr string, router http.Handler, middleware []Middleware) *http.Server {
	composedRouter := router
	for _, mdlware := range middleware {
		composedRouter = mdlware(composedRouter)
	}

	return &http.Server{
		Addr:         addr,
		Handler:      composedRouter,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}
}
