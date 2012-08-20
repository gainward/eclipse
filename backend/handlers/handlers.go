package handlers

import (
	"encoding/json"
	"http"

	"appengine"
	"appengine/datastore"

	"models"
)

const (
	gameReqKey = "gameKey"
)

func GameState(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	encoder := json.NewEncoder(w)

	k, err := datastore.DecodeKey(r.FormValue(gameReqKey))
	if err != nil {
		encoder.Encode(err)
		return
	}

	state, err := models.GetGameState(c, k)
	if err != nil {
		encoder.Encode(err)
		return
	}
	encoder.Encode(state)
}
