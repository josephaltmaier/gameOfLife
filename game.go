package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"strconv"
	"strings"
)

type Game struct {
	gameState map[int]map[int]struct{}
}

func newGame(initialState map[int]map[int]struct{}) *Game {
	return &Game{gameState: initialState}
}

func parseInput(input string) (map[int]map[int]struct{}, error) {
	noWhitespace := strings.ReplaceAll(input, " ", "")
	pairs := strings.Split(noWhitespace, ",")

	gameState := make(map[int]map[int]struct{})
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

		setLive(row, col, gameState)
	}

	return gameState, nil
}

func randomInput(width int, height int, probability float32) map[int]map[int]struct{} {
	gameState := make(map[int]map[int]struct{})
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if rand.Float32() < probability {
				setLive(row, col, gameState)
			}
		}
	}

	return gameState
}

func setLive(row int, col int, gameState map[int]map[int]struct{}) {
	if row < 0 || col < 0 || row >= testHeight || col >= testWidth {
		return
	}

	liveCols, ok := gameState[row]
	if !ok {
		liveCols = make(map[int]struct{}, testHeight)
		gameState[row] = liveCols
	}
	liveCols[col] = struct{}{}
}

func update(gameState map[int]map[int]struct{}) map[int]map[int]struct{} {
	newGameState := make(map[int]map[int]struct{}, testWidth)
	for row, liveCols := range gameState {
		for col := range liveCols {
			// Determine the live status of this cell going forward
			neighbors := countNeighbors(row, col, gameState)
			if neighbors == 2 || neighbors == 3 {
				setLive(row, col, newGameState)
			}

			// Determine the live status of the surrounding dead cells
			for nRow := row - 1; nRow < row+2; nRow++ {
				for nCol := col - 1; nCol < col+2; nCol++ {
					if nLiveCols, ok := gameState[nRow]; ok {
						if _, ok := nLiveCols[nCol]; ok {
							continue
						}
					}
					neighbors := countNeighbors(nRow, nCol, gameState)
					if neighbors == 3 {
						setLive(nRow, nCol, newGameState)
					}
				}
			}
		}
	}

	return newGameState
}

func countNeighbors(row int, col int, gameState map[int]map[int]struct{}) int {
	neighbors := 0
	for nRow := row - 1; nRow < row+2; nRow++ {
		for nCol := col - 1; nCol < col+2; nCol++ {
			if colSet, ok := gameState[nRow]; ok {
				if _, ok := colSet[nCol]; ok {
					if nCol == col && nRow == row {
						continue
					}
					neighbors += 1
					if neighbors == 4 {
						return neighbors
					}
				}
			}
		}
	}

	return neighbors
}

func (g *Game) Update() error {
	g.gameState = update(g.gameState)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(deadColor)
	for row, colSet := range g.gameState {
		for col := range colSet {
			screen.Set(row, col, liveColor)
		}
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return testWidth, testHeight
}
