package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

const (
	blinker = "5:5, 5:6, 5:7"
	beacon  = "1:1, 1:2, 2:1, 3:4, 4:3, 4:4"
	glider  = "1:2, 2:3, 3:1, 3:2, 3:3"
	pulsar  = "2:4, 2:5, 2:6, 2:10, 2:11, 2:12, " +
		"4:2, 4:7, 4:9, 4:14, " +
		"5:2, 5:7, 5:9, 5:14, " +
		"6:2, 6:7, 6:9, 6:14, " +
		"7:4, 7:5, 7:6, 7:10, 7:11, 7:12, " +
		"9:4, 9:5, 9:6, 9:10, 9:11, 9:12, " +
		"10:2, 10:7, 10:9, 10:14, " +
		"11:2, 11:7, 11:9, 11:14, " +
		"12:2, 12:7, 12:9, 12:14, " +
		"14:4, 14:5, 14:6, 14:10, 14:11, 14:12"
	testWidth  = 100
	testHeight = 100
)

var (
	deadColor = color.RGBA{
		R: 0xFF,
		G: 0xFF,
		B: 0xFF,
		A: 0xFF,
	}

	liveColor = color.RGBA{
		R: 0x00,
		G: 0x00,
		B: 0x00,
		A: 0xFF,
	}
)

func main() {
	ebiten.SetWindowTitle("Joseph's CGL Implementation")
	ebiten.SetMaxTPS(10)

	//initialState := randomInput(testWidth, testHeight, .5)
	initialState, err := parseInput(pulsar)
	if err != nil {
		panic(err)
	}
	if err := ebiten.RunGame(newGame(initialState)); err != nil {
		panic(err)
	}

	//initialState := randomInputArray(testWidth, testHeight, .5)
	//initialState, err := parseInputArray(testWidth, testHeight, pulsar)
	//if err != nil {
	//	panic(err)
	//}
	//if err := ebiten.RunGame(newArrayGame(initialState)); err != nil {
	//	panic(err)
	//}
}
