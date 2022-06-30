package main

import "testing"

const (
	smallMapSize      = 100
	largeMapSize      = 10000
	sparseProbability = .1
	denseProbability  = .5
)

func updateBench(b *testing.B, gameState map[int]map[int]struct{}) {
	testGame := newGame(gameState)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = testGame.Update()
	}
}

func BenchmarkPulsar(b *testing.B) {
	gameState, err := parseInput(pulsar)
	if err != nil {
		panic(err)
	}

	updateBench(b, gameState)
}

func BenchmarkSmallSparse(b *testing.B) {
	gameState := randomInput(smallMapSize, smallMapSize, sparseProbability)
	updateBench(b, gameState)
}

func BenchmarkSmallDense(b *testing.B) {
	gameState := randomInput(smallMapSize, smallMapSize, denseProbability)
	updateBench(b, gameState)
}

func BenchmarkLargeSparse(b *testing.B) {
	gameState := randomInput(largeMapSize, largeMapSize, sparseProbability)
	updateBench(b, gameState)
}

func BenchmarkLargeDense(b *testing.B) {
	gameState := randomInput(largeMapSize, largeMapSize, denseProbability)
	updateBench(b, gameState)
}
