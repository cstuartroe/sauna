package gloss

import (
	"fmt"
	"strings"

	"github.com/cstuartroe/sauna/dictionary"
	"github.com/cstuartroe/sauna/lang"
)

func getSuffix(name string) (lang.ProtoSuffix, error) {
	for _, suffixSet := range lang.SuffixSets {
		if suffix, ok := suffixSet[name]; ok {
			return suffix, nil
		}
	}
	return lang.ProtoSuffix{}, fmt.Errorf("failed to find suffix %q", name)
}

type GlossedMorpheme struct {
	Form  string
	Gloss string
}

type GlossedWord struct {
	Word   lang.NewWord
	Kana   string
	Pieces []GlossedMorpheme
}

func ParseGloss(gloss string) ([]GlossedWord, error) {
	words := []GlossedWord{}

	wordStrings := strings.Split(gloss, " ")

	for _, wordString := range wordStrings {
		gw := GlossedWord{}

		morphemeStrings := strings.Split(wordString, "-")

		stem, ok := dictionary.Dictionary[morphemeStrings[0]]
		if !ok {
			return nil, fmt.Errorf("failed to find root %q", morphemeStrings[0])
		}
		word := lang.NewProtoWord(stem)
		gw.Pieces = append(gw.Pieces, GlossedMorpheme{
			Form:  lang.Evolve(word.Form()).Romanization(),
			Gloss: morphemeStrings[0],
		})

		for i := 1; i < len(morphemeStrings); i++ {
			suffix, err := getSuffix(morphemeStrings[i])
			if err != nil {
				return nil, err
			}

			word.AddSuffix(suffix)

			suffixForm := suffix.Romanization()
			if morphemeStrings[i] == "SP" {
				suffixForm = "X"
			}
			gw.Pieces = append(gw.Pieces, GlossedMorpheme{
				Form:  "-" + suffixForm,
				Gloss: "-" + morphemeStrings[i],
			})
		}

		gw.Word = lang.Evolve(word.Form())
		gw.Kana = word.Kana()
		words = append(words, gw)
	}

	return words, nil
}

func Kana(words []GlossedWord) string {
	out := ""
	for _, w := range words {
		out += w.Kana
	}
	return out
}

func Text(words []GlossedWord) string {
	out := ""
	for i, w := range words {
		if i > 0 {
			out += " "
		}
		out += w.Word.Romanization()
	}
	return out
}
