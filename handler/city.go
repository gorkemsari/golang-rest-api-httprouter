package handler

import (
	"net/http"

	"github.com/gorkemsari/golang-rest-api-httprouter/helper"
	repo "github.com/gorkemsari/golang-rest-api-httprouter/repository"
	"github.com/julienschmidt/httprouter"
)

func CityAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	cities, err := repo.GetAllCities()
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, cities)
}

func City(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("id")
	cities, err := repo.GetCityById(id)
	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, cities)
}