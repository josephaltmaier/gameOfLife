package main

import "testing"

func updateArrayBench(b *testing.B, initialState [][]uint8) {
	testGame := newArrayGame(initialState)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = testGame.Update()
	}
}

func BenchmarkArrayPulsar(b *testing.B) {
	initialState, err := parseInputArray(smallMapSize, smallMapSize, pulsar)
	if err != nil {
		panic(err)
	}
	updateArrayBench(b, initialState)
}

func BenchmarkSmallSparseArray(b *testing.B) {
	initialState := randomInputArray(smallMapSize, smallMapSize, sparseProbability)
	updateArrayBench(b, initialState)
}

func BenchmarkSmallDenseArray(b *testing.B) {
	initialState := randomInputArray(smallMapSize, smallMapSize, denseProbability)
	updateArrayBench(b, initialState)
}

func BenchmarkLargeSparseArray(b *testing.B) {
	initialState := randomInputArray(largeMapSize, largeMapSize, sparseProbability)
	updateArrayBench(b, initialState)
}

func BenchmarkLargeDenseArray(b *testing.B) {
	initialState := randomInputArray(largeMapSize, largeMapSize, denseProbability)
	updateArrayBench(b, initialState)
}
