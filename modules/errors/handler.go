package errors

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	Log logrus.FieldLogger
}

func (h Handler) Handle(w http.ResponseWriter, e error) {
	responseBody := newResponse(e)
	w.WriteHeader(responseBody.HTTPStatus)
	responseBodyJSON, err := json.Marshal(responseBody)
	if nil != err {
		fmt.Fprint(w, "An Error ("+err.Error()+") happened when trying to marshall Error to JSON. "+responseBody.String())
		h.Log.WithError(err).Error("An error happened while trying to marshall an other error")
		h.Log.WithError(e).Info("Error")
		return
	}
	fmt.Fprint(w, responseBodyJSON)
	h.Log.WithError(e).Info("Error")
}
