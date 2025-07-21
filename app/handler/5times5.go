package handler

import (
	"log"
	"net/http"
)

type FiveTimesFiveTemplate struct {
	Set     int     `json:"set"`
	Percent float64 `json:"percent"`
	Reps    int     `json:"rep"`
}

type FiveTimesFiveMenu struct {
	Set    int `json:"set"`
	Weight int `json:"weight"`
	Reps   int `json:"rep"`
}

type RequestBody struct {
	Weight int `json:"weight"`
}

type FiveTimesFiveHandler struct{}

func NewFiveTimesFiveHandler() *FiveTimesFiveHandler {
	return &FiveTimesFiveHandler{}
}

// 一旦仮で書く
func (h *FiveTimesFiveHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("ok"))
	if err != nil {
		log.Printf("FiveTimesFiveHandler: failed to write response: %v", err)
	}
}
