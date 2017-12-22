package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Action func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) error

// Main controller structure
type Controller struct{}

// Performs custom action
func (c *Controller) Perform(a Action) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if err := a(w, r, ps); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
