package main

import (
	"github.com/normegil/aphrodite/router"
	"github.com/sirupsen/logrus"
	"github.com/normegil/aphrodite/db"
	"os"
)

const DB_TYPE = "postgres"
const DB_URL = "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
const PORT int = 8080

func main() {
	dbConn, err := db.NewConnection(DB_TYPE, DB_URL)
	if nil != err {
		panic(err)
	}
	defer dbConn.Close()

	err = router.Router{
		Port: PORT,
		LoggingLevel: logrus.DebugLevel,
		DataSource: dbConn,
	}.Listen()
	if nil != err {
		panic(err)
	}
	os.Exit(0)
}