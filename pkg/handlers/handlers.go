package main

import (
	"net/http"

	"github.com/xiuxi/go-course/pkg/render"
	// "github.com/xiuxi/go-course/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {

}
