package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/model"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/service"
)

func CreateLeague(w http.ResponseWriter, r *http.Request) {
	league, err := getLeague(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = service.CreateLeague(league); err != nil {
		log.Printf("Error calling service CreateLeague: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, league)
}

func GetLeagues(w http.ResponseWriter, _ *http.Request) {
	leagues, err := service.GetLeagues()
	if err != nil {
		log.Printf("Error calling service GetLeagues: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, leagues)
}

func GetLeague(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	league, err := service.GetLeague(id)
	if err != nil {
		log.Errorf("Failure retrieving league with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if league == nil {
		http.Error(w, "404 league not found", http.StatusNotFound)
		return
	}
	sendJson(w, league)
}

func UpdateLeague(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	league, err := getLeague(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	league, err = service.UpdateLeague(id, league)
	if err != nil {
		log.Errorf("Failure updating league with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if league == nil {
		http.Error(w, "404 league not found", http.StatusNotFound)
		return
	}
	sendJson(w, league)
}

func DeleteLeague(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	league, err := service.DeleteLeague(id)
	if err != nil {
		log.Errorf("Failure deleting league with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if league == nil {
		http.Error(w, "404 league not found", http.StatusNotFound)
		return
	}
	sendJson(w, result{Success: "OK"})
}

func getLeague(r *http.Request) (*model.League, error) {
	var league model.League
	err := json.NewDecoder(r.Body).Decode(&league)
	if err != nil {
		log.Errorf("Can't serialize request body to league struct: %v", err)
		return nil, err
	}
	return &league, nil
}
