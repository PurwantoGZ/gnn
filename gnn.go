package gnn

import "math"

func F(x float64) float64 {
	return (1.0 / (1.0 + math.Exp(-x)))
}
