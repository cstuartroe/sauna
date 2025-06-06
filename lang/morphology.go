package lang

import "strings"

type ProtoWord struct {
	stem     ProtoForm
	suffixes []ProtoSuffix
}

func NewProtoWord(stem ProtoForm) ProtoWord {
	return ProtoWord{
		stem: stem,
	}
}

func (pw *ProtoWord) AddSuffix(suffix ProtoSuffix) {
	pw.suffixes = append(pw.suffixes, suffix)
}

func (pw ProtoWord) Form() ProtoForm {
	form := pw.stem
	for _, suffix := range pw.suffixes {
		form = applySuffix(form, suffix)
	}
	return form
}

type ProtoSuffix struct {
	lenite        bool
	leadingCoda   MannerOfArticulation
	syllables     []ProtoSyllable
	syllablesFunc func(ProtoForm) []ProtoSyllable
}

func (s ProtoSuffix) Lenite() bool { return s.lenite }

func (s ProtoSuffix) Romanization() string {
	dummyWord := NewProtoWord(
		ProtoForm{
			false,
			[]ProtoSyllable{
				{
					Nucleus: Central,
				},
			},
		},
	)
	dummyWord.AddSuffix(s)

	suffixWord := Evolve(dummyWord.Form())
	rom := suffixWord.Romanization()[1:]
	rom = strings.ReplaceAll(rom, "a", "A")
	rom = strings.ReplaceAll(rom, "e", "I")
	rom = strings.ReplaceAll(rom, "i", "I")
	rom = strings.ReplaceAll(rom, "o", "U")
	rom = strings.ReplaceAll(rom, "u", "U")
	if rom == "" {
		rom = "∅"
	}
	if s.Lenite() {
		rom = "'" + rom
	}

	return rom
}

var epentheticVowel VowelFrontness = Front

func applySuffix(form ProtoForm, suffix ProtoSuffix) ProtoForm {
	newSyllables := []ProtoSyllable{}
	newSyllables = append(newSyllables, form.syllables...)
	lenitionIndex := -1

	lastSyllable := &newSyllables[len(newSyllables)-1]
	suffixSyllables := suffix.syllables
	if suffix.syllablesFunc != nil {
		suffixSyllables = suffix.syllablesFunc(form)
	}

	suffixStartsWithVowel := len(suffixSyllables) > 0 && suffixSyllables[0].Onset == Consonant{}

	if lastSyllable.Coda != 0 {
		if suffix.leadingCoda != 0 {
			lastSyllable.Coda = 0
			lenitionIndex = len(newSyllables)

			syll := ProtoSyllable{
				Onset:   surfaceProtoCoda(form.syllables[len(form.syllables)-1].Coda),
				Nucleus: epentheticVowel,
				Coda:    suffix.leadingCoda,
			}
			newSyllables = append(newSyllables, syll)
			newSyllables = append(newSyllables, suffixSyllables...)
		} else if suffixStartsWithVowel {
			lastSyllable.Coda = 0
			lenitionIndex = len(newSyllables)

			syllsToAdd := []ProtoSyllable{}
			syllsToAdd = append(syllsToAdd, suffixSyllables...)
			syllsToAdd[0].Onset = surfaceProtoCoda(form.syllables[len(form.syllables)-1].Coda)
			newSyllables = append(newSyllables, syllsToAdd...)
		} else {
			newSyllables = append(newSyllables, suffixSyllables...)
		}
	} else {
		lenitionIndex = len(newSyllables) - 1

		if suffix.leadingCoda != 0 {
			lastSyllable.Coda = suffix.leadingCoda
			newSyllables = append(newSyllables, suffixSyllables...)
		} else if suffixStartsWithVowel {
			if lastSyllable.Offglide == 0 && suffixSyllables[0].Offglide == 0 {
				lastSyllable.Offglide = suffixSyllables[0].Nucleus
				lastSyllable.Coda = suffixSyllables[0].Coda
				newSyllables = append(newSyllables, suffixSyllables[1:]...)
			} else {
				lastVowel := lastSyllable.Nucleus
				if lastSyllable.Offglide != 0 {
					lastVowel = lastSyllable.Offglide
				}
				epentheticConsonant := Y
				if lastVowel == Back {
					epentheticConsonant = V
				}
				newSyllables = append(newSyllables, suffixSyllables...)
				newSyllables[lenitionIndex+1].Onset = epentheticConsonant
			}
		} else {
			newSyllables = append(newSyllables, suffixSyllables...)
		}
	}

	if suffix.lenite && lenitionIndex != -1 {
		if lenitionIndex == 0 {
			// _, newSyllables[lenitionIndex].Onset = lenite(ProtoSyllable{}, newSyllables[lenitionIndex].Onset)
		} else {
			newSyllables[lenitionIndex-1].Coda, newSyllables[lenitionIndex].Onset = lenite(newSyllables[lenitionIndex-1], newSyllables[lenitionIndex].Onset)
		}
	}

	return ProtoForm{
		atr:       form.atr,
		syllables: newSyllables,
	}
}

var testSuffixes = []ProtoSuffix{
	{
		syllables: []ProtoSyllable{
			{
				Onset:   T,
				Nucleus: Central,
			},
		},
	},
	{
		syllables: []ProtoSyllable{
			{
				Onset:   P,
				Nucleus: Central,
			},
		},
	},
	{
		syllables: []ProtoSyllable{
			{
				Onset:   V,
				Nucleus: Central,
			},
		},
	},
	{
		syllables: []ProtoSyllable{
			{
				Onset:   N,
				Nucleus: Central,
			},
		},
	},
	{
		syllables: []ProtoSyllable{
			{
				Onset:   R,
				Nucleus: Central,
			},
		},
	},
	{
		syllables: []ProtoSyllable{
			{
				Nucleus: Central,
			},
		},
	},
	{
		syllables: []ProtoSyllable{
			{
				Nucleus:  Central,
				Offglide: Front,
			},
		},
	},
	{
		leadingCoda: Nasal,
		syllables: []ProtoSyllable{
			{
				Onset:   T,
				Nucleus: Central,
			},
		},
	},
	{
		leadingCoda: Stop,
		syllables: []ProtoSyllable{
			{
				Onset:   T,
				Nucleus: Central,
			},
		},
	},
}

var TestSuffixes []ProtoSuffix

func init() {
	for _, suffix := range testSuffixes {
		TestSuffixes = append(TestSuffixes, suffix)
		newSuffix := suffix
		newSuffix.lenite = true
		TestSuffixes = append(TestSuffixes, newSuffix)
	}
}

var NumberSuffixes = map[string]ProtoSuffix{
	"PL": {
		syllables: []ProtoSyllable{
			{
				Nucleus: Front,
			},
		},
	},
}

var PossessorSuffixes = map[string]ProtoSuffix{
	"P1S": {
		leadingCoda: Nasal,
	},
	"P2/3S": {
		leadingCoda: Stop,
	},
	"P1P": {
		leadingCoda: Nasal,
		syllables:   []ProtoSyllable{{M, Front, 0, Stop}},
	},
	"P2/3P": {
		syllables: []ProtoSyllable{{V, Front, 0, Stop}},
	},
}

var CaseSuffixes = map[string]ProtoSuffix{
	"TOP": {
		syllables: []ProtoSyllable{
			{
				Onset:   V,
				Nucleus: Central,
			},
		},
	},
	"NOM": {},
	"PAR": {
		leadingCoda: Stop,
	},
	"EQU": {
		leadingCoda: Fricative,
		syllables:   []ProtoSyllable{{S, Front, 0, 0}},
	},
	"DAT": {
		leadingCoda: Nasal,
		syllables: []ProtoSyllable{
			{
				Onset:   Y,
				Nucleus: Back,
			},
		},
	},
	"GEN": {
		lenite: true,
		syllables: []ProtoSyllable{
			{
				Nucleus: Central,
			},
		},
	},
	"ALL": {
		lenite: true,
		syllables: []ProtoSyllable{
			{
				Nucleus: Central,
			},
			{
				Onset:   N,
				Nucleus: Front,
			},
		},
	},
	"LOC": {
		lenite: true,
		syllables: []ProtoSyllable{
			{
				Nucleus: Central,
				Coda:    Stop,
			},
			{
				Onset:   T,
				Nucleus: Front,
			},
		},
	},
	"ABL": {
		lenite: true,
		syllables: []ProtoSyllable{
			{
				Nucleus: Central,
				Coda:    Fricative,
			},
			{
				Onset:   T,
				Nucleus: Front,
			},
		},
	},
}

var VoiceSuffixes = map[string]ProtoSuffix{
	"PSV": {
		syllables: []ProtoSyllable{
			{
				Onset:   R,
				Nucleus: Central,
			},
		},
	},
	"CAU": {
		syllables: []ProtoSyllable{
			{
				Onset:   S,
				Nucleus: Central,
			},
		},
	},
}

var NegationSuffixes = map[string]ProtoSuffix{
	"NEG": {
		syllables: []ProtoSyllable{
			{
				Onset:   N,
				Nucleus: Central,
			},
		},
	},
}

var TenseSuffixes = map[string]ProtoSuffix{
	"NPT": {
		syllables: []ProtoSyllable{
			{
				Onset:   R,
				Nucleus: Back,
			},
		},
	},
	"PST": {
		leadingCoda: Fricative,
	},
	"CNJ": {
		syllables: []ProtoSyllable{
			{
				Onset:   T,
				Nucleus: Front,
			},
		},
	},
}

var SubjectSuffixes = map[string]ProtoSuffix{
	"S1S": {
		leadingCoda: Nasal,
	},
	"S2/3S": {},
	"SP": {
		syllablesFunc: func(pw ProtoForm) []ProtoSyllable {
			lastSyllable := pw.syllables[len(pw.syllables)-1]
			var suffixSyllables = []ProtoSyllable{}
			if lastSyllable.Offglide == 0 && lastSyllable.Coda == 0 {
				// will lengthen final vowel
				suffixSyllables = append(suffixSyllables, ProtoSyllable{Nucleus: lastSyllable.Nucleus})
			}
			return suffixSyllables
		},
	},
}

var InterrogativeSuffixes = map[string]ProtoSuffix{
	"INT": {
		syllables: []ProtoSyllable{{K, Back, 0, 0}},
	},
}

var SuffixSets = []map[string]ProtoSuffix{
	NumberSuffixes,
	PossessorSuffixes,
	CaseSuffixes,
	VoiceSuffixes,
	NegationSuffixes,
	TenseSuffixes,
	SubjectSuffixes,
	InterrogativeSuffixes,
}
