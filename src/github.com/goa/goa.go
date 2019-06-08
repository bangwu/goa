package goa

import (
	"net/http"
)

type Goa struct {
	app *http.ServeMux
}

func (goa Goa) Route(routeURL string, handler func(w http.ResponseWriter, r *http.Request)) {
	goa.app.Handle(routeURL, http.HandlerFunc(handler))
}

func (goa Goa) ListenAndServe(port string) error {
	return http.ListenAndServe(port, goa.app)
}

func NewGoa() *Goa {
	goa := new(Goa)
	goa.app = http.NewServeMux()
	return goa
}
