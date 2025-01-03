package handlers

import (
	"github.com/go_web/internal/domain"
	"html/template"
	"log"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request, h *Handler) {
	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("./ui/html/signin.html")
		if err != nil {
			log.Println("nothing or damaged html")
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			return
		}

	} else if r.Method == http.MethodPost {

		client := domain.Client{
			Username: r.PostFormValue("username"),
			Password: r.PostFormValue("password"),
		}
		_, err := h.serv.Login(client.Username, client.Password)
		if err == nil {
			http.SetCookie(w, &http.Cookie{
				Name:   "authorized",
				Value:  "true",
				Path:   "/",
				MaxAge: 60,
			})
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		} else if err.Error() == "invalid username or password" {
			http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
		}
	}
}

func SignUp(w http.ResponseWriter, r *http.Request, h *Handler) {
	id := 1
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./ui/html/signup.html")
		if err != nil {
			log.Println("nothing or damaged html")
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
	} else if r.Method == http.MethodPost {
		client := domain.Client{
			ID:       id,
			Name:     r.PostFormValue("name"),
			Username: r.PostFormValue("username"),
			Password: r.PostFormValue("password"),
		}
		id++
		err := h.serv.CreateUser(client)
		if err != nil {
			http.SetCookie(w, &http.Cookie{
				Name:   "authorized",
				Value:  "true",
				Path:   "/",
				MaxAge: 60,
			})
		}
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}
