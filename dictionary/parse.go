package dictionary

import (
	"fmt"

	"github.com/cstuartroe/sauna/lang"
)

var consonants = map[byte]lang.Consonant{
	'm': lang.M,
	'n': lang.N,
	'g': lang.NG,

	'p': lang.P,
	't': lang.T,
	'k': lang.K,

	's': lang.S,

	'v': lang.V,
	'r': lang.R,
	'y': lang.Y,
}

var vowels = map[byte]lang.Vowel{
	'a': lang.A,
	'e': lang.E,
	'o': lang.O,

	'x': lang.SCHWA,
	'i': lang.I,
	'u': lang.U,
}

type parseState int

const (
	// Nothing has been parsed yet
	stateBegin parseState = 0

	// One consonant has been seen after the most recent vowel, and it's unclear whether it's a coda or onset
	stateAmbiguousConsonant parseState = 1

	// The most recent phoneme is a consonant that you know is the onset,
	// either because it's the first phoneme in the word, or because there was another consonant before it
	statePostOnset parseState = 2

	// One vowel has been seen since the most recent consonant/word start
	stateOneVowel parseState = 3

	// Two vowels have been seen since the most recent consonant/word start
	stateTwoVowels parseState = 4
)

func wordFromRomanization(rom string) (lang.ProtoWord, error) {
	syllables := []lang.ProtoSyllable{}
	var atr *bool

	i := 0
	state := stateBegin
	var mostRecentConsonant lang.Consonant
	var currentSyllable lang.ProtoSyllable

	for i < len(rom) {
		letter := rom[i]

		if vowel, ok := vowels[letter]; ok {
			if atr == nil {
				b := lang.IsAtr(vowel)
				atr = &b
			} else if *atr != lang.IsAtr(vowel) {
				return lang.ProtoWord{}, fmt.Errorf("mismatching atr: %q", rom)
			}

			switch state {
			case stateBegin, statePostOnset:
				currentSyllable.Nucleus = vowel.Frontness
				state = stateOneVowel

			case stateAmbiguousConsonant:
				syllables = append(syllables, currentSyllable)
				currentSyllable = lang.ProtoSyllable{
					Onset:   mostRecentConsonant,
					Nucleus: vowel.Frontness,
				}
				state = stateOneVowel

			case stateOneVowel:
				currentSyllable.Offglide = vowel.Frontness
				state = stateTwoVowels

			case stateTwoVowels:
				return lang.ProtoWord{}, fmt.Errorf("three vowels in a row: %q", rom)
			}
		} else if cons, ok := consonants[letter]; ok {
			switch state {
			case stateBegin:
				currentSyllable.Onset = cons
				state = statePostOnset
			case stateOneVowel, stateTwoVowels:
				state = stateAmbiguousConsonant
			case statePostOnset:
				return lang.ProtoWord{}, fmt.Errorf("too many consonants in a row: %q", rom)
			case stateAmbiguousConsonant:
				if !lang.IsValidCoda(mostRecentConsonant) {
					return lang.ProtoWord{}, fmt.Errorf("invalid coda consonant: %q of %q", letter, rom)
				}
				currentSyllable.Coda = mostRecentConsonant.MOA
				syllables = append(syllables, currentSyllable)
				currentSyllable = lang.ProtoSyllable{
					Onset: cons,
				}
				state = statePostOnset
			}

			mostRecentConsonant = cons
		} else {
			return lang.ProtoWord{}, fmt.Errorf("unknown letter %q of %q", letter, rom)
		}

		i++
	}

	switch state {
	case stateBegin:
		return lang.ProtoWord{}, fmt.Errorf("empty word")
	case stateAmbiguousConsonant:
		if !lang.IsValidCoda(mostRecentConsonant) {
			return lang.ProtoWord{}, fmt.Errorf("invalid final consonant: %q", rom)
		}
		currentSyllable.Coda = mostRecentConsonant.MOA
	case statePostOnset:
		return lang.ProtoWord{}, fmt.Errorf("word ends with two consonants: %q", rom)
	}

	syllables = append(syllables, currentSyllable)

	return lang.NewProtoWord(*atr, syllables)
}

var Dictionary = map[string]lang.ProtoWord{}

func init() {
	for name, rom := range _dictionary {
		word, err := wordFromRomanization(rom)
		if err != nil {
			panic(err)
		}
		Dictionary[name] = word
	}
}
