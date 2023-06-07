package flow

import (
	"bytes"
	"encoding/gob"
	"os"
	"path/filepath"
	"time"

	"github.com/LamkasDev/alpaca/cmd/common/arch"
)

type AlpacaFlow struct {
	Seed int64
}

func NewAlpacaFlow() *AlpacaFlow {
	return &AlpacaFlow{
		Seed: time.Now().UnixNano(),
	}
}

func WriteAlpacaFlow(flow *AlpacaFlow) {
	path := filepath.Join(arch.DataPath, "data", "flow.alpflow")
	var data bytes.Buffer
	if err := gob.NewEncoder(&data).Encode(flow); err != nil {
		panic(err)
	}
	os.WriteFile(path, data.Bytes(), 0644)
}

func ReadAlpacaFlow() *AlpacaFlow {
	path := filepath.Join(arch.DataPath, "data", "flow.alpflow")
	file, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var flow AlpacaFlow
	if err := gob.NewDecoder(bytes.NewBuffer(file)).Decode(&flow); err != nil {
		panic(err)
	}

	return &flow
}
