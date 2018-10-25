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
	expectedBoard := Board{{'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}}

	if !reflect.DeepEqual(gotBoard, expectedBoard) {
		t.Errorf("Should produce a board that is all a's. Got %v but expected %v", gotBoard, expectedBoard)
	}
}

func TestWordScore(t *testing.T) {
	scores := map[int]int{3: 1, 4: 1, 5: 2, 6: 3, 7: 5, 8: 11, 9: 11, 12: 11}

	words := []string{"bud", "dude", "hairy", "swampy", "elegant", "corduroy", "gorblimey", "antediluvian"}

	for _, word := range words {
		got := GetWordScore(word)
		expected := scores[len(word)]
		if expected != got {
			t.Errorf("Incorrect score for %s: expected %d but got %d", word, expected, got)
		}
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
	//Graph is a "master node" that has links to every other node
	gotGraph := GetBoardGraph(testBoard)
	n1 := Node{label: 'a', used: false}
	n2 := Node{label: 'a', used: false}
	n3 := Node{label: 'a', used: false}
	n4 := Node{label: 'a', used: false}

	n1.links = []*Node{&n2, &n3, &n4}
	n2.links = []*Node{&n1, &n3, &n4}
	n3.links = []*Node{&n1, &n2, &n4}
	n4.links = []*Node{&n1, &n2, &n3}

	expectedGraph := Node{
		label: '\x00',
		used:  false,
		links: []*Node{&n1, &n2, &n3, &n4}}

	if len(gotGraph.links) != boardSize*boardSize {
		t.Errorf("Not enough nodes returned! Expected %d, got %d", boardSize*boardSize, len(gotGraph.links))
	}
	if !cmpNode(gotGraph, expectedGraph, false) {
		t.Errorf("Bad graph returned! Expected %+v, got %+v", expectedGraph, gotGraph)
	}
}

func cmpNode(a, b Node, dontRecurse bool) bool {
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
	board := Board{{'a', 'a'}, {'a', 'a'}}

	printyBoard := GetPrintableBoard(board)
	expectedPrintyBoard := "a a\na a\n"

	if printyBoard != expectedPrintyBoard {
		t.Errorf("Yon printy board is wronge. Expected \n%s\n, but got \n%s", expectedPrintyBoard, printyBoard)
	}
}

func TestWordInBoard(t *testing.T) {
	testBoard := Board{{'a', 'l', 'm', 's'}, {'b', 'r', 'o', 'w'}, {'c', 'l', 'a', 'm'}, {'d', 'y', 'e', 's'}}
	testWords := []string{"alms", "brow", "clam", "dyes", "loam", "mom", "lame", "solar", "world", "morl", "sos", "zither", "qualm"}
	testResults := []bool{true, true, true, true, true, true, true, true, true, true, false, false, false}

	boggle := Boggle{board: testBoard, graph: GetBoardGraph(testBoard)}
	for i, word := range testWords {
		got := WordInBoard(boggle, word)
		expected := testResults[i]
		if got != expected {
			t.Errorf("Incorrectly identified %s as being in board \n%s", word, GetPrintableBoard(testBoard))
		}
	}
}

func TestWordListScore(t *testing.T) {
	testBoard := Board{{'a', 'l', 'm', 's'}, {'b', 'r', 'o', 'w'}, {'c', 'l', 'a', 'm'}, {'d', 'y', 'e', 's'}}
	boggle := Boggle{board: testBoard, graph: GetBoardGraph(testBoard)}
	testWords := []string{"alms", "brow", "clam", "dyes", "loam", "mom", "lame", "solar", "world", "morl", "sos", "zither", "qualm", "balms"}
	expectedScore := 1 + 1 + 1 + 1 + 1 + 1 + 1 + 2 + 2 + 1 + 0 + 0 + 0 + 2

	gotScore := ScoreWords(boggle, testWords)

	if gotScore != expectedScore {
		t.Errorf("Incorrectly scored wordlist %v; expected %d, got %d", testWords, expectedScore, gotScore)
	}
}
