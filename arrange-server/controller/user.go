package controller

import "net/http"

func (c *Controller) InitUser() {
	c.Routes.User.HandleFunc("/register", registerUser).Methods("POST")
}

func registerUser(w http.ResponseWriter, r *http.Request) {
}
