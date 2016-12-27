package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/normegil/aphrodite/model"
)

type UserCreateRequest struct {
	Name     string
	Password string
}

func UserCreate(env model.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var decoded UserCreateRequest
		err := json.NewDecoder(r.Body).Decode(&decoded)
		if nil != err {
			Error(env.Log, model.NewErrWithCode(40001, err), w)
			return
		}
		user, err := model.NewUser(decoded.Name, decoded.Password)
		if nil != err {
			Error(env.Log, model.NewErrWithCode(40001, err), w)
			return
		}
		err = env.DataSource.UserCreate(*user)
		if nil != err {
			Error(env.Log, model.NewErrWithCode(40001, err), w)
			return
		}
		env.Log.WithField("User", user.Name()).Info("User created")
	}
}
