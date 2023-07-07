package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/model"
	"github.com/slimsamil/KreisliGo-go-ss2023/src/kreisligo/service"
)

func CreateAssociation(w http.ResponseWriter, r *http.Request) {
	association, err := getAssociation(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = service.CreateAssociation(association); err != nil {
		log.Printf("Error calling service CreateAssociation: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, association)
}

func GetAssociations(w http.ResponseWriter, _ *http.Request) {
	associations, err := service.GetAssociations()
	if err != nil {
		log.Printf("Error calling service GetAssociations: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, associations)
}

func GetAssociation(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	association, err := service.GetAssociation(id)
	if err != nil {
		log.Errorf("Failure retrieving association with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if association == nil {
		http.Error(w, "404 association not found", http.StatusNotFound)
		return
	}
	sendJson(w, association)
}

func UpdateAssociation(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	association, err := getAssociation(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	association, err = service.UpdateAssociation(id, association)
	if err != nil {
		log.Errorf("Failure updating association with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if association == nil {
		http.Error(w, "404 association not found", http.StatusNotFound)
		return
	}
	sendJson(w, association)
}

func DeleteAssociation(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	association, err := service.DeleteAssociation(id)
	if err != nil {
		log.Errorf("Failure deleting association with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if association == nil {
		http.Error(w, "404 association not found", http.StatusNotFound)
		return
	}
	sendJson(w, result{Success: "OK"})
}

func getAssociation(r *http.Request) (*model.Association, error) {
	var association model.Association
	err := json.NewDecoder(r.Body).Decode(&association)
	if err != nil {
		log.Errorf("Can't serialize request body to association struct: %v", err)
		return nil, err
	}
	return &association, nil
}
