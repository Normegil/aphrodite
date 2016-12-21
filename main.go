package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/normegil/aphrodite/router"
	"github.com/sirupsen/logrus"
	"github.com/normegil/aphrodite/model"
	"github.com/normegil/aphrodite/db"
)

const PORT int = 8080

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{}
	port := strconv.Itoa(PORT)
	log.WithField("port", PORT).Info("Server listening")
	err := http.ListenAndServe(":" + port, router.New(model.Env{
		Log:log,
		DataSource:db.New(""),
	}))
	if nil != err {
		fmt.Print(err)
	}
}
