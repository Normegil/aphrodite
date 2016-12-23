package main

import (
	"net/http"
	"strconv"

	"github.com/normegil/aphrodite/router"
	"github.com/sirupsen/logrus"
	"github.com/normegil/aphrodite/model"
	"github.com/normegil/aphrodite/db"
)

const DB_TYPE = "postgres"
const DB_URL = "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
const PORT int = 8080

func main() {
	log := initLogger()

	dbConn, err := db.NewConnection(DB_TYPE, DB_URL)
	if nil != err {
		panic(err)
	}
	defer dbConn.Close()

	env := model.Env{
		Log:log,
		DataSource:dbConn,
	}

	port := strconv.Itoa(PORT)
	env.Log.WithField("port", PORT).Info("Server listening")
	if err = http.ListenAndServe(":" + port, router.New(env)); nil != err {
		panic(err)
	}
}

func initLogger() *logrus.Logger {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{}
	return log
}