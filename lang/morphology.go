package lang

type ProtoSuffix struct {
	lenite      bool
	leadingCoda MannerOfArticulation
	syllables   []ProtoSyllable
}

var epentheticVowel VowelFrontness = Front

func ApplySuffix(word ProtoWord, suffix ProtoSuffix) ProtoWord {
	newSyllables := []ProtoSyllable{}
	newSyllables = append(newSyllables, word.syllables...)
	lenitionIndex := -1

	lastSyllable := newSyllables[len(newSyllables)-1]

	if len(suffix.syllables) == 0 {
		lenitionIndex = len(newSyllables) - 1
	} else if lastSyllable.Coda != 0 {
		if suffix.leadingCoda != 0 {
			newSyllables[len(newSyllables)-1].Coda = 0
			lenitionIndex = len(newSyllables)

			syll := ProtoSyllable{
				Onset:   surfaceProtoCoda(word.syllables[len(word.syllables)-1].Coda),
				Nucleus: epentheticVowel,
				Coda:    suffix.leadingCoda,
			}
			newSyllables = append(newSyllables, syll)
			newSyllables = append(newSyllables, suffix.syllables...)
		} else if (suffix.syllables[0].Onset == Consonant{}) {
			newSyllables[len(newSyllables)-1].Coda = 0
			lenitionIndex = len(newSyllables)

			syllsToAdd := []ProtoSyllable{}
			syllsToAdd = append(syllsToAdd, suffix.syllables...)
			syllsToAdd[0].Onset = surfaceProtoCoda(word.syllables[len(word.syllables)-1].Coda)
			newSyllables = append(newSyllables, syllsToAdd...)
		} else {
			newSyllables = append(newSyllables, suffix.syllables...)
		}
	} else {
		lenitionIndex = len(newSyllables) - 1

		if suffix.leadingCoda != 0 {
			newSyllables[len(newSyllables)-1].Coda = suffix.leadingCoda
			newSyllables = append(newSyllables, suffix.syllables...)
		} else if (suffix.syllables[0].Onset == Consonant{}) {
			if lastSyllable.Offglide == 0 && suffix.syllables[0].Offglide == 0 {
				newSyllables[len(newSyllables)-1].Offglide = suffix.syllables[0].Nucleus
				newSyllables[len(newSyllables)-1].Coda = suffix.syllables[0].Coda
				newSyllables = append(newSyllables, suffix.syllables[1:]...)
			} else {
				lastVowel := lastSyllable.Nucleus
				if lastSyllable.Offglide != 0 {
					lastVowel = lastSyllable.Offglide
				}
				epentheticConsonant := Y
				if lastVowel == Back {
					epentheticConsonant = V
				}
				newSyllables = append(newSyllables, suffix.syllables...)
				newSyllables[lenitionIndex+1].Onset = epentheticConsonant
			}
		} else {
			newSyllables = append(newSyllables, suffix.syllables...)
		}
	}

	if suffix.lenite && lenitionIndex != -1 {
		if lenitionIndex == 0 {
			_, newSyllables[lenitionIndex].Onset = lenite(0, 0, newSyllables[lenitionIndex].Onset)
		} else {
			vowelFrontness := newSyllables[lenitionIndex-1].Nucleus
			if newSyllables[lenitionIndex-1].Offglide != 0 {
				vowelFrontness = newSyllables[lenitionIndex-1].Offglide
			}
			newSyllables[lenitionIndex-1].Coda, newSyllables[lenitionIndex].Onset = lenite(vowelFrontness, newSyllables[lenitionIndex-1].Coda, newSyllables[lenitionIndex].Onset)
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
	"ben": {
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
