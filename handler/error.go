package handler

import (
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/normegil/aphrodite/model"
	"fmt"
	"encoding/json"
)

func Error(log logrus.FieldLogger, e error, w http.ResponseWriter) {
	toSend := model.NewError(e);
	w.WriteHeader(toSend.HTTPStatus())
	jsonErr, err := json.Marshal(toSend)
	if nil != err {
		fmt.Fprint(w, "An Error (" + err.Error() + ") happened when trying to marshall Error to JSON. " + toSend.String())
		log.WithError(err).Error("An error happened while trying to marshall an other error")
		log.WithError(e).Info("Error")
		return
	}
	fmt.Fprint(w, jsonErr)
	log.WithError(e).Info("Error")
}