package board

import (
	"math/rand"
	"time"
)

// WordContainer is a container of words
type WordContainer interface {
	ContainsWord(string) bool
}

// Board type def
type Board struct {
	Layout [][]rune
	graph  node
}

// Generator is an interface for injecting into board creation
type Generator interface {
	GenLetter() rune
}

// node is a node
type node struct {
	label rune
	links []*node
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
	board := Board{Layout: make([][]rune, size)}
	for row := range board.Layout {
		board.Layout[row] = make([]rune, size)
		for col := range board.Layout[row] {
			board.Layout[row][col] = gen.GenLetter()
		}
	}
	return board
}

// ContainsWord if it returns a boolean true
func (b Board) ContainsWord(word string) bool {
	return b.graph.findInNode(word)
}

// FindInNode finds first of string in a linked node and rest of string in linked node's linked nodes if needed
func (node *node) findInNode(word string) bool {
	node.used = true
	for _, link := range node.links {
		if !link.used && link.label == rune(word[0]) {
			if len(word) == 1 || link.findInNode(word[1:]) {
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

// CreateBoardGraph gets a board graph
func CreateBoardGraph(board *Board) {
	size := len(board.Layout[0])
	rootNode := node{
		label: '\x00',
		used:  false,
		links: []*node{},
	}
	nodeBoard := make([][]node, size)
	for r := 0; r < size; r++ {
		//need to pre-initialize this, because otherwise lookaheads will crash
		nodeBoard[r] = make([]node, size)
	}
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			nodeBoard[r][c].label = board.Layout[r][c]

			for _, coords := range getNodeNeighborIndexes(r, c, size) {
				//fmt.Println(r, c, coords, nodeBoard)
				nodeBoard[r][c].links = append(nodeBoard[r][c].links, &nodeBoard[coords[0]][coords[1]])
			}
			rootNode.links = append(rootNode.links, &nodeBoard[r][c])
		}
	}
	board.graph = rootNode
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
	size := len(board.Layout[0])
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			output += string(board.Layout[r][c]) + " "
		}
		output = output[:len(output)-1] + "\n"
	}
	return output
}
