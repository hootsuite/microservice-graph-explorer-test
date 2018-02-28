package main

import (
	"net/http"

	"github.com/go-errors/errors"
	"github.com/sirupsen/logrus"
)

// A Middleware is just a function that wraps a handler with additional features
type Middleware func(http.Handler) http.Handler

// PanicLogging is a middleware that intercepts panics and logs them using our built-in logger so they
// show up in Sumologic
func PanicLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				goErr := errors.Wrap(err, 3)
				log.WithFields(logrus.Fields{
					"err":        goErr.Err,
					"stackTrace": string(goErr.Stack()),
				}).Errorf("Panic recovery -> %s\n", goErr.Err)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
