package main

import (
	"fmt"
)

func ladderLength(beginWord, endWord string, wordList []string) int {
	// Convert wordList to a set for O(1) lookups.
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	// If endWord is not in the wordSet, no valid transformation exists.
	if !wordSet[endWord] {
		return 0
	}

	// Define a structure for items in the BFS queue.
	type QueueItem struct {
		word  string
		steps int // number of transformations from beginWord
	}

	// Initialize the queue with the beginWord.
	queue := []QueueItem{{word: beginWord, steps: 1}}

	// Perform BFS.
	for len(queue) > 0 {
		currentItem := queue[0]
		queue = queue[1:]
		currentWord := currentItem.word
		currentSteps := currentItem.steps

		// If we reached endWord, return the number of steps.
		if currentWord == endWord {
			return currentSteps
		}

		// Try changing each letter in the currentWord.
		wordBytes := []rune(currentWord) // using runes in case of Unicode
		for i := 0; i < len(wordBytes); i++ {
			originalChar := wordBytes[i]
			// Replace the letter at position i with every letter 'a' to 'z'
			for c := 'a'; c <= 'z'; c++ {
				if c == originalChar {
					continue
				}
				wordBytes[i] = c
				nextWord := string(wordBytes)
				// If the transformed word is in our set, it's a valid step.
				if wordSet[nextWord] {
					// Append to queue with an incremented step count.
					queue = append(queue, QueueItem{word: nextWord, steps: currentSteps + 1})
					// Remove it from the set so we don't visit it again.
					delete(wordSet, nextWord)
				}
			}
			// Restore the original letter.
			wordBytes[i] = originalChar
		}
	}

	// If the queue is empty and we didn't reach endWord, return 0.
	return 0
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	result := ladderLength(beginWord, endWord, wordList)
	fmt.Println(result)
}
