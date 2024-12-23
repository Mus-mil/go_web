package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type dataHTML struct {
	IsAuthorized bool
}

type Client struct {
	ID       int
	Name     string
	Username string
	Password string
}

func (c Client) NewUser() {
	db := initDB()
	err := editUsersTable(db, c)
	if err != nil {
		return
	}
	closeDB(db)
}

func (c Client) GetUserLogin(username string) Client {
	db := initDB()
	client, err := getUserFromUsername(db, username)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func ServerRun() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", WelcomeHandle)
	mux.HandleFunc("/signin", SignInHandle)
	mux.HandleFunc("/signup", SignUpHandle)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	log.Println("Starting server on port http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}

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

func SignInHandle(w http.ResponseWriter, r *http.Request) {
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

		client := Client{
			Username: r.PostFormValue("username"),
			Password: r.PostFormValue("password"),
		}
		clientLogin := Client.GetUserLogin(client.Username)
		if client.Username == "admin" && client.Password == "admin" {
			http.SetCookie(w, &http.Cookie{
				Name:   "authorized",
				Value:  "true",
				Path:   "/",
				MaxAge: 60,
			})
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}
	}
}

func SignUpHandle(w http.ResponseWriter, r *http.Request) {
	id := 1
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./ui/html/signup.html")
		if err != nil {
			log.Println("nothing or damaged html")
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			return
		}
	} else if r.Method == http.MethodPost {
		client := Client{
			ID:       id,
			Name:     r.PostFormValue("name"),
			Username: r.PostFormValue("username"),
			Password: r.PostFormValue("password"),
		}
		id++
		client.NewUser()
		http.SetCookie(w, &http.Cookie{
			Name:   "authorized",
			Value:  "true",
			Path:   "/",
			MaxAge: 60,
		})
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}
