package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/cstuartroe/sauna/gendocs"
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

func cliGloss() {
	words, err := gloss.ParseGloss(os.Args[2])
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println(gloss.Text(words))
}

func generateWords() {
	for i := 0; i < 3; i++ {
		word, err := lang.RandomProtoWord()
		if err != nil {
			panic(err)
		}
		printEvolution(word)
	}
}

var commands = map[string]func(){
	"paragraph":     generateParagraph,
	"test_suffixes": testSuffixes,
	"real_suffixes": testActualSuffixes,
	"gloss":         cliGloss,
	"generate":      generateWords,
	"docs":          gendocs.GenerateDocs,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please pass a command. Available commands: %v\n", strings.Join(slices.Collect(maps.Keys(commands)), ", "))
		return
	}

	commandName := os.Args[1]

	if cmd, ok := commands[commandName]; ok {
		cmd()
	} else {
		fmt.Printf("Unknown command %q\n", commandName)
	}
}
