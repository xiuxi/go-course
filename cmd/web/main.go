package main

import (
	"fmt"
	"net/http"

	"github.com/xiuxi/go-course/pkg/handlers"
)

const portNum = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Sprintf("%s", portNum))
	http.ListenAndServe(portNum, nil)
}
