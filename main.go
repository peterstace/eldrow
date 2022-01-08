package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	wordFile := flag.String(
		"word-file",
		"/usr/share/dict/words",
		"location of a newline delimited list of dictionary words",
	)
	flag.Parse()

	words, err := loadWords(*wordFile)
	if err != nil {
		log.Fatal(err)
	}
	words = selectWords(words)

	for _, w := range words {
		fmt.Println(w)
	}
}

func loadWords(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scn := bufio.NewScanner(f)
	var lines []string
	for scn.Scan() {
		lines = append(lines, scn.Text())
	}
	return lines, nil
}

func selectWords(rawWords []string) []string {
	var normalised []string
	for _, w := range rawWords {
		if w, ok := normaliseWord(w); ok {
			normalised = append(normalised, w)
		}
	}
	sort.Strings(normalised)
	return normalised
}

func normaliseWord(candidate string) (string, bool) {
	if len(candidate) != 5 {
		return "", false
	}
	for _, c := range candidate {
		if (c < 'A' || c > 'Z') && (c < 'a' || c > 'z') {
			return "", false
		}
	}
	return strings.ToUpper(candidate), true
}
