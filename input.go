package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InputGuess() (guess, error) {
	word, err := Input("guessed word")
	if err != nil {
		return guess{}, err
	}
	word = strings.TrimSpace(word)

	// TODO: word validation
	fmt.Println("WORD", []byte(word))

	meta, err := Input("solution hint")
	if err != nil {
		return guess{}, err
	}
	meta = strings.TrimSpace(meta)

	// TODO: meta validation
	fmt.Println("META", []byte(meta))

	return guess{word: word, meta: meta}, nil
}

func Input(prompt string) (string, error) {
	fmt.Printf(prompt + " > ")
	r := bufio.NewReader(os.Stdin)
	return r.ReadString('\n')
}
