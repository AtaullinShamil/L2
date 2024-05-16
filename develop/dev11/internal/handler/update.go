package handler

import (
	"encoding/json"
	"github.com/AtaullinShamil/wbschool_exam_L2/tree/main/develop/dev11/internal/model"
	"net/http"
)

func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "only Post method allowed",
		})
		return
	}
	updatedEvent, err := getEventForUpdateOrDelete(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "couldn't create event for update",
		})
		return
	}
	err = model.Cache.UpdateEvent(updatedEvent.ID, updatedEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "event was not found by ID",
		})
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]model.Event{
			"result": updatedEvent,
		})
	}
}
