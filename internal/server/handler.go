package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

type Handler struct {
	Logger    *zap.Logger
	Store     Storager
	ApiClient ApiClient
}

type City struct {
	City string
}

func (h *Handler) ChangeLocation(w http.ResponseWriter, r *http.Request) {
	var hand *City

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &hand)
	if err != nil {
		// Error
		return
	}

}
