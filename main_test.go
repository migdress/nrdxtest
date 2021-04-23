package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findQuadruples(t *testing.T) {
	// Set this to 10000 to match the exercise requirement
	// Set test timeout high for high limits
	const limit = 100
	//const limit = 10000

	// Act
	quadruples := findQuadruples(limit)

	// Assert not empty
	require.NotEmpty(t, quadruples)

	for _, q := range quadruples {
		// Assert result is correct
		a3 := math.Pow(float64(q.A), 3)
		b3 := math.Pow(float64(q.B), 3)
		c3 := math.Pow(float64(q.C), 3)
		d3 := math.Pow(float64(q.D), 3)
		require.True(t, a3+b3 == c3+d3)
	}
}
