package main

import (
	"fmt"

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

type namedSuffix struct {
	name   string
	suffix lang.ProtoSuffix
}

func testActualSuffixes() {
	suffixes := []namedSuffix{}
	for name, suffix := range lang.NumberSuffixes {
		suffixes = append(suffixes, namedSuffix{name, suffix})
	}
	for name, suffix := range lang.CaseSuffixes {
		suffixes = append(suffixes, namedSuffix{name, suffix})
	}

	for i := 0; i < 50; i++ {
		pw, err := lang.RandomProtoWord()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		printEvolution(pw)

		for _, ns := range suffixes {
			withSuffix := lang.ApplySuffix(pw, ns.suffix)
			nw := lang.Evolve(withSuffix)
			fmt.Printf("%s + %s -> %s\n", withSuffix.Romanization(), ns.name, nw.Romanization())
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
	testActualSuffixes()
}
