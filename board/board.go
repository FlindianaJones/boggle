package board

import (
	"math/rand"
	"time"
)

// Board type def
type Board [][]rune

// Boggle the mind
type Boggle struct {
	board Board
	graph Node
}

// Generator is an interface for injecting into board creation
type Generator interface {
	GenLetter() rune
}

// Node is a node
type Node struct {
	label rune
	links []*Node
	used  bool
}

// RandomLetter generates a random array. JK - rune.
type RandomLetter struct{}

// GenLetter returns a single random letter
func (r RandomLetter) GenLetter() rune {
	rand.Seed(time.Now().UnixNano())
	char := rand.Intn(26) + 97
	return rune(char)
}

// GenerateBoard hooray!
func GenerateBoard(gen Generator, size int) Board {
	board := make([][]rune, size)
	for row := range board {
		board[row] = make([]rune, size)
		for col := range board[row] {
			board[row][col] = gen.GenLetter()
		}
	}
	return board
}

// GetWordScore implements simple Boggle scoring algorithm
func GetWordScore(word string) int {
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

// WordInBoard finds a word in the board, and returns false if it can't
func WordInBoard(board Boggle, word string) bool {
	return findInNode(&board.graph, word)
}

func findInNode(node *Node, word string) bool {
	node.used = true
	for _, link := range node.links {
		if !link.used && link.label == rune(word[0]) {
			if len(word) == 1 || findInNode(link, word[1:]) {
				node.used = false
				return true
			}
		}
	}
	node.used = false
	return false
}

// GetNodeNeighborIndexes returns all coordinates neighboring
func getNodeNeighborIndexes(row, col, size int) [][]int {
	//setup our iteration boundaries
	rowStart := intMax(row-1, 0)
	rowEnd := intMin(row+1, size-1)
	colStart := intMax(col-1, 0)
	colEnd := intMin(col+1, size-1)

	result := [][]int{}

	for r := rowStart; r <= rowEnd; r++ {
		for c := colStart; c <= colEnd; c++ {
			if r != row || c != col {
				result = append(result, []int{r, c})
			}
		}
	}
	return result
}

// GetBoardGraph gets a board graph
func GetBoardGraph(board Board) Node {
	size := len(board[0])
	rootNode := Node{
		label: '\x00',
		used:  false,
		links: []*Node{},
	}
	nodeBoard := make([][]Node, size)
	for r := 0; r < size; r++ {
		//need to pre-initialize this, because otherwise lookaheads will crash
		nodeBoard[r] = make([]Node, size)
	}
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			nodeBoard[r][c].label = board[r][c]

			for _, coords := range getNodeNeighborIndexes(r, c, size) {
				//fmt.Println(r, c, coords, nodeBoard)
				nodeBoard[r][c].links = append(nodeBoard[r][c].links, &nodeBoard[coords[0]][coords[1]])
			}
			rootNode.links = append(rootNode.links, &nodeBoard[r][c])
		}
	}
	return rootNode
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// GetPrintableBoard is a terrible name, it gets a printyBoard!
func GetPrintableBoard(board Board) string {
	output := ""
	size := len(board[0])
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			output += string(board[r][c]) + " "
		}
		output = output[:len(output)-1] + "\n"
	}
	return output
}

// ScoreWords scores words if they exist in the board
func ScoreWords(boggle Boggle, words []string) int {
	score := 0
	for _, word := range words {
		if WordInBoard(boggle, word) {
			score += GetWordScore(word)
		}
	}
	return score
}
