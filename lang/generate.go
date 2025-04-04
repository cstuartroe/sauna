package lang

import "math/rand"

func WeightedRandomElement[T comparable](weights map[T]float64) T {
	sumWeights := 0.0
	for _, weight := range weights {
		sumWeights += weight
	}
	threshold := rand.Float64() * sumWeights
	cumulative := 0.0
	for value, weight := range weights {
		cumulative += weight
		if cumulative > threshold {
			return value
		}
	}
	panic("Never reached threshold")
}

func UniformDropoff[T comparable](l []T, dropoff float64) map[T]float64 {
	weight := 1.0
	out := map[T]float64{}
	for _, t := range l {
		out[t] = weight
		weight *= dropoff
	}
	return out
}

const defaultDropoff = .8

var protoInitialConsonantWeights = UniformDropoff([]Consonant{{}, K, T, N, V, S /*L,*/, Y, M, P}, defaultDropoff)
var protoMedialConsonantWeights = UniformDropoff([]Consonant{K, T, N, V, S, R /*L,*/, Y, M, P}, defaultDropoff)
var protoNucleusWeights = UniformDropoff([]VowelFrontness{Central, Front, Back}, .6)
var protoOffglideWeights = map[VowelFrontness]float64{0: 20, Front: 1, Central: .75, Back: 1}
var protoCodaWeights = map[MannerOfArticulation]float64{0: 15, Nasal: 2, Stop: 1, Fricative: .75}
var protoSyllableLengthWeights = map[int]float64{1: 2, 2: 10, 3: 3, 4: 1}
var atrWeights = map[bool]float64{true: 1, false: 1}

func validProtoInitialConsonant(syllable ProtoSyllable) bool {
	if syllable.Nucleus == Front && syllable.Onset == Y {
		return false
	}
	if syllable.Nucleus == Back && syllable.Onset == V {
		return false
	}
	return true
}

func RandomProtoWord() (ProtoWord, error) {
	syllables := []ProtoSyllable{}
	numSyllables := WeightedRandomElement(protoSyllableLengthWeights)
	for i := 0; i < numSyllables; i++ {
		syllable := ProtoSyllable{}

		syllable.Nucleus = WeightedRandomElement(protoNucleusWeights)
		syllable.Offglide = WeightedRandomElement(protoOffglideWeights)

		validInitialConsonant := false
		for !validInitialConsonant {
			if i == 0 {
				syllable.Onset = WeightedRandomElement(protoInitialConsonantWeights)
			} else {
				syllable.Onset = WeightedRandomElement(protoMedialConsonantWeights)
			}
			validInitialConsonant = validProtoInitialConsonant(syllable)
		}

		if syllable.Offglide == 0 {
			syllable.Coda = WeightedRandomElement(protoCodaWeights)
		}

		syllables = append(syllables, syllable)
	}

	if len(syllables) >= 2 {
		if rand.Float64() > .5 {
			syllables[0].Coda, syllables[1].Onset = lenite(syllables[0], syllables[1].Onset)
		}
	}

	return NewProtoWord(
		WeightedRandomElement(atrWeights),
		syllables,
	)
}
