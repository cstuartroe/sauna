package lang

type ProtoSuffix struct {
	lenite        bool
	leadingCoda   MannerOfArticulation
	syllables     []ProtoSyllable
	syllablesFunc func(ProtoWord) []ProtoSyllable
}

var epentheticVowel VowelFrontness = Front

func ApplySuffix(word ProtoWord, suffix ProtoSuffix) ProtoWord {
	newSyllables := []ProtoSyllable{}
	newSyllables = append(newSyllables, word.syllables...)
	lenitionIndex := -1

	lastSyllable := &newSyllables[len(newSyllables)-1]
	suffixSyllables := suffix.syllables
	if suffix.syllablesFunc != nil {
		suffixSyllables = suffix.syllablesFunc(word)
	}

	suffixStartsWithVowel := len(suffixSyllables) > 0 && suffixSyllables[0].Onset == Consonant{}

	if lastSyllable.Coda != 0 {
		if suffix.leadingCoda != 0 {
			lastSyllable.Coda = 0
			lenitionIndex = len(newSyllables)

			syll := ProtoSyllable{
				Onset:   surfaceProtoCoda(word.syllables[len(word.syllables)-1].Coda),
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
			syllsToAdd[0].Onset = surfaceProtoCoda(word.syllables[len(word.syllables)-1].Coda)
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

	return ProtoWord{
		atr:       word.atr,
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
	"pl": {
		syllables: []ProtoSyllable{
			{
				Nucleus: Front,
			},
		},
	},
}

var CaseSuffixes = map[string]ProtoSuffix{
	"top": {
		syllables: []ProtoSyllable{
			{
				Onset:   V,
				Nucleus: Central,
			},
		},
	},
	"nom": {},
	"par": {
		leadingCoda: Stop,
	},
	"ben": {
		leadingCoda: Nasal,
		syllables: []ProtoSyllable{
			{
				Onset:   Y,
				Nucleus: Back,
			},
		},
	},
	"gen": {
		lenite: true,
		syllables: []ProtoSyllable{
			{
				Nucleus: Central,
			},
		},
	},
	"all": {
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
	"loc": {
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
	"abl": {
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
	"psv": {
		syllables: []ProtoSyllable{
			{
				Onset:   R,
				Nucleus: Central,
			},
		},
	},
	"cau": {
		syllables: []ProtoSyllable{
			{
				Onset:   S,
				Nucleus: Central,
			},
		},
	},
}

var TenseSuffixes = map[string]ProtoSuffix{
	"prs": {
		syllables: []ProtoSyllable{
			{
				Onset:   R,
				Nucleus: Back,
			},
		},
	},
	"pst": {
		leadingCoda: Fricative,
	},
	"cnt": {
		syllables: []ProtoSyllable{
			{
				Onset:   T,
				Nucleus: Front,
			},
		},
	},
}

var SubjectSuffixes = map[string]ProtoSuffix{
	"s1s": {
		leadingCoda: Nasal,
	},
	"spl": {
		syllablesFunc: func(pw ProtoWord) []ProtoSyllable {
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

var SuffixSets = []map[string]ProtoSuffix{
	NumberSuffixes,
	CaseSuffixes,
	VoiceSuffixes,
	TenseSuffixes,
	SubjectSuffixes,
}
