package main

import (
	"fmt"

	"github.com/cstuartroe/sauna/gloss"
	"github.com/cstuartroe/sauna/lang"
)

func printEvolution(pw lang.ProtoWord) {
	nw := lang.Evolve(pw)

	fmt.Printf("%s -> %s \n", pw.Romanization(), nw.Romanization())
}

func testSuffixes() {
	for i := 0; i < 20; i++ {
		pw, err := lang.RandomProtoWord()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		printEvolution(pw)

		for _, suffix := range lang.TestSuffixes {
			withSuffix := lang.ApplySuffix(pw, suffix)
			printEvolution(withSuffix)
		}

		fmt.Println()
	}
}

func testActualSuffixes() {
	for i := 0; i < 50; i++ {
		pw, err := lang.RandomProtoWord()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		printEvolution(pw)

		for _, suffixSet := range lang.SuffixSets {
			fmt.Println()
			for name, suffix := range suffixSet {
				withSuffix := lang.ApplySuffix(pw, suffix)
				nw := lang.Evolve(withSuffix)
				fmt.Printf("    %s: %s -> %s\n", name, withSuffix.Romanization(), nw.Romanization())
			}
		}

		fmt.Println()
	}
}

func generateParagraph() {
	for i := 0; i < 100; i++ {
		pw, err := lang.RandomProtoWord()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		nw := lang.Evolve(pw)

		fmt.Printf("%s ", nw.Romanization())
	}

	fmt.Println()
}

func main() {
	words, err := gloss.ParseGloss("eat-cnt-gen-par")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println(gloss.Text(words))
}
