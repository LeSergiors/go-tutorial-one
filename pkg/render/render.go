package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders templates using html/templates
func RenderTemplate(w http.ResponseWriter, tmp string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmp)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}
