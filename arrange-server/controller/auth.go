package controller

import (
	"net/http"

	"github.com/hueyjj/arrange-server/model"
)

func (c *Controller) InitAuth() {
	c.Routes.Api.HandleFunc("/register", registerUser).Methods("POST")
	c.Routes.Api.HandleFunc("/login", loginUser).Methods("POST")
	c.Routes.Api.HandleFunc("/logout", logoutUser).Methods("GET")
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	registration := model.RegistrationFromJson(r.Body)
	w.WriteHeader(200)
	w.Write([]byte("hi"))
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hi"))
}

func logoutUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hi"))
}
