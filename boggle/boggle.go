package boggle

import (
	"strings"

	"github.com/FlindianaJones/boggle/board"
	"github.com/FlindianaJones/boggle/dicer"
)

// Boggle the mind
type Boggle struct {
	board board.WordContainer
}

// CreateBoggle creates a board and graph, with specified generator and size
func CreateBoggle(diceType string, size int) Boggle {
	t := strings.ToUpper(diceType)
	switch t {
	case "NEW":
		return Boggle{board: board.GenerateBoard(&dicer.PresetDicer{SourceDice: dicer.NewDice}, size)}
	case "CLASSIC":
		return Boggle{board: board.GenerateBoard(&dicer.PresetDicer{SourceDice: dicer.ClassicDice}, size)}
	default:
		return Boggle{board: board.GenerateBoard(board.RandomLetter{}, size)}
	}
}

// WordInBoard finds a word in the board, and returns false if it can't
func WordInBoard(board board.WordContainer, word string) bool {
	word = strings.ToUpper(word)
	return board.ContainsWord(word)
}

// PrettyPrintBoard returns a string representing an easily printable output of the board's runes
func (b Boggle) PrettyPrintBoard() string {
	return b.board.GetPrintableBoard()
}

// ScoreWords scores words if they exist in the board
func ScoreWords(boggle Boggle, words []string) int {
	score := 0
	for _, word := range words {
		if WordInBoard(boggle.board, word) {
			score += getWordScore(word)
		}
	}
	return score
}

func getWordScore(word string) int {
	length := len(word)
	switch {
	case length < 3:
		return 0
	case length < 5:
		return 1
	case length < 6:
		return 2
	case length < 7:
		return 3
	case length < 8:
		return 5
	default:
		return 11
	}
}
