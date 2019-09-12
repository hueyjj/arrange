package web

import (
	"net/http"

	"github.com/hueyjj/arrange-server/app"
)

type Handler struct {
	HandleFunc func(*Context, http.ResponseWriter, *http.Request)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext()
	c.App = app.New()
	h.HandleFunc(c, w, r)
}

func AuthRequired(h func(*Context, http.ResponseWriter, *http.Request)) http.Handler {
	return Handler{
		HandleFunc: h,
	}
}
