package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getBeefNames() string {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	req, err := http.NewRequest("GET", url, nil)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}
	str := string(body)

	return str
}

func extractBeefNames(w http.ResponseWriter, r *http.Request) {
	beefs := getBeefNames()

	beefs = strings.ReplaceAll(beefs, ".", " ")
	beefs = strings.ReplaceAll(beefs, ",", " ")
	beefs = strings.ReplaceAll(beefs, "\n", " ")

	allBeefs := strings.Split(beefs, " ")

	result := map[string]int{}

	for _, v := range allBeefs {
		tmp := strings.ToLower(v)
		if len(tmp) <= 0 {
			continue
		}

		if _, ok := result[tmp]; ok {
			result[tmp]++
		} else {
			result[tmp] = 1
		}
	}

	fmt.Println(result)
	resp := map[string]interface{}{
		"beef": result,
	}

	json.NewEncoder(w).Encode(resp)

}

func handleRequest() {
	http.HandleFunc("/beef/summary", extractBeefNames)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
