package main

import (
	"log"
	"net/http"

	myserver "github.com/wanghran/learning_golang/app/httpserver/server"
)

func main() {
	http.HandleFunc("/", myserver.HandlerPrint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
