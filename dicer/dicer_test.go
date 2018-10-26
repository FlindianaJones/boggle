package dicer

import (
	"reflect"
	"strings"
	"testing"
)

func TestClassicGenLetter(t *testing.T) {
	gen := PresetDicer{SourceDice: ClassicDice}

	got := gen.GenLetter()

	foundLetter := false
	for i := 0; i < 16 || !foundLetter; i++ {
		if strings.Contains(ClassicDice[i], string(got)) {
			foundLetter = true
		}
	}

	if !foundLetter {
		t.Errorf("Letter %s was not found on any dice!", string(got))
	}
}

func TestNewGenLetter(t *testing.T) {
	gen := PresetDicer{SourceDice: NewDice}

	got := gen.GenLetter()

	foundLetter := false
	for i := 0; i < 16 || !foundLetter; i++ {
		if strings.Contains(NewDice[i], string(got)) {
			foundLetter = true
		}
	}

	if !foundLetter {
		t.Errorf("Letter %s was not found on any dice!", string(got))
	}
}

func TestScrambleDice(t *testing.T) {
	og := make([]string, 16)
	copy(og, ClassicDice)

	gen := PresetDicer{SourceDice: ClassicDice}
	gen.scrambleDice()

	if !reflect.DeepEqual(og, ClassicDice) {
		t.Errorf("Mutated the source dice list!")
	}

	if reflect.DeepEqual(og, gen.scrambledDice) {
		t.Errorf("Did not randomize the source list of %v", gen.SourceDice)
	}
}
