package sets

import "github.com/LamkasDev/alpaca/cmd/alpaca_ai/set"

func NewAlpacaTrainingSetXOR() *set.AlpacaTrainingSet {
	return &set.AlpacaTrainingSet{
		Data: [][]float32{
			{0, 0, 0},
			{0, 1, 1},
			{1, 0, 1},
			{1, 1, 0},
		},
	}
}
