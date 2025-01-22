package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go_web/internal/models"
	"net/http"
)

func (h *Handler) SignInGet(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", gin.H{
		"error": "",
	})
}

func (h *Handler) SignInPost(c *gin.Context) {
	_, err := h.serv.LoginUser(c.PostForm("username"), c.PostForm("password"))
	if err != nil {
		c.HTML(http.StatusOK, "signin.html", gin.H{"error": "неправильный пароль или логин"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/id")
}

func (h *Handler) SignUpGet(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

func (h *Handler) SignUpPost(c *gin.Context) {
	client := models.User{
		Name:     c.PostForm("name"),
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
	}
	err := h.serv.CreateUser(client)
	if err != nil {
		c.HTML(http.StatusOK, "signup.html", gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/id")
}

func (h *Handler) idGet(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"IsAuthorized": true,
	})
}

func (h *Handler) welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"IsAuthorized": false,
	})
}
