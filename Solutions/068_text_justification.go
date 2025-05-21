package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {

	arrRes := []string{}
	//pos := 0
	currLine := []string{}
	currLength := 0

	for i := 0; i < len(words); i++ {
		//fmt.Println("currlength: ", currLength)
		//fmt.Println("word length:", len(words[i]))
		if len(currLine) == 0 {
			// First word in the line: no preceding space.
			currLine = append(currLine, words[i])
			currLength += len(words[i])
		} else if currLength+1+len(words[i]) <= maxWidth {
			// Not the first word: add a space before the word.
			currLine = append(currLine, words[i])
			currLength += 1 + len(words[i])
		} else {
			// flush the current line...
			arrRes = append(arrRes, formatStr(currLine, maxWidth))
			currLine = []string{words[i]}
			currLength = len(words[i])
		}

	}
	//	fmt.Println("pos", pos)
	//	fmt.Println("len res", len(arrRes))
	//arrRes[len(arrRes)-1] = formatLast(arrRes[len(arrRes)-1], maxWidth)
	arrRes = append(arrRes, formatLast(currLine, maxWidth))

	return arrRes

}

func formatStr(currLine []string, maxWidth int) string {
	res := ""
	numWords := len(currLine)
	consumedLength := 0

	//	fmt.Println(currLine)

	for _, elem := range currLine {
		consumedLength += len(elem)
	}

	// totalSpaces is the number of spaces we need to distribute between words
	totalSpaces := maxWidth - consumedLength

	if numWords == 1 {
		// If only one word, left-justify by appending all spaces at the end.
		return currLine[0] + strings.Repeat(" ", totalSpaces)
	}

	gaps := numWords - 1
	base := totalSpaces / gaps
	remainder := totalSpaces % gaps

	// Build the line by appending each word followed by the appropriate spaces.
	for i := 0; i < gaps; i++ {
		res += currLine[i]
		// For the first 'remainder' gaps, add an extra space.
		gapSize := base
		if i < remainder {
			gapSize++
		}
		res += strings.Repeat(" ", gapSize)
	}
	//fmt.Println(res)
	res += currLine[numWords-1] // Append the last word without trailing spaces.
	return res
}

func formatLast(lastLine []string, maxWidth int) string {

	res := ""
	//words := strings.Fields(lastLine)
	//fmt.Println(lastLine)
	//fmt.Println(maxWidth)
	//fmt.Println(len(lastLine))
	//fmt.Println(maxWidth - len(lastLine))

	for j := 0; j < len(lastLine)-1; j++ {
		res += lastLine[j] + " "
	}
	res += lastLine[len(lastLine)-1]
	//fmt.Println(res)
	//fmt.Println(maxWidth - len(res))
	res += strings.Repeat(" ", maxWidth-len(res))

	return res
}

func main() {
	words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth := 20

	result := fullJustify(words, maxWidth)
	for _, line := range result {
		fmt.Printf("\"%s\"\n", line)
	}
}
