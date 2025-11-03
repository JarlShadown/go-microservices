package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Message any  `json:"message,omitempty"`
	Status  int  `json:"status"`
	Error   bool `json:"error"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Message: "Hit the broker",
		Status:  0,
		Error:   false,
	}
	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}
