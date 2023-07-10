package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/model"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/service"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	event, err := getEvent(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = service.CreateEvent(id, event); err != nil {
		log.Printf("Error calling service CreateEvent: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, event)
}

func GetEvents(w http.ResponseWriter, _ *http.Request) {
	events, err := service.GetEvents()
	if err != nil {
		log.Printf("Error calling service GetEvents: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, events)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	event, err := service.GetEvent(id)
	if err != nil {
		log.Errorf("Failure retrieving event with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if event == nil {
		http.Error(w, "404 event not found", http.StatusNotFound)
		return
	}
	sendJson(w, event)
}

func getEvent(r *http.Request) (*model.Event, error) {
	var event model.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		log.Errorf("Can't serialize request body to event struct: %v", err)
		return nil, err
	}
	return &event, nil
}
