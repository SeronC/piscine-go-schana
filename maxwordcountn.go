package piscine

import (
	"sort"
	"strings"
)

func MaxWordCountN(text string, n int) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Split(text, " ")
	for _, word := range words {
		wordCount[word]++
	}

	if len(wordCount) == 0 {
		return map[string]int{"": len(words)}
	}

	type WordCount struct {
		Word  string
		Count int
	}

	wordCounts := make([]WordCount, 0, len(wordCount))
	for word, count := range wordCount {
		wordCounts = append(wordCounts, WordCount{word, count})
	}

	// Sort wordCounts by count (descending) and then by word (ascending)
	sort.SliceStable(wordCounts, func(i, j int) bool {
		if wordCounts[i].Count != wordCounts[j].Count {
			return wordCounts[i].Count > wordCounts[j].Count
		}
		return wordCounts[i].Word < wordCounts[j].Word
	})

	result := make(map[string]int)
	for i := 0; i < n && i < len(wordCounts); i++ {
		result[wordCounts[i].Word] = wordCounts[i].Count
	}

	return result
}
