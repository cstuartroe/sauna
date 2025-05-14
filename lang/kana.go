package lang

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

var frontKana = map[Consonant]string{
	{}: "イ",
	K:  "キ",
	S:  "シ",
	T:  "チ",
	N:  "ニ",
	P:  "ヒ",
	M:  "ミ",
	Y:  "𛄠",
	R:  "リ",
	V:  "ヰ",
}

var backKana = map[Consonant]string{
	{}: "オ",
	K:  "コ",
	S:  "ソ",
	T:  "ト",
	N:  "ノ",
	P:  "ホ",
	M:  "モ",
	Y:  "ヨ",
	R:  "ロ",
	V:  "ヲ",
}

var lowKana = map[Consonant]string{
	{}: "ア",
	K:  "カ",
	S:  "サ",
	T:  "タ",
	N:  "ナ",
	P:  "ハ",
	M:  "マ",
	Y:  "ヤ",
	R:  "ラ",
	V:  "ワ",
}

var highKana = map[Consonant]string{
	{}: "エ",
	K:  "ケ",
	S:  "セ",
	T:  "テ",
	N:  "ネ",
	P:  "ヘ",
	M:  "メ",
	Y:  "𛄡",
	R:  "レ",
	V:  "ヱ",
}

var codaKana = map[MannerOfArticulation]string{
	Nasal:     "ン",
	Stop:      "ツ",
	Fricative: "ス",
}

var codaKanaToOnset = map[string]Consonant{
	codaKana[Nasal]:     N,
	codaKana[Stop]:      T,
	codaKana[Fricative]: S,
}

var codaKanaList = slices.Collect(maps.Values(codaKana))

var lenitedKanaMap = map[string]string{
	"カ": "ガ",
	"ケ": "ゲ",
	"キ": "ギ",
	"コ": "ゴ",

	"タ": "ダ",
	"テ": "デ",
	"チ": "ヂ",
	"ト": "ド",

	"ハ": "バ",
	"ヒ": "ビ",
	"フ": "ブ",
	"ヘ": "ベ",
	"ホ": "ボ",

	"ツ": "ヅ",
}

const longSign string = "ー"

func getCodaKana(coda MannerOfArticulation, nextOnset *Consonant) string {
	if nextOnset == nil {
		return codaKana[coda]
	}

	if *nextOnset == R && coda == Fricative {
		return codaKana[Stop]
	}

	return codaKana[coda]
}

func syllableToKana(syll ProtoSyllable, atr bool, nextOnset *Consonant) []string {
	out := []string{}
	var kanaMap map[Consonant]string
	if syll.Nucleus == Front {
		kanaMap = frontKana
	} else if syll.Nucleus == Back {
		kanaMap = backKana
	} else if atr {
		kanaMap = highKana
	} else {
		kanaMap = lowKana
	}

	onset := syll.Onset
	lenite := false
	if onset == NG {
		onset = K
		lenite = true
	}

	kana, ok := kanaMap[onset]
	if !ok {
		panic(fmt.Sprintf("Unknown kana combo: %s %d %t", syll.Onset.IPA(), syll.Nucleus, atr))
	}
	if lenite {
		lenitedKana, ok := lenitedKanaMap[kana]
		if ok {
			kana = lenitedKana
		}
	}
	out = append(out, kana)

	if syll.Offglide == 0 {
		// do nothing
	} else if syll.Offglide == syll.Nucleus {
		out = append(out, longSign)
	} else {
		out = append(out, syllableToKana(ProtoSyllable{Nucleus: syll.Offglide}, atr, nil)...)
	}

	if syll.Coda != 0 {
		out = append(out, getCodaKana(syll.Coda, nextOnset))
	}

	return out
}

func (w ProtoWord) Kana() string {
	syllableStrings := [][]string{}

	for i, syll := range w.stem.syllables {
		var nextOnset *Consonant = nil
		if i+1 < len(w.stem.syllables) {
			nextOnset = &w.stem.syllables[i+1].Onset
		}
		syllableStrings = append(syllableStrings, syllableToKana(syll, w.stem.atr, nextOnset))
	}

	for _, suffix := range w.suffixes {
		lastSyllable := syllableStrings[len(syllableStrings)-1]
		if suffix.lenite {
			if slices.Contains(codaKanaList, lastSyllable[len(lastSyllable)-1]) {
				if lenitedKana, ok := lenitedKanaMap[lastSyllable[len(lastSyllable)-1]]; ok {
					lastSyllable[len(lastSyllable)-1] = lenitedKana
				}
			} else if len(syllableStrings) > 1 {
				if lenitedKana, ok := lenitedKanaMap[lastSyllable[0]]; ok {
					penult := syllableStrings[len(syllableStrings)-2]
					if penult[len(penult)-1] == codaKana[Stop] {
						penult[len(penult)-1] = ""
					} else {
						lastSyllable[0] = lenitedKana
					}
				}
			}
		}
		if suffix.leadingCoda != 0 {
			lastCoda := lastSyllable[len(lastSyllable)-1]
			if slices.Contains(codaKanaList, lastCoda) {
				lastSyllable[len(lastSyllable)-1] = ""
				syllableStrings = append(syllableStrings, []string{frontKana[codaKanaToOnset[lastCoda]], codaKana[suffix.leadingCoda]})
			} else {
				syllableStrings[len(syllableStrings)-1] = append(lastSyllable, codaKana[suffix.leadingCoda])
			}
		}
		if suffix.syllablesFunc != nil {
			if len(lastSyllable) == 1 {
				syllableStrings[len(syllableStrings)-1] = append(lastSyllable, longSign)
			}
		}
		for i, syll := range suffix.syllables {
			var nextOnset *Consonant = nil
			if i+1 < len(suffix.syllables) {
				nextOnset = &suffix.syllables[i+1].Onset
			}
			syllableStrings = append(syllableStrings, syllableToKana(syll, w.stem.atr, nextOnset))
		}
	}

	out := ""
	for _, syll := range syllableStrings {
		out += strings.Join(syll, "")
	}
	return out
}
