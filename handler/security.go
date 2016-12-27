package handler

import (
	"context"
	"net/http"

	"github.com/normegil/aphrodite/model"
	"github.com/normegil/aphrodite/security"
)

func AuthenticationLogger(env model.Env, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authenticator := &security.Authenticator{env.DataSource}
		user, err := authenticator.Authenticate(r)
		if nil != err {
			if err != security.AuthenticationFailed {
				Error(env.Log, err, w)
				return
			}
			env.Log.WithError(err).Error("Authentication failed")
		}

		ctx := context.WithValue(r.Context(), "User", user)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
