package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/normegil/aphrodite/model"
	"github.com/normegil/aphrodite/modules/environment"
	"github.com/normegil/aphrodite/modules/errors"
)

type UserCreateRequest struct {
	Name     string
	Password string
}

func UserCreate(env environment.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var decoded UserCreateRequest
		err := json.NewDecoder(r.Body).Decode(&decoded)
		if nil != err {
			errors.Handler{env.Log}.Handle(w, errors.NewErrWithCode(40001, err))
			return
		}
		user, err := model.NewUser(decoded.Name, decoded.Password)
		if nil != err {
			errors.Handler{env.Log}.Handle(w, errors.NewErrWithCode(40001, err))
			return
		}
		err = env.DataSource.UserCreate(*user)
		if nil != err {
			errors.Handler{env.Log}.Handle(w, errors.NewErrWithCode(40001, err))
			return
		}
		env.Log.WithField("User", user.Name()).Info("User created")
	}
}
