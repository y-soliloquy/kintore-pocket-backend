package handler

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"os"
)

// MinPlateWeight は、ジムのプレートに基づいた最小の重量刻みを定義します。
// 計算された重量を、使用可能なプレートに合わせて丸めるために使われます。
// トレーニーはプレート2枚を1組として利用します。
// 一般的な最小のプレートの重さは1.25kgです。
const MinPlateWeight float64 = 2.5

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

func (h *FiveTimesFiveHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "FiveTimesFiveHandler: Invalid input", http.StatusBadRequest)
		return
	}

	data, err := os.ReadFile("../../data/5x5.json")
	if err != nil {
		http.Error(w, "FiveTimesFiveHandler: Template not found", http.StatusInternalServerError)
		return
	}

	var template []FiveTimesFiveTemplate
	if err := json.Unmarshal(data, &template); err != nil {
		http.Error(w, "FiveTimesFiveHandler: Invalid template", http.StatusInternalServerError)
		return
	}

	var menus []FiveTimesFiveMenu
	for _, t := range template {
		weight := CalculateWeight(float64(req.Weight) * t.Percent)
		menus = append(menus, FiveTimesFiveMenu{
			Set:    t.Set,
			Weight: int(weight),
			Reps:   t.Reps,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(menus); err != nil {
		log.Printf("FiveTimesFiveHandler: failed to encode JSON: %v", err)
	}
}

// CalculateWeight トレーニングメニュー上の重さを計算する関数
func CalculateWeight(x float64) float64 {
	return math.Round(x/MinPlateWeight) * MinPlateWeight
}
