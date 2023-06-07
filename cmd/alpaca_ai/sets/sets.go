package sets

import "github.com/LamkasDev/alpaca/cmd/alpaca_ai/set"

type AlpacaTrainingSetFetch func() *set.AlpacaTrainingSet

var AlpacaTrainingSets = map[string]AlpacaTrainingSetFetch{
	"and":  func() *set.AlpacaTrainingSet { return NewAlpacaTrainingSetAND() },
	"nand": func() *set.AlpacaTrainingSet { return NewAlpacaTrainingSetNOT(NewAlpacaTrainingSetAND()) },
	"or":   func() *set.AlpacaTrainingSet { return NewAlpacaTrainingSetOR() },
	"xor":  func() *set.AlpacaTrainingSet { return NewAlpacaTrainingSetXOR() },
}
