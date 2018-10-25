package boggle

import (
	"testing"
)

type MockWordContainer struct {
	called int
}

func (c *MockWordContainer) ContainsWord(word string) bool {
	c.called++
	switch word {
	case "sos", "zither", "qualm":
		return false
	default:
		return true
	}
}

func TestWordInBoard(t *testing.T) {
	mockBoard := &MockWordContainer{}
	testWords := []string{"alms", "brow", "clam", "dyes", "loam", "mom", "lame", "solar", "world", "morl", "sos", "zither", "qualm"}

	for i, word := range testWords {
		WordInBoard(mockBoard, word)
		expected := i + 1
		if mockBoard.called != expected {
			t.Errorf("Didn't call ContainsWord on the board!")
		}
	}
}

func TestWordListScore(t *testing.T) {
	boggle := Boggle{board: &MockWordContainer{}}
	testWords := []string{"alms", "brow", "clam", "dyes", "loam", "mom", "lame", "solar", "world", "morl", "sos", "zither", "qualm", "balms"}
	expectedScore := 1 + 1 + 1 + 1 + 1 + 1 + 1 + 2 + 2 + 1 + 0 + 0 + 0 + 2

	gotScore := ScoreWords(boggle, testWords)

	if gotScore != expectedScore {
		t.Errorf("Incorrectly scored wordlist %v; expected %d, got %d", testWords, expectedScore, gotScore)
	}
}

func TestWordScore(t *testing.T) {
	scores := map[int]int{3: 1, 4: 1, 5: 2, 6: 3, 7: 5, 8: 11, 9: 11, 12: 11}

	words := []string{"bud", "dude", "hairy", "swampy", "elegant", "corduroy", "gorblimey", "antediluvian"}

	for _, word := range words {
		got := getWordScore(word)
		expected := scores[len(word)]
		if expected != got {
			t.Errorf("Incorrect score for %s: expected %d but got %d", word, expected, got)
		}
	}
}
