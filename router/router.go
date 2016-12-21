package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/normegil/aphrodite/handler"
)

func New() http.Handler {
	router := httprouter.New()
	router.GET("/", handler.PrintHello)
	return router
}
