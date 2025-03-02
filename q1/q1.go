package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//-- from file
	jsonFile, err := os.Open("./files/hard.json")
	if err != nil {
		fmt.Println("err:", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result [][]int
	json.Unmarshal([]byte(byteValue), &result)

	var final [][]int
	nextrow := result[len(result)-1]

	for i := len(result) - 2; i >= 0; i-- {
		var temp []int

		for j := 0; j < i+1; j++ {
			left := result[i][j] + nextrow[j]
			right := result[i][j] + nextrow[j+1]

			if left > right {
				temp = append(temp, left)
			} else {
				temp = append(temp, right)
			}
		}

		final = append(final, temp)
		nextrow = temp
	}

	fmt.Println(final[len(result)-2][0])
}
