package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Logger logrus.FieldLogger

type Env struct {
	Logger
}

func PrintHelloHandler(log Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		log.WithField("Request", r).Info("Request received")
		fmt.Fprint(w, "Hello, new World !")
	}
}
