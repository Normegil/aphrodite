package router

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/normegil/aphrodite/handler"
	"github.com/normegil/aphrodite/log"
	"github.com/normegil/aphrodite/model"
	"github.com/sirupsen/logrus"
)

type Router struct {
	Port         int
	LoggingLevel logrus.Level
	DataSource   model.DataSource
	router       *httprouter.Router
}

func (r Router) Listen() error {
	env := model.Env{
		Log:        log.New(r.LoggingLevel),
		DataSource: r.DataSource,
	}

	r.router = httprouter.New()
	err := r.registerRoutes(env)
	if nil != err {
		return err
	}

	h := handler.AuthenticationLogger(env, handler.RequestLogger(env.Log, r.router))

	port := strconv.Itoa(r.Port)
	env.Log.WithField("port", port).Info("Server listening")
	if err = http.ListenAndServe(":"+port, h); nil != err {
		return err
	}
	return nil
}

func (r Router) registerRoutes(env model.Env) error {
	if nil == r.router {
		return errors.New("You forgot to initialize router before registering routes")
	}
	r.router.GET("/image", handler.ImageGetAll(env))
	r.router.PUT("/user", handler.UserCreate(env))
	return nil
}
