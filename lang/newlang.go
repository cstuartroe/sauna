package lang

import "slices"

type NewSyllable struct {
	Onset    Consonant
	Nucleus  Vowel
	Offglide Vowel
	Coda     Consonant
}

type NewWord struct {
	syllables []NewSyllable
}

func Evolve(protoword ProtoWord) NewWord {
	newword := NewWord{}

	for i, syllable := range protoword.syllables {
		newsyllable := NewSyllable{
			Onset:   syllable.Onset,
			Nucleus: surfaceProtoVowel(syllable.Nucleus, protoword.atr),
		}

		var prevSyll *NewSyllable = &NewSyllable{}
		if i > 0 {
			prevSyll = &newword.syllables[i-1]
		}

		switch prevSyll.Coda {
		case Consonant{}:
			if newsyllable.Onset == P {
				newsyllable.Onset = H
			}
		case N:
			if newsyllable.Onset.POA == Labial {
				prevSyll.Coda.POA = Labial
			}
			if newsyllable.Onset == V {
				newsyllable.Onset = M
			}
		case T, S:
			if newsyllable.Onset.MOA == Nasal || newsyllable.Onset == R {
				prevSyll.Coda = R
			}
			if newsyllable.Onset == V {
				newsyllable.Onset = P
			}
			if newsyllable.Onset == Y {
				newsyllable.Onset = K
			}
			if prevSyll.Coda == T && slices.Contains([]MannerOfArticulation{Stop, Fricative}, newsyllable.Onset.MOA) {
				prevSyll.Coda = newsyllable.Onset
			}
		}

		if syllable.Offglide != 0 {
			newsyllable.Offglide = surfaceProtoVowel(syllable.Offglide, protoword.atr)
		}
		if syllable.Coda != 0 {
			newsyllable.Coda = surfaceProtoCoda(syllable.Coda)
		}

		if protoword.atr {
			if syllable.Offglide != 0 {
				if syllable.Nucleus == Central {
					newsyllable.Nucleus.Frontness = syllable.Offglide
				}
				if syllable.Offglide == Central {
					newsyllable.Offglide = Vowel{
						Height:    newsyllable.Nucleus.Height - 1,
						Frontness: newsyllable.Nucleus.Frontness,
					}
				}
			}
			if newsyllable.Nucleus == SCHWA {
				newsyllable.Nucleus = E
			}
		} else {
			if syllable.Offglide != 0 && syllable.Offglide != Central {
				newsyllable.Offglide.Height = High
			}
		}

		newword.syllables = append(newword.syllables, newsyllable)
	}

	return newword
}

func (w NewWord) Romanization() string {
	s := ""
	for _, syllable := range w.syllables {
		s += syllable.Onset.IPA()
		s += syllable.Nucleus.IPA()
		if (syllable.Offglide != Vowel{}) {
			s += syllable.Offglide.IPA()
		}
		s += syllable.Coda.IPA()
	}
	return s
}
