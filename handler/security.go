package handler

import (
	"github.com/normegil/aphrodite/model"
	"net/http"
	"context"
)

func AuthenticationLogger(env model.Env, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := "admin"



		ctx := context.WithValue(r.Context(), "USER", user)
		env.Log.WithField("User", user).Debug("User identified")
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
