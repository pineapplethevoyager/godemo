package controllers

import (
	"net/http"
	"regexp"
)
type userController struct {
	userIDPattern *regexp.Regexp //not traditional class
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) { //go method
	w.Write([]byte("Hello from the User Controller!"))
}

func newUserController() *userController{
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}