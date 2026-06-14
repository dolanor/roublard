package main

import (
	"math"

	"github.com/g3n/engine/core"
)

type Player struct{}

type Position struct {
	X int
	Y int
	Z float32
}

func (p *Position) GetManhattanDistance(other *Position) int {
	xDist := math.Abs(float64(p.X - other.X))
	yDist := math.Abs(float64(p.Y - other.Y))

	return int(xDist) + int(yDist)
}

type Renderable struct {
	node core.INode
}

type Movable struct{}

type Monster struct {
	Name string
}
