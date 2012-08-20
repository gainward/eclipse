package backend

import (
	"http"

	"backend/hanlders"
)

func init() {
	http.HandleFunc("/", handlers.GameState)
}
