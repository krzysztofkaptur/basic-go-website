package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tc = map[string]*template.Template{}

func RenderTemplate(w http.ResponseWriter, tname string) {
	// create a template cache

	// get requested template from cache

	// render the template

	_, ok := tc[tname]

	if !ok {
		err := createCacheTemplate(tname)
		if err != nil {
			log.Println("Error:", err)
		}
	}

	tc[tname].Execute(w, "")
}

func createCacheTemplate(tname string) error {
	templates := []string{
		fmt.Sprintf("./templates/%v.page.tmpl", tname),
		"./layouts/base.layout.tmpl",
	}

	temp, err := template.ParseFiles(templates...)
	if err != nil {
		log.Println("Error:", err)
	}

	tc[tname] = temp

	return nil
}
