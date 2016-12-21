package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const PORT int = 8080

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World !")
	})

	err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
	if nil != err {
		fmt.Print(err)
	}
}
