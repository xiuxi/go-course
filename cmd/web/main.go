package main

import (
	"fmt"
	"net/http"

	"github.com/tsawler/go-course/pkg/handler"
)

const portNum = ":8080"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)
	fmt.Println(fmt.Sprintf("%s", portNum))
	http.ListenAndServe(portNum, nil)
}
