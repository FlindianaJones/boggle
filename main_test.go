package main

import "testing"

type MockReader struct {
	words []string
	count int
}

func (mr *MockReader) ReadString(b byte) (string, error) {
	if len(mr.words) == 0 {
		mr.words = []string{"aardwolf\n", "sdfdhfdsh\n", "test\n", "butts\n", "absentee\n"}
	}
	ret := mr.words[mr.count]
	mr.count++
	return ret, nil
}

func TestWordEntry(t *testing.T) {
	testReader := MockReader{}
	result := make(chan string)

	go EnterWord(&testReader, result)

	got := <-result

	if got != "aardwolf" {
		t.Error("Why the aardwolf hate?")
	}

	go EnterWord(&testReader, result)

	got = <-result

	if got != "" {
		t.Errorf("Good job validating gibberish: expected empty string, got %s", got)
	}
}
