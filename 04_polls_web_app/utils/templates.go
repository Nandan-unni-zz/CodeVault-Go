package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates : Load Templates from the templates folder
func LoadTemplates(pattern string) {
	templates = template.Must(template.ParseGlob(pattern))
}

// RenderTemplate : Render Templates by demand
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}
