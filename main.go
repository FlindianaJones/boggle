package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/FlindianaJones/boggle/boggle"
)

func main() {
	gameBoard := boggle.CreateBoggle(4)

	fmt.Println(gameBoard.PrettyPrintBoard())

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")

	guesses := []string{}

	for len(guesses) < 10 {
		text, _ := reader.ReadString('\n')
		guesses = append(guesses, text[:len(text)-1])
	}

	fmt.Printf("your final score was %d\n", boggle.ScoreWords(gameBoard, guesses))

}
