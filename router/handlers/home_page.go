package handlers

import (
	"dev-diary/utils"
	"html/template"
	"net/http"
)

func HomePageHandler(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles(utils.GetTemplateFilepath("home.html"), utils.GetTemplateFilepath("base.html")))

	template.ExecuteTemplate(response, "base", nil)
}
