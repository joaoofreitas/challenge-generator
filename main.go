package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	res, err := http.Get("https://raw.githubusercontent.com/karan/Projects/master/README.md")
	if err != nil {
		print(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		print(err)
	}
	rawMD := string(body)

	subjectRegex := regexp.MustCompile(`([\r\n].*?)(:=?\r|\n)(..?(?:------).*)`)
	for i, match := range subjectRegex.FindAllString(rawMD, -1) {
		fmt.Println(match, "found at index", i)
	}

	var nameRegex = regexp.MustCompile(`(?m)^\*\*.*\*\*`)
	for i, match := range nameRegex.FindAllString(rawMD, -1) {
		fmt.Println(match, "found at index", i)
	}

	var descriptionRegex = regexp.MustCompile(`(?m)^\*\*.*\*`)
	for i, match := range descriptionRegex.FindAllString(rawMD, -1) {
		fmt.Println(match, " match found at index", i)
	}
}
