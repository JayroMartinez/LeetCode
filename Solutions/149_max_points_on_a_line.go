package main

import (
	"fmt"
)

func maxPoints(points [][]int) int {
	length := len(points)
	var maxPo int
	//var vert bool
	//var hor bool

	if length == 0 {
		maxPo = 0

	} else if length == 1 {
		maxPo = 1
	} else {
		maxPo = 2
		for i := 0; i < length-2; i++ {
			x1 := points[i][0]
			y1 := points[i][1]

			for j := i + 1; j < length-1; j++ {
				x2 := points[j][0]
				y2 := points[j][1]
				auxPo := 2
				/*if x1 == x2 {
					vert = true
					hor = false
				} else if y1 == y2 {
					hor = true
					vert = false
				}*/
				for k := j + 1; k < length; k++ {
					/*if vert && x1 == points[k][0] {
						auxPo++
					} else if hor && y1 == points[k][1] {
						auxPo++
					} else*/if (y2-y1)*(points[k][0]-x1) == (points[k][1]-y1)*(x2-x1) {
						auxPo++
					}
				}

				if auxPo > maxPo {
					maxPo = auxPo
				}
			}
		}
	}
	return maxPo
}

func main() {
	points := [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}
	result := maxPoints(points)
	fmt.Println(result)
}
