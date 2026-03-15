package main

import "math/rand"

type wordList struct {
	name  string
	words []string
}

var wordLists = []wordList{
	{name: "english", words: english200},
	{name: "english 1k", words: english1k},
	{name: "english 5k", words: english5k},
}

func generateWords(n int, listIdx int) []string {
	src := wordLists[listIdx].words
	words := make([]string, n)
	for i := range words {
		words[i] = src[rand.Intn(len(src))]
	}
	return words
}
