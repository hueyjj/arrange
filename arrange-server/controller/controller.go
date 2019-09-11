package controller

import (
	"github.com/gorilla/mux"
)

type Controller struct {
	Routes *Routes
}

type Routes struct {
	Root *mux.Router // "/"
	Api  *mux.Router // "api/v1"

	User *mux.Router // "api/v1/user"
}

func Init(root *mux.Router) *Controller {
	controller := &Controller{
		Routes: &Routes{},
	}

	controller.Routes.Root = root
	controller.Routes.Api = controller.Routes.Root.PathPrefix("/api/v1").Subrouter()

	controller.Routes.User = controller.Routes.Api.PathPrefix("/user").Subrouter()

	controller.InitUser()

	return controller
}
