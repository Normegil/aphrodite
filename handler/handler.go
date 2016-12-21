package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"github.com/normegil/aphrodite/model"
	"github.com/satori/go.uuid"
)

func PrintHelloHandler(env model.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		env.Log.WithField("Request", r).Info("Request received")
		idString := params.ByName("id")
		id, err := uuid.FromString(idString)
		if nil != err {
			fmt.Fprint(w, "Cannot parse ID: " + idString)
			return;
		}
		fmt.Fprint(w, env.DataSource.Image(model.ID(id)).Name())
	}
}
