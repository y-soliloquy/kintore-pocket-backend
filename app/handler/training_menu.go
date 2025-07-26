package handler

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
)

// MinPlateWeight は、ジムのプレートに基づいた最小の重量刻みを定義します。
// 計算された重量を、使用可能なプレートに合わせて丸めるために使われます。
// トレーニーはプレート2枚を1組として利用します。
// 一般的な最小のプレートの重さは1.25kgです。
const MinPlateWeight float64 = 2.5

type TrainingTemplate struct {
	Set     int     `json:"set"`
	Percent float64 `json:"percent"`
	Reps    int     `json:"reps"`
}

type TrainingMenu struct {
	Set    int     `json:"set"`
	Weight float64 `json:"weight"`
	Reps   int     `json:"reps"`
}

type RequestBodyTrainingMenu struct {
	Weight int `json:"weight"`
}

type TrainingMenuHandler struct {
	BasePath string
}

func NewTrainingMenuHandler(path string) *TrainingMenuHandler {
	return &TrainingMenuHandler{
		BasePath: path,
	}
}

// ゆくゆくは、メニュータイプも受け取って、このAPIだけで複数のメニューを返すようにしたい。
// その際は名前を汎用的なものに変更する
func (h *TrainingMenuHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req RequestBodyTrainingMenu
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("failed to decode request: %v", err)
		http.Error(w, "TrainingMenuHandler: Invalid input", http.StatusBadRequest)
		return
	}

	templateName := r.URL.Query().Get("template")
	if templateName == "" {
		http.Error(w, "template query parameter is required", http.StatusBadRequest)
		return
	}

	templatePath := filepath.Join(h.BasePath, templateName)
	data, err := os.ReadFile(templatePath)
	if err != nil {
		log.Printf("failed to read file: %v", err)
		http.Error(w, "failed to read template", http.StatusInternalServerError)
		return
	}

	var template []TrainingTemplate
	if err := json.Unmarshal(data, &template); err != nil {
		log.Printf("failed to parse: %v", err)
		http.Error(w, "TrainingMenuHandler: Invalid template", http.StatusInternalServerError)
		return
	}

	var menus []TrainingMenu
	for _, t := range template {
		weight := CalculateWeight(float64(req.Weight) * t.Percent)
		menus = append(menus, TrainingMenu{
			Set:    t.Set,
			Weight: weight,
			Reps:   t.Reps,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(menus); err != nil {
		log.Printf("TrainingMenuHandler: failed to encode JSON: %v", err)
	}
}

// CalculateWeight トレーニングメニュー上の重さを計算する関数
func CalculateWeight(x float64) float64 {
	const epsilon = 0.0001
	return math.Round((x+epsilon)/MinPlateWeight) * MinPlateWeight
}
