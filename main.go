package main

import (
	"fmt"
	"math"
)

type pair struct {
	A int
	B int
}

type quadruple struct {
	A int
	B int
	C int
	D int
}

func main() {
	// 100 for quick test
	const limit = 100

	for _, q := range findQuadruples(limit) {
		fmt.Printf("{a: %v, b: %v, c:%v, d:%v}\n", q.A, q.B, q.C, q.D)
	}
}

func findQuadruples(limit int) []quadruple {
	quadruples := []quadruple{}
	cubeSumResults := map[float64][]pair{}

	for a := 1; a <= limit; a++ {
		for b := a; b <= limit; b++ {
			cubeSum := math.Pow(float64(a), 3) + math.Pow(float64(b), 3)

			if _, ok := cubeSumResults[cubeSum]; ok {
				for _, n := range cubeSumResults[cubeSum] {
					quadruples = append(quadruples, quadruple{
						A: n.A,
						B: n.B,
						C: a,
						D: b,
					})
				}
			} else {
				cubeSumResults[cubeSum] = []pair{}
			}

			cubeSumResults[cubeSum] = append(cubeSumResults[cubeSum], pair{a, b})
		}
	}

	return quadruples
}
