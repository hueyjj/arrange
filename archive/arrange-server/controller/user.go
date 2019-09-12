package controller

import "net/http"

func (c *Controller) InitUser() {
	c.Routes.User.HandleFunc("/profile", userProfile).Methods("GET")
}

func userProfile(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hi"))
}
