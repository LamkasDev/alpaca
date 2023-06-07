package run

import (
	"fmt"
	"math/rand"

	"github.com/LamkasDev/alpaca/cmd/alpaca_ai/ai"
	"github.com/LamkasDev/alpaca/cmd/alpaca_ai/flow"
	"github.com/LamkasDev/alpaca/cmd/alpaca_ai/neuron"
	"github.com/jwalton/gchalk"
)

func LoadAI(cai *ai.AlpacaAI) {
	const cset = "nand"
	/* cai.Set = set.ReadAlpacaTrainingSet(cset)
	if cai.Set == nil {
		cai.Set = sets.AlpacaTrainingSets[cset]()
		set.WriteAlpacaTrainingSet(cai.Set, cset)
	} */
	cai.Flow = flow.ReadAlpacaFlow()
	if cai.Flow == nil {
		cai.Flow = flow.NewAlpacaFlow()
		flow.WriteAlpacaFlow(cai.Flow)
	}
}

func RunAI(cai *ai.AlpacaAI) {
	r := rand.New(rand.NewSource(cai.Flow.Seed))
	ccomplex := neuron.NewAlpacaNeuronXOR(r)

	targetEpoch := 1000 * 1000

	fmt.Printf("=> seed: %s | target epoch: %s\n", gchalk.Yellow(fmt.Sprint(cai.Flow.Seed)), gchalk.Yellow(fmt.Sprint(targetEpoch)))
	neuron.PrintNeuronComplex(cai, ccomplex)
	for i := 0; i < targetEpoch; i++ {
		neuron.TrainNeuronComplex(ccomplex)
	}
	neuron.PrintNeuronComplex(cai, ccomplex)
	neuron.PrintNeuronComplexResults(ccomplex)

	return
}
