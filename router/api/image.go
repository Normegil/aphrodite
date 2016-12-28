package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/normegil/aphrodite/modules/environment"
	"github.com/normegil/aphrodite/modules/errors"
)

func ImageGetAll(env environment.Environment) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		images := env.DataSource.AllImages(0, 20)
		jsonImages, err := json.Marshal(images)
		if nil != err {
			errors.Handler{env.Log}.Handle(w, err)
		}
		fmt.Fprint(w, string(jsonImages))
	}
}
