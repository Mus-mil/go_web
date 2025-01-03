package handlers

import (
	"github.com/go_web/internal/service"
	"net/http"
)

type dataHTML struct {
	IsAuthorized bool
}

type Handler struct {
	serv *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{serv: serv}
}

func RegisterHandlers(h *Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", WelcomeHandle)
	mux.HandleFunc("/signin", withH(SignIn, h))
	mux.HandleFunc("/signup", withH(SignUp, h))
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	return mux
}

func withH(handler func(w http.ResponseWriter, r *http.Request, h *Handler), h *Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, h)
	}
}
