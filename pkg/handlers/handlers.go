package handlers

import (
	"net/http"

	render "github.com/krzysztofkaptur/basic-go-website/pkg/render"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home")

}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about")
}
