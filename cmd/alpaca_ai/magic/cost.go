package magic

import (
	"github.com/LamkasDev/alpaca/cmd/alpaca_ai/set"
)

func Cost(set *set.AlpacaTrainingSet, weights []float32, bias float32) float32 {
	cost := float32(0)
	for _, dataRow := range set.Data {
		y := float32(0)
		for cw := 0; cw < len(dataRow)-1; cw++ {
			x := dataRow[cw]
			y += x * weights[cw]
		}
		y = Sigmoid(y + bias)

		destination := y - dataRow[len(dataRow)-1]
		cost += destination * destination
	}
	cost /= float32(len(set.Data))

	return cost
}
