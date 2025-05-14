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

func printEvolution(pf lang.ProtoForm) {
	nw := lang.Evolve(pf)

	fmt.Printf("%s -> %s \n", pf.Romanization(), nw.Romanization())
}

func testSuffixes() {
	for i := 0; i < 20; i++ {
		ps, err := lang.RandomProtoStem()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		printEvolution(ps)

		for _, suffix := range lang.TestSuffixes {
			pw := lang.NewProtoWord(ps)
			pw.AddSuffix(suffix)
			printEvolution(pw.Form())
		}

		fmt.Println()
	}
}

func testActualSuffixes() {
	for i := 0; i < 50; i++ {
		ps, err := lang.RandomProtoStem()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		printEvolution(ps)

		for _, suffixSet := range lang.SuffixSets {
			fmt.Println()
			for name, suffix := range suffixSet {
				pw := lang.NewProtoWord(ps)
				pw.AddSuffix(suffix)
				nw := lang.Evolve(pw.Form())
				fmt.Printf("    %s: %s -> %s\n", name, pw.Form().Romanization(), nw.Romanization())
			}
		}

		fmt.Println()
	}
}

func generateParagraph() {
	for i := 0; i < 100; i++ {
		ps, err := lang.RandomProtoStem()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		nw := lang.Evolve(ps)

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
		ps, err := lang.RandomProtoStem()
		if err != nil {
			panic(err)
		}
		printEvolution(ps)
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
