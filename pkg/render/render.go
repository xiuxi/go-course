package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmp string) {
	parseTemplate, _ := template.ParseFiles("./templates/" + tmp)
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
