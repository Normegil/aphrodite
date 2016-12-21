package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
)

const PORT int = 8080

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hello, new World !")
	})

	err := http.ListenAndServe(":"+strconv.Itoa(PORT), router)
	if nil != err {
		fmt.Print(err)
	}
}
