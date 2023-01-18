package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

/*
templates irá receber as páginas em html do projeto.
pode-se declarar templates dentro de uma função, porém,	para templates que são utilizados
ao longo de todo o pacote, a declaração deve ocorrer fora de uma função.
*/
var templates *template.Template

/*
Os campos da struct obrigatoriamente devem começar com uma letra maíuscula
*/
type usuario struct {
	Nome  string
	Email string
	Agora string
}

func main() {
	/*Carrega todas as páginas que atendam ao critério *.html
	Para projetos maiores que envolvam diversas áreas, o preferível seria ter diversos templates
	cada um carregando um padrão específico de uma determinada área, ou seja: fin_*.html, rh_*.html e etc"
	*/
	templates = template.Must(template.ParseGlob("*.html"))
	http.HandleFunc("/", defaultpage)
	http.HandleFunc("/usuario", usuarios)
	//http.NotFound(http.NotFoundHandler(),pageNotFound)

	//Carrega o servidor de páginas e executa até ser interrompido
	println("Carregando o servidor de páginas na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func defaultpage(w http.ResponseWriter, r *http.Request) {
	/*
		O templates possui uma coleção de templates (páginas) que foram carregados.
		o método executeTempate possui três parâmetros.
		Response Writer, que neste caso é o próprio w que foi recebido como parâmetro
		O nome do arquivo, para este caso home.html
		E o pacote de dados que será fundido ao template para mostrar uma página dinânimca
	*/
	templates.ExecuteTemplate(w, "home.html", nil)
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	agora := fmt.Sprintf("%2d:%2d:%2d - %2d/%2d/%4d", time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Day(), time.Now().Month(), time.Now().Year())
	time.Now().Day()
	dados_usuario := usuario{
		"Vagner de Araujo",
		"vagnerdearaujo@gmail.com",
		agora,
	}
	templates.ExecuteTemplate(w, "dados_usuario.html", dados_usuario)
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deu ruim, página não encontrada !!!!"))
}
