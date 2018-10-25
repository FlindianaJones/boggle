package dictionary

import (
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

const trueReturn = `{"wordFound":"true"}`
const falseReturn = `{"wordFound":"false"}`

func TestGetValidWords(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://whispering-falls-21983.herokuapp.com/dictionary?word=cat",
		httpmock.NewStringResponder(200, trueReturn))

	httpmock.RegisterResponder("GET", "https://whispering-falls-21983.herokuapp.com/dictionary?word=fergulous",
		httpmock.NewStringResponder(200, falseReturn))

	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test List Of One Valid One Invalid Words (mock)",
			args: args{
				words: []string{"cat", "fergulous"},
			},
			want: []string{"cat"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getValidWords(tt.args.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValidWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGetValidWordsAPI(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test List Of One Valid One Invalid Words (FOR REAL)",
			args: args{
				words: []string{"dog", "reponse"},
			},
			want: []string{"dog"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getValidWords(tt.args.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValidWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkWord(t *testing.T) {
	want := "test"

	testChan := make(chan string)

	go checkWord("test", testChan)

	got := <-testChan

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetValidWords() = %v, want %v", got, want)
	}
}
