package searcher

import (
	"github.com/go-ego/riot/types"
	"reflect"
)

const (
	// SecondsInADay seconds in a day
	SecondsInADay = 86400
	// MaxTokenProximity max token proximity
	MaxTokenProximity = 2
)

/*******************************************************************************
    评分
*******************************************************************************/

// ScoringFields scoring fields
type ScoringFields struct {
	Timestamp int64
	View      int64
}

// ScoringCriteria custom scoring criteria
type ScoringCriteria struct {
}

// Score score and sort
func (criteria ScoringCriteria) Score(
	doc types.IndexedDoc, fields interface{}) []float32 {
	if reflect.TypeOf(fields) != reflect.TypeOf(ScoringFields{}) {
		return []float32{}
	}
	wsf := fields.(ScoringFields)
	output := make([]float32, 3)
	if doc.TokenProximity > MaxTokenProximity {
		output[0] = 1.0 / float32(doc.TokenProximity)
	} else {
		output[0] = 1.0
	}
	output[1] = float32(wsf.Timestamp / (SecondsInADay * 3))
	output[2] = doc.BM25 * (1 + float32(wsf.View)/10000)
	return output
}
