package handlers

import (
	"net/http"
	"rmccallister/pkg/render"
)

// home page handler

func Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "templates/home.html")

}

//about page handler

func about(w http.ResponseWriter, r *http.Request) {

}
