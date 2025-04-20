package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmp string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmp)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}
