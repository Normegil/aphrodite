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
				env.Log.WithError(err).Error("Unable to authenticate user")
				return
			}
			env.Log.WithError(err).Error("Authentication failed")
		}

		ctx := context.WithValue(r.Context(), "USER", user)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
