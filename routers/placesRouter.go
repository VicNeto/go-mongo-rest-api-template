package routers

import (
	"go-rest-mongodb/controllers"

	"github.com/gorilla/mux"
)

// AddPlacesRouter godoc
func AddPlacesRouter(r *mux.Router) *mux.Router {
	s := r.PathPrefix("/places").Subrouter()
	s.HandleFunc("", controllers.GetAllPlaces).Methods("GET")
	s.HandleFunc("/{id}", controllers.GetPlaceById).Methods("GET")
	s.HandleFunc("", controllers.CreatePlace).Methods("POST")
	s.HandleFunc("", controllers.UpdatePlace).Methods("PUT")
	s.HandleFunc("/{id}", controllers.DeletePlace).Methods("DELETE")
	return s
}
