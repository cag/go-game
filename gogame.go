package gogame

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

type pointState uint8

const (
	unset pointState = iota
	Black
	White
)

func nameForPointState(s pointState) string {
	switch s {
	case unset:
		return "unset"
	case Black:
		return "Black"
	case White:
		return "White"
	}
	return ""
}

type move struct {
	stone    pointState
	pointIdx int
}

type point struct {
	index  int
	adjpts []*point
}

type game struct {
	history          []move
	board            []point
	position         []pointState
	lastPosition     []pointState
	pointIdxPlus1Map map[string]int
	formatter        func() string
}

func linkPoints(p1, p2 *point) {
	p1.adjpts = append(p1.adjpts, p2)
	p2.adjpts = append(p2.adjpts, p1)
}

func (p *point) chainSurrounded(position []pointState) bool {
	visited := make([]bool, len(position))
	var traveler func(p2 *point) bool
	traveler = func(p2 *point) bool {
		visited[p2.index] = true
		curPtState := position[p2.index]
		for _, adjpt := range p2.adjpts {
			adjptState := position[adjpt.index]
			if adjptState == unset {
				return false
			}
			if adjptState == curPtState && !visited[adjpt.index] {
				if !traveler(adjpt) {
					return false
				}
			}
		}
		return true
	}
	return traveler(p)
}

func (p *point) removeChain(position []pointState) {
	stone := position[p.index]
	if stone == unset {
		panic(fmt.Sprintf("Can't remove unset point %v from position %v", p, position))
	}
	var traveler func(p2 *point)
	traveler = func(p2 *point) {
		position[p2.index] = unset
		for _, adjpt := range p2.adjpts {
			if position[adjpt.index] == stone {
				traveler(adjpt)
			}
		}
	}
	traveler(p)
}

func Standard(width, height int) *game {
	g := game{
		board:            make([]point, width*height),
		position:         make([]pointState, width*height),
		pointIdxPlus1Map: make(map[string]int),
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			var name string
			if i < 8 {
				name = fmt.Sprintf("%c%d", []rune("A")[0]+rune(i), j+1)
			} else {
				name = fmt.Sprintf("%c%d", []rune("A")[0]+rune(i+1), j+1)
			}

			idx := i*height + j

			if i > 0 {
				linkPoints(&g.board[idx], &g.board[idx-height])
			}
			if j > 0 {
				linkPoints(&g.board[idx], &g.board[idx-1])
			}

			g.board[idx].index = idx
			g.pointIdxPlus1Map[name] = idx + 1
		}
	}

	g.formatter = func() string {
		var buffer bytes.Buffer
		for j := height - 1; j >= 0; j-- {
			for i := 0; i < width; i++ {
				if i > 0 {
					buffer.WriteByte(' ')
				}
				idx := i*height + j
				switch g.position[idx] {
				case unset:
					buffer.WriteRune('·')
				case Black:
					buffer.WriteRune('●')
				case White:
					buffer.WriteRune('○')
				}
			}
			buffer.WriteByte('\n')
		}
		return buffer.String()
	}
	return &g
}

func (g *game) Move(stone pointState, pointName string) error {
	g.history = append(g.history, move{stone: stone, pointIdx: g.pointIdxPlus1Map[pointName] - 1})

	idx := g.pointIdxPlus1Map[pointName] - 1
	curPtState := g.position[idx]
	if curPtState != unset {
		return errors.New(fmt.Sprintf("%s cannot move to %s: space is occupied", nameForPointState(stone), pointName))
	}
	newPosition := make([]pointState, len(g.position))
	copy(newPosition, g.position)
	newPosition[idx] = stone
	for _, adjpt := range g.board[idx].adjpts {
		adjptState := newPosition[adjpt.index]
		if adjptState != unset && adjptState != stone && adjpt.chainSurrounded(newPosition) {
			adjpt.removeChain(newPosition)
		}
	}
	if g.board[idx].chainSurrounded(newPosition) {
		return errors.New(fmt.Sprintf("%s cannot move to %s: suicide rule", nameForPointState(stone), pointName))
	}
	if reflect.DeepEqual(g.lastPosition, newPosition) {
		return errors.New(fmt.Sprintf("%s cannot move to %s: ko rule", nameForPointState(stone), pointName))
	}
	g.lastPosition = g.position
	g.position = newPosition
	return nil
}

func (g *game) String() string {
	return g.formatter()
}
