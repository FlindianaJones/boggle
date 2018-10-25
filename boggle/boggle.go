package boggle

import "github.com/FlindianaJones/boggle/board"

// Boggle the mind
type Boggle struct {
	board board.WordContainer
}

// CreateBoggle creates a... uh... you know.
func CreateBoggle(size int) Boggle {
	return Boggle{board: board.GenerateBoard(board.RandomLetter{}, size)}
}

// WordInBoard finds a word in the board, and returns false if it can't
func WordInBoard(board board.WordContainer, word string) bool {
	return board.ContainsWord(word)
}

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
