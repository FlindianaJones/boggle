package dicer

import (
	"math/rand"
	"time"
)

// Generator is an interface for injecting into board creation
type Generator interface {
	GenLetter() rune
}

// PresetDicer generates random letters from the set of classic dice
type PresetDicer struct {
	SourceDice    []string
	scrambledDice []string
	calls         int32
}

// GenLetter returns a single rune from a random dice; randomizes dice if they haven't been yet
func (pd *PresetDicer) GenLetter() rune {
	if len(pd.scrambledDice) == 0 {
		rand.Seed(time.Now().UnixNano())
		pd.scrambleDice()
	}
	index := rand.Intn(6)
	pd.calls++
	//fmt.Printf("Getting letter %d from die %s, side %d\n", pd.calls, pd.scrambledDice[pd.calls%16], index)
	return rune(pd.scrambledDice[pd.calls%16][index])
}

func (pd *PresetDicer) scrambleDice() {
	cachedSource := make([]string, len(pd.SourceDice))
	copy(cachedSource, pd.SourceDice)
	for i := 16; i > 0; i-- {
		grabIndex := 0
		//can't pass a zero to rand.Intn, so have to protect it
		if i > 1 {
			grabIndex = rand.Intn(i - 1)
		}
		pd.scrambledDice = append(pd.scrambledDice, cachedSource[grabIndex])
		if i > 1 {
			cachedSource = append(cachedSource[:grabIndex], cachedSource[grabIndex+1:]...)
		}
	}
}

// ClassicDice is the set of classic die strings
var ClassicDice = []string{
	"AACIOT",
	"ABILTY",
	"ABJMOQu",
	"ACDEMP",
	"ACELRS",
	"ADENVZ",
	"AHMORS",
	"BIFORX",
	"DENOSW",
	"DKNOTU",
	"EEFHIY",
	"EGKLUY",
	"EGINTV",
	"EHINPS",
	"ELPSTU",
	"GILRUW"}

// NewDice is the set of new dice strings
var NewDice = []string{
	"AAEEGN",
	"ABBJOO",
	"ACHOPS",
	"AFFKPS",
	"AOOTTW",
	"CIMOTU",
	"DEILRX",
	"DELRVY",
	"DISTTY",
	"EEGHNW",
	"EEINSU",
	"EHRTVW",
	"EIOSST",
	"ELRTTY",
	"HIMNUQu",
	"HLNNRZ"}
