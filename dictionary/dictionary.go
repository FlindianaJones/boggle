package dictionary

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Create a function that takes a list of words and returns only valid words
// Use the following web API to check if a word is valid
// API: method GET, path /dictionary, query word
const dictionaryURL = "https://whispering-falls-21983.herokuapp.com/"

func getValidWords(words []string) []string {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	doneChannel := make(chan string)

	for _, word := range words {
		go CheckWord(word, doneChannel)
	}

	responseList := []string{}
	responses := 0

	for {
		select {
		case v := <-doneChannel:
			responses++
			if v != "" {
				//fmt.Println("received word", v)
				responseList = append(responseList, v)
			}
			if responses == len(words) {
				return responseList
			}
		case <-ctx.Done():
			panic(errors.New("Overall request collection timed out"))
		}
	}
}

// CheckWord a single word against remote API
func CheckWord(word string, doneChannel chan string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	//fmt.Println("sending word", word)

	resp, err := http.DefaultClient.Get(fmt.Sprintf("%sdictionary?word=%s", dictionaryURL, word))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println("received word", word, "and body", string(body))
	if strings.Contains(string(body), "true") {
		//fmt.Println("Writing to the done channel")
		doneChannel <- word
	} else {
		doneChannel <- ""
	}

	for {
		select {
		case <-ctx.Done():
			if len(body) == 0 {
				panic(fmt.Errorf("HTTP Request timed out while checking %s", word))
			}
		}
	}
}
