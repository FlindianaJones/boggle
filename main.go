package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/FlindianaJones/boggle/boggle"
	"github.com/FlindianaJones/boggle/dictionary"
)

func main() {
	play := true
	reader := bufio.NewReader(os.Stdin)
	for play {
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
		guesses := []string{}
		checkChan := make(chan string)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		done := false

		go EnterWord(reader, checkChan)
		for !done {
			select {
			case resp := <-checkChan:
				if resp != "" {
					//fmt.Printf("Adding guess = %s to guesses\n", resp)
					guesses = append(guesses, resp)
				}
				go EnterWord(reader, checkChan)
			case <-ctx.Done():
				done = true
			}
		}

		fmt.Printf("your final score was %d\n", boggle.ScoreWords(gameBoard, guesses))
		fmt.Println("Play again? (y)es/(n)o")
		playAgain, err := reader.ReadString('\n')
		playAgain = strings.ToUpper(playAgain[:len(playAgain)-1])
		play = false
		if err != nil || playAgain == "Y" || playAgain == "YES" {
			play = true
		}
		if !play {
			fmt.Println("Okay, that's fine, I'm not upset at all that you said", playAgain)
			time.Sleep(5 * time.Second)
			fmt.Println(noose)
		}
	}
}

// EnterWord is intended to be used concurrently; checking online dictionary, writing to channel
func EnterWord(reader Reader, rChan chan string) {
	fmt.Print("Enter Word: ")
	enteredWord, err := reader.ReadString('\n')
	if err != nil {
		rChan <- ""
	}
	enteredWord = enteredWord[:len(enteredWord)-1]

	go dictionary.CheckWord(enteredWord, rChan)
}

// Reader provides a simple interface for testing with a mock CLI input
type Reader interface {
	ReadString(byte) (string, error)
}

const noose = "  +----+\n" +
	"  |    |\n" +
	"  O    |\n" +
	" /|\\   |\n" +
	" / \\   |\n" +
	"       |\n" +
	"=========''']\n"
