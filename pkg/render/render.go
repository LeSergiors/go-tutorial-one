package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders templates using html/templates
func RenderTemplateTest(w http.ResponseWriter, tmp string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmp, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Check to see if we already have the template in out cache
	_, inMap := tc[t]
	if !inMap {
		// Need to create the template
		log.Println("Creating template", t, "...")
		err = createTemplate(t)
		if err != nil {
			log.Println("Error creating template:", err)
		}
	} else {
		// We have the template in the cache
		log.Println("using cached template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

}

func createTemplate(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	tc[t] = tmpl
	return nil
}
