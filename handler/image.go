package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/normegil/aphrodite/model"
)

func ImageGetAll(env model.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		images := env.DataSource.AllImages(0, 20)
		jsonImages, err := json.Marshal(images)
		if nil != err {
			Error(env.Log, err, w)
		}
		fmt.Fprint(w, string(jsonImages))
	}
}
