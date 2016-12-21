package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/normegil/aphrodite/handler"
)

func New(env handler.Env) http.Handler {
	router := httprouter.New()
	router.GET("/", handler.PrintHelloHandler(env))
	return router
}
