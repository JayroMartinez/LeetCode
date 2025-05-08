package main

import (
	"fmt"
)

func convert(s string, numRows int) string {

	strVec := make([]string, numRows)

	xPos := 0
	desc := true

	if numRows > 1 {

		for _, ch := range s {

			strVec[xPos] += string(ch)
			if desc && xPos < numRows-1 {
				xPos++
			} else if desc {
				desc = false
				xPos--
			} else if !desc && xPos > 0 {
				xPos--
			} else {
				xPos++
				desc = true
			}

		}
	} else {
		return s
	}

	aux := ""

	for it := range numRows {
		aux += strVec[it]
	}

	return aux

}

func main() {
	s := "PAYPALISHIRING"
	numRows := 4
	result := convert(s, numRows)
	fmt.Println("Result:", result)
}
