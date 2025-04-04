package lang

import "fmt"

type ProtoSyllable struct {
	Onset    Consonant
	Nucleus  VowelFrontness
	Offglide VowelFrontness
	Coda     MannerOfArticulation
}

type ProtoWord struct {
	atr       bool
	syllables []ProtoSyllable
}

func NewProtoWord(atr bool, syllables []ProtoSyllable) (ProtoWord, error) {
	word := ProtoWord{
		atr:       atr,
		syllables: syllables,
	}
	var err error

	for i, syllable := range syllables {
		if syllable.Onset == NG {
			err = fmt.Errorf("ng is not a valid onset for proto words")
		} else if syllable.Onset == H {
			err = fmt.Errorf("h is not a valid onset for proto words")
		} else if syllable.Coda == Approximant {
			err = fmt.Errorf("approximant is not a valid MOA for proto coda")
			// } else if syllable.Coda == Lateral {
			// 	err = fmt.Errorf("lateral is not a valid MOA for proto coda")
		} else if i > 0 && (syllable.Onset == Consonant{}) {
			err = fmt.Errorf("non-initial syllables cannot have null onset")
		} else if syllable.Nucleus == 0 {
			err = fmt.Errorf("nucleus cannot have a null onset")
		}
	}

	return word, err
}

func surfaceProtoVowel(frontness VowelFrontness, atr bool) Vowel {
	height := Mid
	if atr {
		height++
	}
	if frontness == Central {
		height--
	}
	return Vowel{height, frontness}
}

func surfaceProtoCoda(moa MannerOfArticulation) Consonant {
	return Consonant{Coronal, moa}
}

func lenite(frontness VowelFrontness, coda MannerOfArticulation, onset Consonant) (MannerOfArticulation, Consonant) {
	switch coda {
	case 0:
		switch onset {
		case P:
			return 0, V
		case T:
			return 0, R
		case K:
			if frontness == Back {
				return 0, V
			} else {
				return 0, Y
			}
		}
	case Nasal:
		switch onset {
		case P:
			return Nasal, M
		case T:
			return Nasal, N
		case K:
			return Nasal, Consonant{Velar, Nasal}
		}
	case Stop:
		switch onset {
		case P:
			return 0, P
		case T:
			return 0, T
		case K:
			return 0, K
		}
	}

	return coda, onset
}

func (w ProtoWord) Romanization() string {
	s := ""
	for _, syllable := range w.syllables {
		s += syllable.Onset.IPA()
		s += surfaceProtoVowel(syllable.Nucleus, w.atr).IPA()
		if syllable.Offglide != 0 {
			s += surfaceProtoVowel(syllable.Offglide, w.atr).IPA()
		}
		if syllable.Coda != 0 {
			s += surfaceProtoCoda(syllable.Coda).IPA()
		}
	}
	return s
}
