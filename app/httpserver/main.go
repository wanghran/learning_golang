package main

import (
	"fmt"
	"log"
	"net/http"

	myserver "github.com/wanghran/learning_golang/app/httpserver/server"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hi there, I love %s!", r.URL.Path[1:])
	myserver.PrintForFun()
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
