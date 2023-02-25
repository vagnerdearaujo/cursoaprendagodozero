package utils

import (
	"html/template"
	"net/http"
)

// Gerencia o carregamento de templates (páginas html) que serão utilizados na aplicação.
var templates *template.Template

// Carrega todos os arquivos de um determinado tipo como sendo um template,
// neste caso todas as páginas html
func CarregarTemplates() {
	//Carrega todos os .html's da pasta views
	templates = template.Must(templates.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// ExecutarTemplate Carrega um template pelo nome da página e o renderiza na tela
func ExecutarTemplate(w http.ResponseWriter, templatename string, dados interface{}) {
	templates.ExecuteTemplate(w, templatename, dados)
}
