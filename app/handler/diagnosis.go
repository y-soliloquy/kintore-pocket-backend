package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Answers struct {
	Answer string
}

type RequestBodyDiagnosis struct {
	Answers []string `json:"answers"`
}

type ResponseDiagnosis struct {
	Type           string   `json:"type"`
	Recomendations []string `json:"recommendations"`
}

type DiagnosisHandler struct {
	// dbやconfigを入れて拡張できるようにして多く
	// db *sql.DB
}

func NewDiagnosisHandler() *DiagnosisHandler {
	return &DiagnosisHandler{}
}

func (h *DiagnosisHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req RequestBodyDiagnosis
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("failed to decode request: %v", err)
		http.Error(w, "DiagnosisHandler: Invalid input", http.StatusBadRequest)
		return
	}

	counts := map[string]int{}
	for _, answer := range req.Answers {
		counts[answer]++
	}

	// 最も多いタイプ
	var maxType string
	maxCount := -1
	for typ, cnt := range counts {
		if cnt > maxCount {
			maxType = typ
			maxCount = cnt
		}
	}

	// タイプごとのメニューを定義（将来的にはDBや設定ファイルでもOK）
	recommendationMap := map[string][]string{
		"A": {"ピラミッド法", "アセンディング法", "ディセンディング法"},
		"B": {"5x5法", "3x3法"},
		"C": {"有酸素運動"},
	}

	resp := ResponseDiagnosis{
		Type:           maxType,
		Recomendations: recommendationMap[maxType],
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("DiagnoseHandler: failed to encode response: %v", err)
	}
}
