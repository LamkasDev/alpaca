package neuron

import (
	"fmt"
	"math/rand"

	"github.com/LamkasDev/alpaca/cmd/alpaca_ai/ai"
	"github.com/LamkasDev/alpaca/cmd/alpaca_ai/magic"
	"github.com/LamkasDev/alpaca/cmd/alpaca_ai/sets"
	"github.com/jwalton/gchalk"
)

type AlpacaNeuronComplex struct {
	Neurons []*AlpacaNeuron
	Result  *AlpacaNeuron
}

func NewAlpacaNeuronXOR(r *rand.Rand) *AlpacaNeuronComplex {
	return &AlpacaNeuronComplex{
		Neurons: []*AlpacaNeuron{NewAlpacaNeuron(r, sets.AlpacaTrainingSets["or"]()), NewAlpacaNeuron(r, sets.AlpacaTrainingSets["nand"]())},
		Result:  NewAlpacaNeuron(r, sets.AlpacaTrainingSets["and"]()),
	}
}

func TrainNeuronComplex(ccomplex *AlpacaNeuronComplex) {
	// TODO: finish this part (2:04:51).
	// cost := CostNeuronComplex(ccomplex)
	for _, cneuron := range ccomplex.Neurons {
		TrainNeuron(cneuron)
	}
	TrainNeuron(ccomplex.Result)
}

func ForwardNeuronComplex(ccomplex *AlpacaNeuronComplex, weights []float32) float32 {
	destinations := make([]float32, len(ccomplex.Neurons))
	for i, cneuron := range ccomplex.Neurons {
		for cw := 0; cw < len(cneuron.Weights); cw++ {
			destinations[i] += cneuron.Weights[cw] * weights[cw]
		}
		destinations[i] = magic.Sigmoid(destinations[i] + cneuron.Bias)
	}
	destination := float32(0)
	for cw := 0; cw < len(ccomplex.Result.Weights); cw++ {
		destination += ccomplex.Result.Weights[cw] * destinations[cw]
	}
	return magic.Sigmoid(destination + ccomplex.Result.Bias)
}

func CostNeuronComplex(ccomplex *AlpacaNeuronComplex) float32 {
	cost := float32(0)
	for _, row := range ccomplex.Result.Set.Data {
		y := ForwardNeuronComplex(ccomplex, row[:len(row)-1])
		d := y - row[len(row)-1]
		cost += d * d
	}
	cost /= float32(len(ccomplex.Result.Set.Data))

	return cost
}

func PrintNeuronComplex(cai *ai.AlpacaAI, ccomplex *AlpacaNeuronComplex) {
	fmt.Printf("result = %s\n", SprintNeuron(cai, ccomplex.Result))
	for i, cneuron := range ccomplex.Neurons {
		fmt.Printf("n[%d] = %s\n", i, SprintNeuron(cai, cneuron))
	}
}

func PrintNeuronComplexResults(ccomplex *AlpacaNeuronComplex) {
	for i := float32(0); i < totalWeights; i++ {
		for j := float32(0); j < totalWeights; j++ {
			fmt.Printf("%s | %s = %s\n", gchalk.Green(fmt.Sprint(i)), gchalk.Green(fmt.Sprint(j)), gchalk.Red(fmt.Sprint(magic.Sigmoid(i*ccomplex.Result.Weights[0]+j*ccomplex.Result.Weights[1]+ccomplex.Result.Bias))))
		}
	}
}
