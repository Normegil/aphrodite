package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/normegil/aphrodite/model"
	"encoding/json"
	"fmt"
)

func ImageGetAll(env model.Env) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		images := env.DataSource.AllImages(0, 20)
		jsonImages, err := json.Marshal(images)
		if nil != err {
			fmt.Fprint(w, "ERROR", err)
		}
		fmt.Fprint(w, string(jsonImages))
	}
}
