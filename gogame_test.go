package gogame

import (
	"fmt"
	"testing"
)

func TestNewStandardGame(t *testing.T) {
	g := Standard(19, 19)
	if g == nil {
		t.Fail()
		return
	}
	if len(g.history) != 0 {
		t.Fail()
	}
	if len(g.board) != len(g.position) {
		t.Fail()
	}
	if len(g.board) != 19*19 {
		t.Fail()
	}
	if g.pointIdxPlus1Map == nil {
		t.Fail()
	}
	if g.pointIdxPlus1Map["C4"] != 42 {
		t.Fail()
	}
	if g.pointIdxPlus1Map["I9"] != 0 {
		t.Fail()
	}
}

func ExampleStandardGame() {
	g := Standard(5, 5)
	fmt.Println("\n", g)
	// Output:
	// · · · · ·
	// · · · · ·
	// · · · · ·
	// · · · · ·
	// · · · · ·
}

func ExampleMoves() {
	g := Standard(5, 5)
	g.Move(Black, "D3")
	g.Move(White, "C3")
	g.Move(Black, "D2")
	fmt.Println("\n", g)
	// Output:
	// · · · · ·
	// · · · · ·
	// · · ○ ● ·
	// · · · ● ·
	// · · · · ·
}

func ExampleOccupiedSpace() {
	g := Standard(5, 5)
	g.Move(Black, "D3")
	fmt.Println(g.Move(Black, "D3"))
	// Output: Black cannot move to D3: space is occupied
}

func ExampleSingleCapture() {
	g := Standard(5, 5)
	g.Move(Black, "D3")
	g.Move(Black, "C4")
	g.Move(Black, "B3")
	g.Move(White, "C3")
	g.Move(Black, "C2")
	fmt.Println("\n", g)
	// Output:
	// · · · · ·
	// · · ● · ·
	// · ● · ● ·
	// · · ● · ·
	// · · · · ·
}

func ExampleMultipleCapture() {
	g := Standard(5, 5)
	g.Move(Black, "A1")
	g.Move(Black, "B1")
	g.Move(Black, "A2")
	g.Move(White, "C1")
	g.Move(White, "B2")
	g.Move(White, "A3")
	fmt.Println("\n", g)
	// Output:
	// · · · · ·
	// · · · · ·
	// ○ · · · ·
	// · ○ · · ·
	// · · ○ · ·
}

func ExampleSuicideRule() {
	g := Standard(5, 5)
	g.Move(Black, "D3")
	g.Move(Black, "C4")
	g.Move(Black, "B3")
	g.Move(Black, "C2")
	fmt.Println(g.Move(White, "C3"))
	// Output: White cannot move to C3: suicide rule
}

func ExampleNearSuicideCapture() {
	g := Standard(5, 5)
	g.Move(Black, "D3")
	g.Move(Black, "C4")
	g.Move(Black, "B3")
	g.Move(Black, "C2")
	g.Move(White, "B4")
	g.Move(White, "A3")
	g.Move(White, "B2")
	g.Move(White, "C3")
	fmt.Println("\n", g)
	// Output:
	// · · · · ·
	// · ○ ● · ·
	// ○ · ○ ● ·
	// · ○ ● · ·
	// · · · · ·
}

func ExampleKoRule() {
	g := Standard(5, 5)
	g.Move(Black, "D3")
	g.Move(Black, "C4")
	g.Move(Black, "B3")
	g.Move(Black, "C2")
	g.Move(White, "B4")
	g.Move(White, "A3")
	g.Move(White, "B2")
	g.Move(White, "C3")
	fmt.Println(g.Move(Black, "B3"))
	// Output: Black cannot move to B3: ko rule
}
