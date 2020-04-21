package server

import (
	"fmt"
	http "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hi there, I love %s!", r.URL.Path[1:])
}
