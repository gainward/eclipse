package handlers

import (
	"encoding/json"
	"net/http"

	"appengine"
	"appengine/datastore"

	"backend/models"
)

const (
	gameReqKey = "gameKey"
)

func GameState(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	encoder := json.NewEncoder(w)

	// Handle the case of empty string
	k, err := datastore.DecodeKey(r.FormValue(gameReqKey))
	if err != nil {
		c.Infof("Got key: ...%v...", r.FormValue(gameReqKey))
		encoder.Encode(err)
		c.Infof("Failure to decode request gameKey.")
		c.Infof("Error: %+v", err)
		return
	}

	state, err := models.GetGameState(c, k)
	if err != nil {
		encoder.Encode(err)
		c.Infof("Failure to find game state.")
		return
	}
	encoder.Encode(state)
}
