package gogame

import "github.com/twmb/algoimpl/go/graph"

type pointState uint8

const (
	unset pointState = iota
	Black
	White
)

type move struct {
	stone    pointState
	pointIdx int
}

type game struct {
	history  []move
	board    []graph.Node
	position []pointState
}

func Standard(width, height uint) *game {
	return nil
}

func (g *game) Move(player pointState, pointName string) error {
	return nil
}
