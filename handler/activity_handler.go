package handler

import (
	"fmt"
	"log"
	"net/http"
)

func ActivitiesHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Activities handler")
	query := r.URL.Query()
	activity := query.Get("activity")
	if activity == "" {
		activity = "N/A"
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Activities: %v\n", activity)
}
