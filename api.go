package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type API struct {
	service Service
}

func NewAPIService(service Service) *API {
	return &API{service: service}
}

func (s *API) Start(address string) error {

	http.HandleFunc("/", s.handleGetActivityRequest)

	fmt.Printf("Server started at %s\n", address)

	return http.ListenAndServe(address, nil)
}

func (s *API) handleGetActivityRequest(w http.ResponseWriter, r *http.Request) {
	activity, err := s.service.GetActivity(context.Background())
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err.Error()})
	}
	writeJSON(w, http.StatusOK, activity)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(v)
}
