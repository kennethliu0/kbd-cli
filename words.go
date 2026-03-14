package main

import "math/rand"

var commonWords = []string{
	"the", "be", "to", "of", "and", "a", "in", "that", "have", "i",
	"it", "for", "not", "on", "with", "he", "as", "you", "do", "at",
	"this", "but", "his", "by", "from", "they", "we", "say", "her", "she",
	"or", "an", "will", "my", "one", "all", "would", "there", "their", "what",
	"so", "up", "out", "if", "about", "who", "get", "which", "go", "me",
	"when", "make", "can", "like", "time", "no", "just", "him", "know", "take",
	"people", "into", "year", "your", "good", "some", "could", "them", "see", "other",
	"than", "then", "now", "look", "only", "come", "its", "over", "think", "also",
	"back", "after", "use", "two", "how", "our", "work", "first", "well", "way",
	"even", "new", "want", "because", "any", "these", "give", "day", "most", "us",
	"great", "between", "need", "large", "under", "never", "each", "right", "hand", "high",
	"place", "very", "through", "still", "long", "been", "same", "another", "much", "should",
	"last", "life", "state", "since", "world", "house", "keep", "might", "while", "found",
	"own", "part", "old", "home", "small", "end", "put", "help", "here", "show",
	"every", "big", "name", "number", "group", "ask", "such", "turn", "few", "run",
	"move", "must", "tell", "point", "city", "play", "live", "find", "head", "where",
	"those", "may", "down", "side", "change", "line", "set", "act", "try", "around",
	"close", "night", "real", "left", "open", "seem", "next", "walk", "begin", "both",
	"school", "child", "grow", "country", "more", "many", "did", "man", "woman", "said",
	"too", "does", "made", "let", "off", "had", "before", "call", "being", "thing",
}

func generateWords(n int) []string {
	words := make([]string, n)
	for i := range words {
		words[i] = commonWords[rand.Intn(len(commonWords))]
	}
	return words
}
