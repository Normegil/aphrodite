package middlewares

import (
	"context"
	"net/http"

	"github.com/normegil/aphrodite/modules/environment"
	"github.com/normegil/aphrodite/modules/errors"
	"github.com/normegil/aphrodite/modules/security"
)

func Authentication(env environment.Environment, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authenticator := &security.Authenticator{env.DataSource}
		user, err := authenticator.Authenticate(r)
		if nil != err {
			if err != security.AuthenticationFailed {
				errors.Handler{env.Log}.Handle(w, err)
				return
			}
			env.Log.WithError(err).Error("Authentication failed")
		}

		ctx := context.WithValue(r.Context(), "User", user)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
