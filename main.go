package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmp string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmp)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")

}

func addValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	fmt.Printf("Starting application on port %s \n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
