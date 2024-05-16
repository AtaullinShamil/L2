package handler

import (
	"encoding/json"
	"github.com/AtaullinShamil/wbschool_exam_L2/tree/main/develop/dev11/internal/model"
	"net/http"
)

func EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Wrong method : only Get is allowed",
		})
		return
	}
	date, err := GetDateQuery(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Query parameter missing",
		})
		return
	}
	events, err := model.Cache.FindEventsByDate(date, "day")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Couldn't find events",
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]model.Event{
		"result": events,
	})
}
