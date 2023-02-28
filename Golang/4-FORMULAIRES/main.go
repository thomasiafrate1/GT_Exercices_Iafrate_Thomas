package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	tmpl1 := template.Must(template.ParseFiles("formulaire.html"))
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl1.Execute(w, nil)
		nom := r.FormValue("nom")
		prénom := r.FormValue("prénom")
		email := r.FormValue("email")
		message := r.FormValue("message")
		age := r.FormValue("age")
		ville := r.FormValue("ville")
		fmt.Printf("Nom: %s\nPrénom: %s\nEmail: %s\nMessage: %s\nAge: %s\nVille: %s\n", nom, prénom, email, message, age, ville)
	})
	http.ListenAndServe(":8080", nil)
}
