package magic

import (
	"maze.io/x/math32"
)

func Sigmoid(x float32) float32 {
	return 1 / (1 + math32.Exp(-x))
}
