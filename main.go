package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type Exercise struct {
	name        string
	description string
}

type Subject struct {
	name      string
	exercises []Exercise
}

func parseMD(text string) []string { // Can be optimized in terms of not even store the first introductory part of the README
	var sections []string
	var tmp string

	re := regexp.MustCompile(`(?m)^-{3,}`)
	textToArray := strings.Split(strings.Replace(text, "\r\n", "\n", -1), "\n")
	textToArray = append(textToArray, "\n----")

	for _, line := range textToArray {
		if !re.MatchString(line) {
			tmp += line + "\n"
		} else {
			sections = append(sections, tmp)
			tmp = ""
		}
	}
	return sections
}

func main() {
	var listOfSubjects []Subject

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
	rawMDList := parseMD(rawMD)

	subjectRegex := regexp.MustCompile(`([\r\n].*?)(:=?\r|\n)(..?(?:------).*)`)
	nameParsingRegex := regexp.MustCompile("[^a-zA-Z0-9]+") // Used for removing all non alphabetic characters from the names of the exercises

	for i, subjectMatches := range subjectRegex.FindAllString(rawMD, -1) {
		var listOfExercises []Exercise
		var nameMatches, descriptionMatches []string

		nameRegex := regexp.MustCompile(`(?m)^\*\*.*\*\*`)
		descriptionRegex := regexp.MustCompile(`(?m)^\*\*.*$`)

		nameMatches = nameRegex.FindAllString(rawMDList[i+1], -1)
		descriptionMatches = descriptionRegex.FindAllString(rawMDList[i+1], -1)

		if len(nameMatches) != len(descriptionMatches) {
			panic("Parsing Error...")
		}

		for j := 0; j < len(nameMatches); j++ {
			var exercise Exercise

			processedString := nameParsingRegex.ReplaceAllString(nameMatches[j], "") // Used for removing all non alphabetic characters from the names of the exercises

			exercise.name = processedString
			exercise.description = descriptionMatches[j]
			listOfExercises = append(listOfExercises, exercise)
		}
		subject := Subject{name: subjectMatches, exercises: listOfExercises}
		listOfSubjects = append(listOfSubjects, subject)
	}
	fmt.Printf("%+v\n", listOfSubjects)
}
