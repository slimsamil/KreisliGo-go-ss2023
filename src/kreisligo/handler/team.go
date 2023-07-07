package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/model"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/service"
)

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	team, err := getTeam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = service.CreateTeam(team); err != nil {
		log.Printf("Error calling service CreateTeam: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, team)
}

func GetTeams(w http.ResponseWriter, _ *http.Request) {
	teams, err := service.GetTeams()
	if err != nil {
		log.Printf("Error calling service GetTeams: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, teams)
}

func GetTeam(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	team, err := service.GetTeam(id)
	if err != nil {
		log.Errorf("Failure retrieving team with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if team == nil {
		http.Error(w, "404 team not found", http.StatusNotFound)
		return
	}
	sendJson(w, team)
}

func UpdateTeam(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	team, err := getTeam(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	team, err = service.UpdateTeam(id, team)
	if err != nil {
		log.Errorf("Failure updating team with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if team == nil {
		http.Error(w, "404 team not found", http.StatusNotFound)
		return
	}
	sendJson(w, team)
}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	team, err := service.DeleteTeam(id)
	if err != nil {
		log.Errorf("Failure deleting team with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if team == nil {
		http.Error(w, "404 team not found", http.StatusNotFound)
		return
	}
	sendJson(w, result{Success: "OK"})
}

func getTeam(r *http.Request) (*model.Team, error) {
	var team model.Team
	err := json.NewDecoder(r.Body).Decode(&team)
	if err != nil {
		log.Errorf("Can't serialize request body to team struct: %v", err)
		return nil, err
	}
	return &team, nil
}
