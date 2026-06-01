package main

import "github.com/g3n/engine/core"

type Player struct{}

type Position struct {
	X int
	Y int
	Z float32
}

type Renderable struct {
	node core.INode
}

type Movable struct{}

type Monster struct {
	Name string
}
