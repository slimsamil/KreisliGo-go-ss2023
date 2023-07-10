package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/model"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/service"
)

func CreateGame(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	game, err := getGame(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = service.CreateGame(id, game); err != nil {
		log.Printf("Error calling service CreateGame: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, game)
}

func GetGames(w http.ResponseWriter, _ *http.Request) {
	games, err := service.GetGames()
	if err != nil {
		log.Printf("Error calling service GetGames: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, games)
}

func GetGame(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	game, err := service.GetGame(id)
	if err != nil {
		log.Errorf("Failure retrieving game with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if game == nil {
		http.Error(w, "404 game not found", http.StatusNotFound)
		return
	}
	sendJson(w, game)
}

func UpdateGame(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	game, err := getGame(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	game, err = service.UpdateGame(id, game)
	if err != nil {
		log.Errorf("Failure updating game with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if game == nil {
		http.Error(w, "404 game not found", http.StatusNotFound)
		return
	}
	sendJson(w, game)
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	game, err := service.DeleteGame(id)
	if err != nil {
		log.Errorf("Failure deleting game with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if game == nil {
		http.Error(w, "404 game not found", http.StatusNotFound)
		return
	}
	sendJson(w, result{Success: "OK"})
}

func getGame(r *http.Request) (*model.Game, error) {
	var game model.Game
	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		log.Errorf("Can't serialize request body to game struct: %v", err)
		return nil, err
	}
	return &game, nil
}
