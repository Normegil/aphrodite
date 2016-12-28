package router

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/normegil/aphrodite/modules/environment"
	"github.com/normegil/aphrodite/modules/middlewares"
	"github.com/normegil/aphrodite/router/api"
)

type Router struct {
	Port        int
	Environment environment.Environment
	router      *httprouter.Router
}

func (r Router) Listen() error {
	r.router = httprouter.New()
	err := r.registerRoutes(r.Environment)
	if nil != err {
		return err
	}

	router := middlewares.Authentication(r.Environment, middlewares.RequestLogger(r.Environment.Log, r.router))

	port := strconv.Itoa(r.Port)
	r.Environment.Log.WithField("port", port).Info("Server listening")
	if err = http.ListenAndServe(":" + port, router); nil != err {
		return err
	}
	return nil
}

func (r Router) registerRoutes(env environment.Environment) error {
	if nil == r.router {
		return errors.New("You forgot to initialize router before registering routes")
	}
	r.router.GET("/rest/image", api.ImageGetAll(env))

	r.router.GET("/rest/user/:name", api.UserGet(env))
	r.router.PUT("/rest/user", api.UserCreate(env))
	return nil
}
