package sets

import "github.com/LamkasDev/alpaca/cmd/alpaca_ai/set"

func NewAlpacaTrainingSetNOT(source *set.AlpacaTrainingSet) *set.AlpacaTrainingSet {
	cset := &set.AlpacaTrainingSet{
		Data: make([][]float32, len(source.Data)),
	}
	copy(cset.Data, source.Data)
	for i, row := range source.Data {
		if row[len(row)-1] == 1 {
			cset.Data[i][len(row)-1] = 0
		} else {
			cset.Data[i][len(row)-1] = 1
		}
	}

	return cset
}
