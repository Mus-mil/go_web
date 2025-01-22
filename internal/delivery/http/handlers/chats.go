package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func WelcomeHandle(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("authorized")
	isAuthorized := err == nil

	tmpl, err := template.ParseFiles("./ui/html/welcome.html")
	if err != nil {
		log.Println("nothing or damaged html")
	}
	data := dataHTML{
		IsAuthorized: isAuthorized,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		return
	}
}
