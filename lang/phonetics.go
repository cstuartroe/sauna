package lang

import "fmt"

type Phoneme interface {
	IPA() string
}

type Syllable interface {
	Moras() int
}

type Word interface {
	Romanization() string
}

type VowelHeight int

const (
	Low  VowelHeight = 1
	Mid  VowelHeight = 2
	High VowelHeight = 3
)

type VowelFrontness int

const (
	Back    VowelFrontness = 1
	Central VowelFrontness = 2
	Front   VowelFrontness = 3
)

type Vowel struct {
	Height    VowelHeight
	Frontness VowelFrontness
}

func (v Vowel) IPA() string {
	switch v {
	case Vowel{Low, Central}:
		return "a"
	case Vowel{Mid, Central}:
		return "õ"
	case Vowel{Mid, Front}:
		return "e"
	case Vowel{High, Front}:
		return "i"
	case Vowel{Mid, Back}:
		return "o"
	case Vowel{High, Back}:
		return "u"
	default:
		return fmt.Sprintf("%+v", v)
	}
}

var (
	I     = Vowel{High, Front}
	E     = Vowel{Mid, Front}
	A     = Vowel{Low, Central}
	SCHWA = Vowel{Mid, Central}
	O     = Vowel{Mid, Back}
	U     = Vowel{High, Back}
)

type PlaceOfArticulation int

const (
	Labial  PlaceOfArticulation = 1
	Coronal PlaceOfArticulation = 2
	Palatal PlaceOfArticulation = 3
	Velar   PlaceOfArticulation = 4
)

type MannerOfArticulation int

const (
	Nasal       MannerOfArticulation = 1
	Stop        MannerOfArticulation = 2
	Fricative   MannerOfArticulation = 3
	Approximant MannerOfArticulation = 4
)

type Consonant struct {
	POA PlaceOfArticulation
	MOA MannerOfArticulation
}

func (c Consonant) IPA() string {
	switch c {
	case Consonant{}:
		return ""

	case Consonant{Labial, Nasal}:
		return "m"
	case Consonant{Coronal, Nasal}:
		return "n"
	case Consonant{Velar, Nasal}:
		return "g"

	case Consonant{Labial, Stop}:
		return "p"
	case Consonant{Coronal, Stop}:
		return "t"
	case Consonant{Velar, Stop}:
		return "k"

	case Consonant{Coronal, Fricative}:
		return "s"
	case Consonant{Velar, Fricative}:
		return "h"

	case Consonant{Labial, Approximant}:
		return "v"
	case Consonant{Coronal, Approximant}:
		return "r"
	case Consonant{Palatal, Approximant}:
		return "y"

	default:
		return fmt.Sprintf("%+v", c)
	}
}

var (
	M  = Consonant{Labial, Nasal}
	N  = Consonant{Coronal, Nasal}
	NG = Consonant{Velar, Nasal}

	P = Consonant{Labial, Stop}
	T = Consonant{Coronal, Stop}
	K = Consonant{Velar, Stop}

	S = Consonant{Coronal, Fricative}
	H = Consonant{Velar, Fricative}

	V = Consonant{Labial, Approximant}
	R = Consonant{Coronal, Approximant}
	Y = Consonant{Palatal, Approximant}
)
