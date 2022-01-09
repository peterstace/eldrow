package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strings"
)

type option struct {
	candidate string
	score     float64
}

func gameLoop(dictionary []string) {
	log.Printf("uniquifying dictionary (before=%d)...", len(dictionary))
	dictionary = uniquifySlice(dictionary)
	log.Printf("done (after=%d)", len(dictionary))

	// Update these manually as you guess.
	log.Printf("applying existing guesses (before=%d)...", len(dictionary))
	for _, g := range []guess{
		{"LANES", "XXXYX"},
	} {
		oldDictionary := dictionary
		dictionary = nil
		for _, word := range oldDictionary {
			if compatible(word, g) {
				dictionary = append(dictionary, word)
			}
		}
	}
	log.Printf("done (after=%d)", len(dictionary))

	var options []option
	for _, candidate := range sample(dictionary, 500) {
		var total, comp int
		for _, simulatedActual := range sample(dictionary, 200) {
			g := guess{meta: calculateMeta(candidate, simulatedActual), word: candidate}
			for _, probe := range sample(dictionary, 200) {
				total++
				if compatible(probe, g) {
					comp++
				}
			}
		}
		score := float64(total-comp) / float64(total)
		options = append(options, option{candidate, score})
	}

	sort.Slice(options, func(i, j int) bool {
		return options[i].score > options[j].score
	})
	for _, opt := range options {
		fmt.Println(opt)
	}
}

func sample(words []string, n int) []string {
	if len(words) <= n {
		return words
	}
	var selected []string
	for i := 0; i < n; i++ {
		for {
			candidate := words[rand.Intn(len(words))]
			if !sliceContains(selected, candidate) {
				selected = append(selected, candidate)
				break
			}
		}
	}
	return selected
}

const (
	green  byte = 'G'
	yellow byte = 'Y'
	grey   byte = 'X'
)

type guess struct {
	word string
	meta string
}

func calculateMeta(guess, actual string) string {
	var meta strings.Builder
	for i := range guess {
		if guess[i] == actual[i] {
			meta.WriteByte(green)
			continue
		}

		var found bool
		for j := range actual {
			if guess[i] == actual[j] {
				found = true
				break
			}
		}
		if found {
			meta.WriteByte(yellow)
		} else {
			meta.WriteByte(grey)
		}
	}
	return meta.String()
}

func compatible(candidate string, guess guess) bool {
	for i := range guess.word {
		switch guess.meta[i] {
		case green:
			if candidate[i] != guess.word[i] {
				return false
			}
		case yellow:
			if candidate[i] == guess.word[i] {
				return false
			}

			contains := false
			for j := range candidate {
				if candidate[j] == guess.word[i] {
					contains = true
					break
				}
			}
			if !contains {
				return false
			}
		case grey:
			for j := range candidate {
				if candidate[j] == guess.word[i] {
					return false
				}
			}
		default:
			panic(guess.meta[i])
		}
	}
	return true
}
