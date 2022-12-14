package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

type handler struct {
	Logger  *zap.Logger
	CityLoc chan<- string
	Service CityExistence
}

type City struct {
	City string
}

func (h *handler) changeLocation(w http.ResponseWriter, r *http.Request) {
	var hand *City

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &hand)
	if err != nil {
		// Error
		return
	}

	err = h.Service.AddNewCity(h.Logger, r.Context(), hand.City)
	if err != nil {
		//
		return
	}

	h.CityLoc <- hand.City
	return
}
