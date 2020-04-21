package main

import (
	"fmt"

	page "github.com/wanghran/learning_golang/library/wikipage"
)

func main() {
	p := page.Page{Title: "save", Body: []byte("save")}
	fmt.Println(p.Body)
}
