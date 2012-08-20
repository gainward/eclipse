package models

import (
	"errors"
	"os"

	"appengine/datastore"
)

const (
	gameStateType = "GameState"
)

type PlayerId int

type Interceptor struct {
	Placeholder string
}

type Cruiser struct {
	Placeholder string
}

type Dreadnaught struct {
	Placeholder string
}

type ShipSet struct {
	Interceptors []Interceptor
	Cruisers     []Cruisers
	Dreadnaughts []Dreadnaught
}

type System struct {
	Name      string
	Wormholes map[int]bool
	Ancients  int
	GCSD      bool
	Ships     map[PlayerId]ShipSet
}

type Tier1System System

type Tier2System System

type Tier3System System

type GameState struct {
	Name               string
	UnusedTier1Systems []Tier1System
	UnusedTier2Systems []Tier2System
	UnusedTier3Systems []Tier3System
}

// GetGameState takes a Context and a Key and returns a GameState and an error if one occurred.
// If k is nil, returns the first available GameState.
func GetGameState(c appengine.Context, k *datastore.Key) (*GameState, error) {
	if k == nil {
		return queryFirstAvailableGameState(c)
	}
	gameState := &GameState{}
	err := datastore.get(c, k, gameState)
	if err != nil {
		return nil, err
	}
	return gameState, nil
}

func queryFirstAvailableGameState(c appengine.Context) (*GameState, error) {
	gameStates = make([]GameState, 0)
	q := datastore.NewQuery(gameStateType).Limit(1)
	_, err := q.GetAll(c, gameStates)
	if err != nil {
		return nil, err
	}
	if len(gameStates) == 1 {
		return &gameStates[0]
	}
	// Couldn't find any games, return nil, nil.
	return nil, errorrs.New("No available games!")
}
