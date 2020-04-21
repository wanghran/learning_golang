package httpserver

import (
	"log"
	"net/http"

	myserver "github.com/wanghran/learning_golang/app/httpserver/server"
)

func main() {
	http.HandleFunc("/", myserver.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
