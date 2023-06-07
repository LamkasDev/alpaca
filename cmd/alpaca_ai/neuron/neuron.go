package neuron

import (
	"fmt"
	"math/rand"

	"github.com/LamkasDev/alpaca/cmd/alpaca_ai/ai"
	"github.com/LamkasDev/alpaca/cmd/alpaca_ai/magic"
	"github.com/LamkasDev/alpaca/cmd/alpaca_ai/set"
	"github.com/jwalton/gchalk"
)

type AlpacaNeuron struct {
	Set     *set.AlpacaTrainingSet
	Weights []float32
	Bias    float32
}

const totalWeights = 2
const eps = float32(1e-1)
const rate = float32(1e-1)

func NewAlpacaNeuron(r *rand.Rand, set *set.AlpacaTrainingSet) *AlpacaNeuron {
	cneuron := AlpacaNeuron{
		Set:     set,
		Weights: make([]float32, totalWeights),
	}
	for cw := 0; cw < totalWeights; cw++ {
		cneuron.Weights[cw] = r.Float32()
	}
	cneuron.Bias = r.Float32()

	return &cneuron
}

func TrainNeuron(cneuron *AlpacaNeuron) {
	cost := magic.Cost(cneuron.Set, cneuron.Weights, cneuron.Bias)
	destinationWeights := []float32{
		(magic.Cost(cneuron.Set, []float32{cneuron.Weights[0] + eps, cneuron.Weights[1]}, cneuron.Bias) - cost) / eps,
		(magic.Cost(cneuron.Set, []float32{cneuron.Weights[0], cneuron.Weights[1] + eps}, cneuron.Bias) - cost) / eps,
	}
	destinationBias := (magic.Cost(cneuron.Set, []float32{cneuron.Weights[0], cneuron.Weights[1]}, cneuron.Bias+eps) - cost) / eps

	for cw := 0; cw < len(cneuron.Weights); cw++ {
		cneuron.Weights[cw] -= rate * destinationWeights[cw]
	}
	cneuron.Bias -= rate * destinationBias
}

func PrintNeuron(cai *ai.AlpacaAI, cneuron *AlpacaNeuron) {
	fmt.Println(SprintNeuron(cai, cneuron))
}

func SprintNeuron(cai *ai.AlpacaAI, cneuron *AlpacaNeuron) string {
	return fmt.Sprintf("w: %s | bias: %s | cost: %s", gchalk.Red(fmt.Sprint(cneuron.Weights)), gchalk.Red(fmt.Sprint(cneuron.Bias)), gchalk.Red(fmt.Sprint(magic.Cost(cneuron.Set, cneuron.Weights, cneuron.Bias))))
}

func PrintNeuronResults(cneuron *AlpacaNeuron) {
	for i := float32(0); i < totalWeights; i++ {
		for j := float32(0); j < totalWeights; j++ {
			fmt.Printf("%s | %s = %s\n", gchalk.Green(fmt.Sprint(i)), gchalk.Green(fmt.Sprint(j)), gchalk.Red(fmt.Sprint(magic.Sigmoid(i*cneuron.Weights[0]+j*cneuron.Weights[1]+cneuron.Bias))))
		}
	}
}
