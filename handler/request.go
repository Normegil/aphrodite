package handler

import (
	"net/http"
	"github.com/normegil/aphrodite/model"
)

func RequestLogger(env model.Env, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		env.Log.WithField("Request", r).Debug("Request received")
		h.ServeHTTP(w, r)
	})
}
