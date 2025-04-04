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

func wordFromRomanization(rom string) (lang.ProtoWord, error) {
	syllables := []lang.ProtoSyllable{}
	var atr *bool

	i := 0
	var mostRecentConsonant lang.Consonant
	var currentSyllable lang.ProtoSyllable

	for i < len(rom) {
		letter := rom[i]

		if vowel, ok := vowels[letter]; ok {
			if (mostRecentConsonant != lang.Consonant{}) {
				syllables = append(syllables, currentSyllable)
				currentSyllable = lang.ProtoSyllable{
					Onset: mostRecentConsonant,
				}
				mostRecentConsonant = lang.Consonant{}
			}

			if atr == nil {
				b := lang.IsAtr(vowel)
				atr = &b
			} else if *atr != lang.IsAtr(vowel) {
				return lang.ProtoWord{}, fmt.Errorf("mismatching atr: %q", rom)
			}

			if currentSyllable.Nucleus == 0 {
				currentSyllable.Nucleus = vowel.Frontness
			} else if currentSyllable.Offglide == 0 {
				currentSyllable.Offglide = vowel.Frontness
			} else {
				return lang.ProtoWord{}, fmt.Errorf("three vowels in a row: %q", rom)
			}
		} else if cons, ok := consonants[letter]; ok {
			if (mostRecentConsonant != lang.Consonant{}) {
				if currentSyllable.Coda != 0 {
					return lang.ProtoWord{}, fmt.Errorf("three consonants in a row: %q", rom)
				}
				if !lang.IsValidCoda(mostRecentConsonant) {
					return lang.ProtoWord{}, fmt.Errorf("invalid coda consonant: %q of %q", letter, rom)
				}
				currentSyllable.Coda = cons.MOA
				syllables = append(syllables, currentSyllable)
				currentSyllable = lang.ProtoSyllable{}
			}
			mostRecentConsonant = cons
		} else {
			return lang.ProtoWord{}, fmt.Errorf("unknown letter %q of %q", letter, rom)
		}

		i++
	}

	if (mostRecentConsonant != lang.Consonant{}) {
		return lang.ProtoWord{}, fmt.Errorf("trailing consonant: %q", rom)
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
