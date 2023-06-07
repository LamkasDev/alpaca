package ai

import (
	"github.com/LamkasDev/alpaca/cmd/alpaca_ai/config"
	"github.com/LamkasDev/alpaca/cmd/alpaca_ai/flow"
)

type AlpacaAI struct {
	Config config.AlpacaConfig
	// Set    *set.AlpacaTrainingSet
	Flow *flow.AlpacaFlow
}

func SetupAI() AlpacaAI {
	ai := AlpacaAI{}
	var err error
	if ai.Config, err = config.ReadConfig(ai.Config); err != nil {
		panic(err)
	}

	return ai
}
