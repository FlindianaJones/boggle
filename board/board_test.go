package board

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestRandomLetterGenerator(t *testing.T) {
	letters := [100]rune{}
	gen := RandomLetter{}
	for i := 0; i < len(letters); i++ {
		letters[i] = gen.GenLetter()
	}

	var last rune
	hasDifferences := false
	var nonLetter rune
	for _, next := range letters {
		if last == '\x00' {
			last = next
			continue
		}
		if last != next {
			hasDifferences = true
		}
		if next < 'a' || next > 'z' {
			nonLetter = next
		}
	}

	if nonLetter != '\x00' {
		t.Errorf("Should produce a letter rune - int val from 96 to 122 inclusive, got %v", nonLetter)
	}
	if !hasDifferences {
		t.Errorf("Should produce something other than %v of %s", len(letters), strconv.QuoteRune(letters[0]))
	}
}

func BenchmarkGenLetter(b *testing.B) {
	gen := RandomLetter{}
	for i := 0; i < b.N; i++ {
		gen.GenLetter()
	}

}

func BenchmarkBoardGenerator(b *testing.B) {
	gen := RandomLetter{}
	for i := 0; i < b.N; i++ {
		GenerateBoard(gen, 4)
	}
}

type MockLetterGen struct{}

func (board MockLetterGen) GenLetter() rune {
	return 'a'
}

func TestBoardGenerator(t *testing.T) {
	boardSize := 4
	gotBoard := GenerateBoard(MockLetterGen{}, boardSize)
	expectedBoard := Board{Layout: [][]rune{{'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}}}

	if !reflect.DeepEqual(gotBoard, expectedBoard) {
		t.Errorf("Should produce a board that is all a's. Got %v but expected %v", gotBoard, expectedBoard)
	}
}

func TestGetNodeNeighbors(t *testing.T) {

	nodesToNeighbors := map[string][][]int{
		"0,0": [][]int{{0, 1}, {1, 0}, {1, 1}},
		"2,0": [][]int{{1, 0}, {1, 1}, {2, 1}, {3, 0}, {3, 1}},
		"1,1": [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}},
		"3,3": [][]int{{2, 2}, {2, 3}, {3, 2}},
	}

	for _, rc := range [][]int{{0, 0}, {2, 0}, {1, 1}, {3, 3}} {
		got := getNodeNeighborIndexes(rc[0], rc[1], 4)
		expected := nodesToNeighbors[fmt.Sprintf("%d,%d", rc[0], rc[1])]
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Bad neighbors for %v: expected %v, got %v", rc, expected, got)
		}
	}

}

func TestGetBoardGraphSize(t *testing.T) {
	boardSize := 2
	testBoard := GenerateBoard(MockLetterGen{}, boardSize)
	CreateBoardGraph(&testBoard)

	n1 := node{label: 'a', used: false}
	n2 := node{label: 'a', used: false}
	n3 := node{label: 'a', used: false}
	n4 := node{label: 'a', used: false}

	n1.links = []*node{&n2, &n3, &n4}
	n2.links = []*node{&n1, &n3, &n4}
	n3.links = []*node{&n1, &n2, &n4}
	n4.links = []*node{&n1, &n2, &n3}

	//Graph is a "master node" that has links to every other node
	expectedGraph := node{
		label: '\x00',
		used:  false,
		links: []*node{&n1, &n2, &n3, &n4}}

	if len(testBoard.graph.links) != boardSize*boardSize {
		t.Errorf("Not enough nodes returned! Expected %d, got %d", boardSize*boardSize, len(testBoard.graph.links))
	}
	if !cmpNode(testBoard.graph, expectedGraph, false) {
		t.Errorf("Bad graph returned! Expected %+v, got %+v", expectedGraph, testBoard.graph)
	}
}

func cmpNode(a, b node, dontRecurse bool) bool {
	if a.label == b.label && len(a.links) == len(b.links) {
		if !dontRecurse {
			for i := range a.links {
				if !cmpNode(*a.links[i], *b.links[i], true) {
					return false
				}
			}
		}
		return true
	}
	return false
}

func TestPrintBoard(t *testing.T) {
	board := Board{Layout: [][]rune{{'a', 'a'}, {'a', 'a'}}}

	printyBoard := GetPrintableBoard(board)
	expectedPrintyBoard := "a a\na a\n"

	if printyBoard != expectedPrintyBoard {
		t.Errorf("Yon printy board is wronge. Expected \n%s\n, but got \n%s", expectedPrintyBoard, printyBoard)
	}
}
