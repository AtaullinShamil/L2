package handler

import (
	"encoding/json"
	"github.com/AtaullinShamil/wbschool_exam_L2/tree/main/develop/dev11/internal/model"
	"net/http"
	"time"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "only Post method allowed",
		})
		return
	}

	input := inputData{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	start, err := time.Parse("2006.01.02", input.Start)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid start time format",
		})
		return
	}
	end, err := time.Parse("2006.01.02", input.End)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid end time format",
		})
		return
	}

	event, err := model.Cache.CreateEvent(input.Title, input.Description, start, end)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "couldn't create event",
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]model.Event{
		"result": event,
	})
}
