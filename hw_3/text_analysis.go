package main

import (
	"strconv"
	"strings"
)

const topСount = 10

type wordSlice struct {
	Value string
	Count int
}

func TextAnalysis(text string) string {
	wordCountMap := prepareMap(text)

	wordSlices := prepareSlices(wordCountMap)
	wordSlices = descendingSort(wordSlices)
	wordSlices = lexicographicSort(wordSlices)

	if len(wordSlices) > topСount {
		return outputResult(wordSlices[:topСount])
	}

	return outputResult(wordSlices)
}

func prepareMap(text string) map[string]int {
	words := strings.Fields(text)
	wordCountMap := make(map[string]int)
	for _, word := range words {
		normalized := strings.Trim(word, ",.!?:;\"'")
		wordCountMap[normalized]++
	}

	return wordCountMap
}

func prepareSlices(wordCountMap map[string]int) []wordSlice {
	var wordSlices []wordSlice
	for word, count := range wordCountMap {
		wordSlices = append(wordSlices, wordSlice{Value: word, Count: count})
	}

	return wordSlices
}

func descendingSort(wordSlices []wordSlice) []wordSlice {
	for i := 0; i < len(wordSlices)-1; i++ {
		for j := 0; j < len(wordSlices)-i-1; j++ {
			if wordSlices[i].Count > wordSlices[j].Count {
				temporalValue := wordSlices[i]
				wordSlices[i] = wordSlices[j]
				wordSlices[j] = temporalValue
			}
		}
	}

	return wordSlices
}

func lexicographicSort(wordSlices []wordSlice) []wordSlice {
	for i := 0; i < len(wordSlices)-1; i++ {
		for j := 0; j < len(wordSlices)-i-1; j++ {
			if wordSlices[i].Count == wordSlices[j].Count &&
				strings.ToLower(wordSlices[i].Value) < strings.ToLower(wordSlices[j].Value) {
				temporalValue := wordSlices[i]
				wordSlices[i] = wordSlices[j]
				wordSlices[j] = temporalValue
			}
		}
	}

	return wordSlices
}

func outputResult(wordSlices []wordSlice) string {
	var result string
	for i := 0; i < len(wordSlices); i++ {
		result += wordSlices[i].Value + " (" + strconv.Itoa(wordSlices[i].Count) + ")\n"
	}

	return result
}
