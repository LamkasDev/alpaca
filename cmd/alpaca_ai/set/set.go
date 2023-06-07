package set

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"

	"github.com/LamkasDev/alpaca/cmd/common/arch"
)

type AlpacaTrainingSet struct {
	Data [][]float32
}

func WriteAlpacaTrainingSet(set *AlpacaTrainingSet, name string) {
	path := filepath.Join(arch.DataPath, "data", fmt.Sprintf("%s.alpset", name))
	var data bytes.Buffer
	if err := gob.NewEncoder(&data).Encode(set); err != nil {
		panic(err)
	}
	os.WriteFile(path, data.Bytes(), 0644)
}

func ReadAlpacaTrainingSet(name string) *AlpacaTrainingSet {
	path := filepath.Join(arch.DataPath, "data", fmt.Sprintf("%s.alpset", name))
	file, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var set AlpacaTrainingSet
	if err := gob.NewDecoder(bytes.NewBuffer(file)).Decode(&set); err != nil {
		panic(err)
	}

	return &set
}
