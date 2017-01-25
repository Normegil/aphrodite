package middlewares

import (
	"net/http"

	"github.com/Sirupsen/logrus"
)

func RequestLogger(log logrus.FieldLogger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithField("Request", r).Debug("Request received")
		h.ServeHTTP(w, r)
	})
}
