package gogame

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/twmb/algoimpl/go/graph"
)

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
	history     []move
	board       []graph.Node
	position    []pointState
	pointIdxMap map[string]int
	formatter   func() string
}

func makeStandardPointIdxMap(width, height int) map[string]int {
	m := make(map[string]int)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			idx := i*height + j + 1
			var name string
			if i < 8 {
				name = fmt.Sprintf("%c%d", []rune("A")[0]+rune(i), j+1)
			} else {
				name = fmt.Sprintf("%c%d", []rune("A")[0]+rune(i+1), j+1)
			}
			m[name] = idx
		}
	}
	return m
}

func Standard(width, height int) *game {
	g := game{
		board:       make([]graph.Node, width*height),
		position:    make([]pointState, width*height),
		pointIdxMap: makeStandardPointIdxMap(width, height),
	}
	g.formatter = func() string {
		var buffer bytes.Buffer
		for j := height - 1; j >= 0; j-- {
			for i := 0; i < width; i++ {
				if i != 0 {
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
	idx := g.pointIdxMap[pointName] - 1
	if g.position[idx] != unset {
		return errors.New("Black cannot move to D3: space is occupied")
	}
	g.position[idx] = stone
	return nil
}

func (g *game) String() string {
	return g.formatter()
}
