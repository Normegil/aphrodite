package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/normegil/aphrodite/router"
	"github.com/Normegil/log"
)

const PORT int = 8080

func main() {
	log := log.StructuredLog{

	}
	itoa := strconv.Itoa(PORT)
	log.Log(log.DEBUG, log.Structure{}, "Listening on " + itoa)
	err := http.ListenAndServe(":"+ itoa, router.New())
	if nil != err {
		fmt.Print(err)
	}
}
