package main

import (
	"fmt"
	"log"
	"io"
	"net/http"
	"github.com/gorilla/mux"
)
type Posts []Post

func main(){ 
	router := mux.NewRouter()

}

type Post struct {
	Title string `json:title`
	Body string `json:body`
}
