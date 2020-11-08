package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates : Load Templates from the templates folder
func LoadTemplates(folder string) {
	templates = template.Must(template.ParseGlob(folder))
	print("Loaded templates.\n")
}

// LoadStaticFiles : Load Static Files from the static folder
func LoadStaticFiles(folder string) {
	fs := http.FileServer(http.Dir(folder))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	print("Loaded static files.\n")
}

// RenderTemplate : Render Templates by demand
func RenderTemplate(res http.ResponseWriter, templateName string, data interface{}) {
	templates.ExecuteTemplate(res, templateName, data)
}
