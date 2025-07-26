package diagnosis

type TrainingType string

const (
	MuscleGainType TrainingType = "A"
	PowerType      TrainingType = "B"
	AerobicType    TrainingType = "C"
)

type SingleResult struct {
	TrainingType    TrainingType `json:"type"`
	Recommendations []string     `json:"recommendations"`
}

type Result struct {
	Results []SingleResult `json:"results"`
}

var recommendationMap = map[TrainingType][]string{
	MuscleGainType: {"ピラミッド法", "アセンディング法", "ディセンディング法"},
	PowerType:      {"5x5法", "3x3法"},
	AerobicType:    {"有酸素運動"},
}

func Diagnose(answers []string) Result {
	counts := map[string]int{}
	for _, a := range answers {
		counts[a]++
	}

	// 最大票数を調べる
	maxCount := -1
	for _, cnt := range counts {
		if cnt > maxCount {
			maxCount = cnt
		}
	}

	// 最大票数と同じタイプを抽出
	var results []SingleResult
	for typStr, cnt := range counts {
		if cnt == maxCount {
			t := TrainingType(typStr)
			results = append(results, SingleResult{
				TrainingType:    t,
				Recommendations: recommendationMap[t],
			})
		}
	}

	return Result{
		Results: results,
	}
}
