package main

import (
	"os"

	"github.com/normegil/aphrodite/modules/database"
	"github.com/normegil/aphrodite/modules/environment"
	"github.com/normegil/aphrodite/modules/log"
	"github.com/normegil/aphrodite/router"
	"github.com/sirupsen/logrus"
)

const DB_TYPE = "postgres"
const DB_URL = "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
const PORT = 8080
const LOG_LEVEL = logrus.DebugLevel

func main() {
	dbConn, err := database.NewConnection(DB_TYPE, DB_URL)
	if nil != err {
		panic(err)
	}
	defer dbConn.Close()

	err = router.Router{
		Port: PORT,
		Environment: environment.Environment{
			Log:        log.New(LOG_LEVEL),
			DataSource: dbConn,
		},
	}.Listen()
	if nil != err {
		panic(err)
	}
	os.Exit(0)
}
