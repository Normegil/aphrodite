package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/normegil/aphrodite/router"
	"github.com/sirupsen/logrus"
	"github.com/normegil/aphrodite/handler"
)

const PORT int = 8080

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{}
	port := strconv.Itoa(PORT)
	log.WithField("port", PORT).Info("Server listening")
	err := http.ListenAndServe(":" + port, router.New(handler.Env{log}))
	if nil != err {
		fmt.Print(err)
	}
}
