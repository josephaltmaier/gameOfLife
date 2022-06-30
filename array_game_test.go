package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestArrayGame(t *testing.T) {
	suite.Run(t, new(ArrayGameTestSuite))
}

type ArrayGameTestSuite struct {
	suite.Suite
}

func (s *ArrayGameTestSuite) TestParse() {
	expectedBlinker := makeEmptyGameStateArray(testWidth, testHeight)
	expectedBlinker[5][5] = 1
	expectedBlinker[5][6] = 1
	expectedBlinker[5][7] = 1

	expectedGlider := makeEmptyGameStateArray(testWidth, testHeight)
	expectedGlider[1][2] = 1
	expectedGlider[2][3] = 1
	expectedGlider[3][1] = 1
	expectedGlider[3][2] = 1
	expectedGlider[3][3] = 1

	parsedBlinker, err := parseInputArray(testWidth, testHeight, blinker)
	s.Require().NoError(err)
	s.Equal(expectedBlinker, parsedBlinker)

	parsedGlider, err := parseInputArray(testWidth, testHeight, glider)
	s.Require().NoError(err)
	s.Equal(expectedGlider, parsedGlider)
}

func (s *ArrayGameTestSuite) TestBlinker() {
	expectedInitialState := makeEmptyGameStateArray(testWidth+2, testHeight+2)
	expectedInitialState[6][6] = 1
	expectedInitialState[6][7] = 1
	expectedInitialState[6][8] = 1
	expectedTransitionState := makeEmptyGameStateArray(testWidth+2, testHeight+2)
	expectedTransitionState[5][7] = 1
	expectedTransitionState[6][7] = 1
	expectedTransitionState[7][7] = 1

	initialState, err := parseInputArray(testWidth, testHeight, blinker)
	s.Require().NoError(err)

	game := newArrayGame(initialState)
	s.Equal(expectedInitialState, game.gameState)
	err = game.Update()
	s.Require().NoError(err)
	s.Equal(expectedTransitionState, game.gameState)

	err = game.Update()
	s.Require().NoError(err)
	s.Equal(expectedInitialState, game.gameState)
}
