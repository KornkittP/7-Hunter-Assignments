package main

import (
	"fmt"
)

func decode(code string) {
	codeLength := len(code) + 1
	temp := make([]int, 0)
	for i := 0; i < codeLength; i++ {
		temp = append(temp, codeLength)
	}

	min := codeLength

	for i, v := range code {
		if string(v) == "=" {
			temp[i+1] = temp[i]
		}

		if string(v) == "R" {
			temp[i+1] = temp[i] + 1
		}

		if string(v) == "L" {
			temp[i+1] = temp[i] - 1
			if temp[i+1] < min {
				min = temp[i+1]
			}
		}
	}

	for i, _ := range temp {
		temp[i] = temp[i] - min
	}

	fmt.Println(code, temp)
}

func main() {
	decode("LLRR=")
	decode("RLRL=")
	decode("RR=LR")
	decode("RR=LL")
}
