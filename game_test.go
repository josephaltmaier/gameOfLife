package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestGame(t *testing.T) {
	suite.Run(t, new(GameTestSuite))
}

type GameTestSuite struct {
	suite.Suite
}

func (s *GameTestSuite) TestParse() {
	expectedBlinker := map[int]map[int]struct{}{
		5: {
			5: struct{}{},
			6: struct{}{},
			7: struct{}{},
		},
	}
	expectedGlider := map[int]map[int]struct{}{
		1: {
			2: struct{}{},
		},
		2: {
			3: struct{}{},
		},
		3: {
			1: struct{}{},
			2: struct{}{},
			3: struct{}{},
		},
	}

	parsedBlinker, err := parseInput(blinker)
	s.Require().NoError(err)
	s.Equal(expectedBlinker, parsedBlinker)

	parsedGlider, err := parseInput(glider)
	s.Require().NoError(err)
	s.Equal(expectedGlider, parsedGlider)
}

func (s *GameTestSuite) TestBlinker() {
	expectedGameState := map[int]map[int]struct{}{
		4: {6: struct{}{}},
		5: {6: struct{}{}},
		6: {6: struct{}{}},
	}

	initialState, err := parseInput(blinker)
	s.Require().NoError(err)

	game := newGame(initialState)
	err = game.Update()
	s.Require().NoError(err)
	s.Equal(expectedGameState, game.gameState)

	err = game.Update()
	s.Require().NoError(err)
	s.Equal(initialState, game.gameState)
}
