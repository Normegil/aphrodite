package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/normegil/aphrodite/handler"
	"github.com/normegil/aphrodite/model"
)

func New(env model.Env) http.Handler {
	router := httprouter.New()
	router.GET("/image/:id", handler.PrintHelloHandler(env))
	return router
}
