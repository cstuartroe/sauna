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
	Pieces []GlossedMorpheme
}

func ParseGloss(gloss string) ([]GlossedWord, error) {
	words := []GlossedWord{}

	dummyWord, err := lang.NewProtoWord(
		false,
		[]lang.ProtoSyllable{
			{
				Nucleus: lang.Central,
			},
		},
	)
	if err != nil {
		return words, err
	}

	wordStrings := strings.Split(gloss, " ")

	for _, wordString := range wordStrings {
		gw := GlossedWord{}

		morphemeStrings := strings.Split(wordString, "-")

		word, ok := dictionary.Dictionary[morphemeStrings[0]]
		if !ok {
			return nil, fmt.Errorf("failed to find root %q", morphemeStrings[0])
		}
		gw.Pieces = append(gw.Pieces, GlossedMorpheme{
			Form:  lang.Evolve(word).Romanization(),
			Gloss: morphemeStrings[0],
		})

		for i := 1; i < len(morphemeStrings); i++ {
			suffix, err := getSuffix(morphemeStrings[i])
			if err != nil {
				return nil, err
			}

			word = lang.ApplySuffix(word, suffix)

			suffixWord := lang.Evolve(lang.ApplySuffix(dummyWord, suffix))
			suffixForm := suffixWord.Romanization()[1:]
			suffixForm = strings.ReplaceAll(suffixForm, "a", "A")
			suffixForm = strings.ReplaceAll(suffixForm, "e", "I")
			suffixForm = strings.ReplaceAll(suffixForm, "o", "U")
			suffixForm = strings.ReplaceAll(suffixForm, "i", "I")
			suffixForm = strings.ReplaceAll(suffixForm, "u", "U")
			if suffixForm == "" {
				suffixForm = "∅"
			}
			if morphemeStrings[i] == "SP" {
				suffixForm = "X"
			}
			if suffix.Lenite() {
				suffixForm = "'" + suffixForm
			}
			gw.Pieces = append(gw.Pieces, GlossedMorpheme{
				Form:  "-" + suffixForm,
				Gloss: "-" + morphemeStrings[i],
			})
		}

		gw.Word = lang.Evolve(word)
		words = append(words, gw)
	}

	return words, nil
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
