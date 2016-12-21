package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func PrintHello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello, new World !")
}
