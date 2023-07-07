package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/model"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/service"
)

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	player, err := getPlayer(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = service.CreatePlayer(player); err != nil {
		log.Printf("Error calling service CreatePlayer: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, player)
}

func GetPlayers(w http.ResponseWriter, _ *http.Request) {
	players, err := service.GetPlayers()
	if err != nil {
		log.Printf("Error calling service GetPlayers: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, players)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	player, err := service.GetPlayer(id)
	if err != nil {
		log.Errorf("Failure retrieving player with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if player == nil {
		http.Error(w, "404 player not found", http.StatusNotFound)
		return
	}
	sendJson(w, player)
}

func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	player, err := getPlayer(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	player, err = service.UpdatePlayer(id, player)
	if err != nil {
		log.Errorf("Failure updating player with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if player == nil {
		http.Error(w, "404 player not found", http.StatusNotFound)
		return
	}
	sendJson(w, player)
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	player, err := service.DeletePlayer(id)
	if err != nil {
		log.Errorf("Failure deleting player with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if player == nil {
		http.Error(w, "404 player not found", http.StatusNotFound)
		return
	}
	sendJson(w, result{Success: "OK"})
}

func getPlayer(r *http.Request) (*model.Player, error) {
	var player model.Player
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		log.Errorf("Can't serialize request body to player struct: %v", err)
		return nil, err
	}
	return &player, nil
}
