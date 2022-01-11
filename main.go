package main

import (
	"fmt"
	// "log"
	"net/http"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Inicio)
	fmt.Println("Servidor andando..")
	// log.Println("Servidor volando....")
	http.ListenAndServe(":8080", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request) {

	plantillas.ExecuteTemplate(w, "inicio", nil)

}

// voy 23:44
