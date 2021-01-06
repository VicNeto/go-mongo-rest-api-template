package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-mongodb/models"
	. "go-rest-mongodb/repository"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var placesRepository PlacesRepository

// GetAllPlaces godoc
// @Summary Places
// @Description Places list
// @Tags places
// @Produce json
// @Success 200 {array} models.Place
// @Router /api/places [get]
func GetAllPlaces(w http.ResponseWriter, r *http.Request) {
	places, err := placesRepository.FindAll()
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Info(places)
	jsonResponse(w, http.StatusOK, places)
}

// GetPlaceById godoc
// @Summary Place detail
// @Description Place detail
// @Tags places
// @Accept  json
// @Produce json
// @Success 200 {object} models.Place
// @Router /api/places/{id} [get]
func GetPlaceById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented!")
}

// CreatePlace godoc
// @Summary post request example
// @Description post request example
// @Tags places
// @Accept json
// @Produce json
// @Param place body models.Place true "Add place"
// @Success 200 {string} string "success"
// @Failure 500 {string} string "fail"
// @Router /api/places [post]
func CreatePlace(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var place models.Place
	if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
		errorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	log.Info(place)
	insertResult, err := placesRepository.Insert(place)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	place.ID = insertResult.(primitive.ObjectID)
	jsonResponse(w, http.StatusCreated, place)
}

func UpdatePlace(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented")
}

func DeletePlace(w http.ResponseWriter, r *http.Request) {
	//Vars returns the route variables for the current request, if any.
	vars := mux.Vars(r)
	//Get id from the current request
	id := vars["id"]
	fmt.Println(id)
	if err := placesRepository.Delete(id); err != nil {
		errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func errorResponse(w http.ResponseWriter, code int, msg string) {
	jsonResponse(w, code, map[string]string{"error": msg})
}

func jsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
