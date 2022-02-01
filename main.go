package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)
type Posts []Post
var templates  *template.Template
func main(){ 
	templates = template.Must(template.ParseGlob("templates/*.html"))
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	http.Handle("/", router)
	http.ListenAndServe(":8081", nil)

}
func indexHandler(res http.ResponseWriter, req *http.Request){
	templates.ExecuteTemplate(res, "index.html", nil)

}
type Post struct {
	Title string `json:title`
	Body string `json:body`
}
