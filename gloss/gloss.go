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

func ParseGloss(gloss string) ([]lang.NewWord, error) {
	words := []lang.NewWord{}

	wordStrings := strings.Split(gloss, " ")

	for _, wordString := range wordStrings {
		morphemeStrings := strings.Split(wordString, "-")

		word, ok := dictionary.Dictionary[morphemeStrings[0]]
		if !ok {
			return nil, fmt.Errorf("failed to find root %q", morphemeStrings[0])
		}

		for i := 1; i < len(morphemeStrings); i++ {
			suffix, err := getSuffix(morphemeStrings[i])
			if err != nil {
				return nil, err
			}

			word = lang.ApplySuffix(word, suffix)
		}

		words = append(words, lang.Evolve(word))
	}

	return words, nil
}

func Text(words []lang.NewWord) string {
	out := ""
	for i, w := range words {
		if i > 0 {
			out += " "
		}
		out += w.Romanization()
	}
	return out
}
