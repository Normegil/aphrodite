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
			env.Log.WithError(err).Error("Unable to read request")
			return
		}
		user, err := model.NewUser(decoded.Name, decoded.Password)
		if nil != err {
			env.Log.WithError(err).Error("User creation failed")
			return
		}
		err = env.DataSource.UserCreate(*user)
		if nil != err {
			env.Log.WithError(err).Error("User registration failed")
			return
		}
		env.Log.WithField("User", user.Name()).Info("User created")
	}
}
