package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/normegil/aphrodite/router"
)

const PORT int = 8080

func main() {
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), router.New())
	if nil != err {
		fmt.Print(err)
	}
}
