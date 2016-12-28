package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/normegil/aphrodite/model"
	"github.com/normegil/aphrodite/modules/environment"
	"github.com/normegil/aphrodite/modules/errors"
	"fmt"
)

func UserGet(env environment.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		name := params.ByName("name")
		user, err := env.DataSource.User(name)
		if nil != err {
			errors.Handler{env.Log}.Handle(w, err)
		}
		jsonUser, err := json.Marshal(user)
		if nil != err {
			errors.Handler{env.Log}.Handle(w, err)
		}
		fmt.Fprint(w, string(jsonUser))
	}
}

type userCreateRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func UserCreate(env environment.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var decoded userCreateRequest
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
