package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/FlindianaJones/boggle/boggle"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select Dice Set: (C)lassic, (N)ew, (R)andom")
	mode, _ := reader.ReadString('\n')
	mode = strings.ToUpper(mode[:len(mode)-1])
	//R or Random case
	gameBoard := boggle.CreateBoggle("RANDOM", 4)
	if mode == "C" || mode == "CLASSIC" {
		gameBoard = boggle.CreateBoggle("CLASSIC", 4)
	}
	if mode == "N" || mode == "NEW" {
		gameBoard = boggle.CreateBoggle("NEW", 4)
	}

	fmt.Println(gameBoard.PrettyPrintBoard())
	fmt.Print("Enter text: ")
	guesses := []string{}

	for len(guesses) < 1 {
		text, _ := reader.ReadString('\n')
		guesses = append(guesses, text[:len(text)-1])
	}

	fmt.Printf("your final score was %d\n", boggle.ScoreWords(gameBoard, guesses))

}
