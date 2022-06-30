package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"strconv"
	"strings"
)

type ArrayGame struct {
	width         int
	height        int
	gameState     [][]uint8
	nextGameState [][]uint8
}

func newArrayGame(initialState [][]uint8) *ArrayGame {
	// Add empty spaces around all the edges to simplify neighbor checks
	width := len(initialState[0]) + 2
	height := len(initialState) + 2
	gameState := makeEmptyGameStateArray(width, height)
	for row, colList := range initialState {
		for col, state := range colList {
			gameState[row+1][col+1] = state
		}
	}

	return &ArrayGame{
		width:         width,
		height:        height,
		gameState:     gameState,
		nextGameState: makeEmptyGameStateArray(width, height),
	}
}

func parseInputArray(width int, height int, input string) ([][]uint8, error) {
	noWhitespace := strings.ReplaceAll(input, " ", "")
	pairs := strings.Split(noWhitespace, ",")

	gameState := makeEmptyGameStateArray(width, height)
	for i, pair := range pairs {
		rowAndCol := strings.Split(pair, ":")
		if len(rowAndCol) != 2 {
			return nil, fmt.Errorf("pair at index %d split into %d elements, must have exactly 2: %s", i, len(rowAndCol), pair)
		}

		row, err := strconv.Atoi(rowAndCol[0])
		if err != nil {
			return nil, fmt.Errorf("could not convert row %s to an int for pair at index %d: %s", rowAndCol[0], i, pair)
		}

		col, err := strconv.Atoi(rowAndCol[1])
		if err != nil {
			return nil, fmt.Errorf("could not convert column %s to an int for pair at index %d: %s", rowAndCol[1], i, pair)
		}

		gameState[row][col] = 1
	}

	return gameState, nil
}

func randomInputArray(width int, height int, probability float32) [][]uint8 {
	gameState := makeEmptyGameStateArray(width, height)
	for row := 1; row < height; row++ {
		for col := 1; col < width; col++ {
			if rand.Float32() < probability {
				gameState[row][col] = 1
			}
		}
	}

	return gameState
}

func makeEmptyGameStateArray(width int, height int) [][]uint8 {
	emptyGameState := make([][]uint8, height)
	for i := range emptyGameState {
		emptyGameState[i] = make([]uint8, width)
	}

	return emptyGameState
}

func (g *ArrayGame) Update() error {
	for row := 1; row < g.height-1; row++ {
		for col := 1; col < g.width-1; col++ {
			g.nextGameState[row][col] = 0

			neighbors := g.gameState[row-1][col-1] +
				g.gameState[row-1][col] +
				g.gameState[row-1][col+1] +
				g.gameState[row][col-1] +
				g.gameState[row][col+1] +
				g.gameState[row+1][col-1] +
				g.gameState[row+1][col] +
				g.gameState[row+1][col+1]

			curCell := g.gameState[row][col]
			if curCell == 1 && neighbors > 1 && neighbors < 4 {
				g.nextGameState[row][col] = 1
			} else if curCell == 0 && neighbors == 3 {
				g.nextGameState[row][col] = 1
			}
		}
	}

	var tmp = g.gameState
	g.gameState = g.nextGameState
	g.nextGameState = tmp
	return nil
}

func (g *ArrayGame) Draw(screen *ebiten.Image) {
	screen.Fill(deadColor)
	for row, colList := range g.gameState {
		for col, state := range colList {
			if state == 1 {
				screen.Set(row, col, liveColor)
			}
		}
	}
}

func (g *ArrayGame) Layout(_, _ int) (int, int) {
	return testWidth, testHeight
}
