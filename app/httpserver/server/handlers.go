package server

import (
	"fmt"
	"net/http"
)

func HandlerPrint(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hi there, I love %s!", r.URL.Path[1:])
}
