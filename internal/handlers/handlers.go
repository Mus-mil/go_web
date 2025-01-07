package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go_web/internal/service"
	"net/http"
	"time"
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

func RegisterRoutes(h *Handler) *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("ui/html/*")
	router.StaticFS("/static", http.Dir("ui/static"))

	auth := router.Group("/auth")
	{
		auth.GET("/signin", h.SignInGet)
		auth.POST("/signin", h.SignInPost)
		auth.GET("/signup", h.SignUpGet)
		auth.POST("/signup", h.SignUpPost)
	}
	router.GET("/id", h.idGet)
	router.GET("/", h.welcome)
	return router
}

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}
