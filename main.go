package main

import (
	"fmt"
	"github.com/deadsy/sdfx/sdf"
	"math"
	"math/rand"
)

type tset []*sdf.Triangle3

// RaisePoint will set all points a the given x/y coordinate to the given height
func (s tset) RaisePoint(x, y, z float64) {
	for _, t := range s {
		for k := range t.V {
			if t.V[k].X == x && t.V[k].Y == y {
				t.V[k].Z = z
			}
		}
	}
}

func main() {
	// heights need to be divisible by 2
	heights := []float64{}
	for i := 0; i <= 256; i++ {
		heights = append(heights, rand.Float64()*20)
	}
	saveMesh(heights)
}

func saveMesh(heightOffsets []float64) {

	if len(heightOffsets)%2 != 1 {
		panic("input length must be divisible by 2")
	}

	baseHeight := 5.0
	squareSize := 30.0
	numXYSquares := math.Floor(math.Sqrt(float64(len(heightOffsets))))
	sideLength := numXYSquares * squareSize

	fmt.Printf("squares: %0.2f\n", numXYSquares)

	tris := tset{}
	for col := 0.0; col < numXYSquares; col++ {
		for row := 0.0; row < numXYSquares; row++ {

			xo := col * squareSize
			yo := row * squareSize

			tl := sdf.V3{X: xo, Y: yo, Z: 0}
			tr := sdf.V3{X: xo + squareSize, Y: yo, Z: 0}
			bl := sdf.V3{X: xo, Y: yo + squareSize, Z: 0}
			br := sdf.V3{X: xo + squareSize, Y: yo + squareSize, Z: 0}

			// add a square
			tris = append(
				tris,
				sdf.NewTriangle3(tl, tr, bl),
				sdf.NewTriangle3(tr, br, bl),
			)
		}
	}

	curOffset := 0
	for col := 1.0; col < numXYSquares; col++ {
		for row := 1.0; row < numXYSquares; row++ {
			if curOffset > len(heightOffsets)-1 {
				continue
			}

			xo := col * squareSize
			yo := row * squareSize

			tris.RaisePoint(xo, yo, heightOffsets[curOffset])
			curOffset++

		}
	}

	// create a base
	tris = append(
		tris,
		//N wall
		sdf.NewTriangle3(
			sdf.V3{X: 0, Y: 0, Z: 0},
			sdf.V3{X: 0, Y: 0, Z: 0-baseHeight},
			sdf.V3{X: sideLength, Y: 0, Z: 0},
		),
		sdf.NewTriangle3(
			sdf.V3{X: sideLength, Y: 0, Z: 0},
			sdf.V3{X: 0, Y: 0, Z: 0-baseHeight},
			sdf.V3{X: sideLength, Y: 0, Z: 0-baseHeight},
		),
		//E wall
		sdf.NewTriangle3(
			sdf.V3{X: sideLength, Y: 0, Z: 0},
			sdf.V3{X: sideLength, Y: 0, Z: 0-baseHeight},
			sdf.V3{X: sideLength, Y: sideLength, Z: 0},
		),
		sdf.NewTriangle3(
			sdf.V3{X: sideLength, Y: 0, Z: 0-baseHeight},
			sdf.V3{X: sideLength, Y: sideLength, Z: 0-baseHeight},
			sdf.V3{X: sideLength, Y: sideLength, Z: 0},
		),
		//S wall
		sdf.NewTriangle3(
			sdf.V3{X: 0, Y: sideLength, Z: 0-baseHeight},
			sdf.V3{X: 0, Y: sideLength, Z: 0},
			sdf.V3{X: sideLength, Y: sideLength, Z: 0},
		),
		sdf.NewTriangle3(
			sdf.V3{X: 0, Y: sideLength, Z: 0-baseHeight},
			sdf.V3{X: sideLength, Y: sideLength, Z: 0},
			sdf.V3{X: sideLength, Y: sideLength, Z: 0-baseHeight},
		),
		//W wall
		sdf.NewTriangle3(
			sdf.V3{X: 0, Y: 0, Z: 0},
			sdf.V3{X: 0, Y: sideLength, Z: 0},
			sdf.V3{X: 0, Y: 0, Z: 0-baseHeight},
		),
		sdf.NewTriangle3(
			sdf.V3{X: 0, Y: 0, Z: 0-baseHeight},
			sdf.V3{X: 0, Y: sideLength, Z: 0},
			sdf.V3{X: 0, Y: sideLength, Z: 0-baseHeight},
		),
		//bottom
		sdf.NewTriangle3(
			sdf.V3{X: 0, Y: sideLength, Z: 0-baseHeight},
			sdf.V3{X: sideLength, Y: 0, Z: 0-baseHeight},
			sdf.V3{X: 0, Y: 0, Z: 0-baseHeight},
		),
		sdf.NewTriangle3(
			sdf.V3{X: sideLength, Y: sideLength, Z: 0-baseHeight},
			sdf.V3{X: sideLength, Y: 0, Z: 0-baseHeight},
			sdf.V3{X: 0, Y: sideLength, Z: 0-baseHeight},
		),
	)

	sdf.SaveSTL("output.stl", tris)
}
