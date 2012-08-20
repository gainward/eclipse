package backend

import (
	"net/http"

	"backend/handlers"
)

func init() {
	http.HandleFunc("/s/getGameState", handlers.GameState)
}
