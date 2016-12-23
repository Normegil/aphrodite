package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/normegil/aphrodite/handler"
	"github.com/normegil/aphrodite/model"
	"github.com/sirupsen/logrus"
	"github.com/normegil/aphrodite/log"
	"strconv"
	"errors"
)

type Router struct {
	Port int
	LoggingLevel logrus.Level
	DataSource model.DataSource
	router *httprouter.Router
}

func (r Router) Listen() error {
	env := model.Env{
		Log: log.New(r.LoggingLevel),
		DataSource: r.DataSource,
	}

	r.router = httprouter.New()
	err := r.registerRoutes(env)
	if nil != err {
		return err
	}

	port := strconv.Itoa(r.Port)
	env.Log.WithField("port", port).Info("Server listening")
	if err = http.ListenAndServe(":" + port, r.router); nil != err {
		return err
	}
	return nil
}

func (r Router) registerRoutes(env model.Env) error {
	if nil == r.router {
		return errors.New("You forgot to initialize router before registering routes")
	}
	r.router.GET("/image", handler.ImageGetAll(env))
	return nil
}