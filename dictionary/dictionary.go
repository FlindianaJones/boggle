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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	doneChannel := make(chan string)

	for _, word := range words {
		go checkWord(word, doneChannel)
	}

	responseList := []string{}
	responses := 0

	for {
		select {
		case v := <-doneChannel:
			responses++
			if v != "" {
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

func checkWord(word string, doneChannel chan string) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)

	defer cancel()

	resp, err := http.DefaultClient.Get(fmt.Sprintf("%sdictionary?word=%s", dictionaryURL, word))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if strings.Contains(string(body), "true") {
		doneChannel <- word
	} else {
		doneChannel <- ""
	}

	for {
		select {
		case <-ctx.Done():
			panic(errors.New("HTTP Request timed out"))
		}
	}
}
