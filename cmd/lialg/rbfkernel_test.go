package lialg

import (
	"math"
	"testing"
)

func TestRBFKernelBasic(t *testing.T) {
	x := []float64{1, 1}
	y := []float64{2, 2}

	value := RbfKernel(x, y, 1)
	expected := math.Exp(-1)

	// check if floating point error is within bounds
	if math.Abs(value-expected) > 1e-14 {
		t.Errorf("Expected: %f ; Got: %f", expected, value)
	}
}
