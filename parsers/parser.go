package parsers

import (
	"../structs"
	"strings"
)

func parseMD(text string) []string { // Can be optimized in terms of not even store the first introductory part of the README
	var sections []string
	var tmp string

	textToArray := strings.Split(strings.Replace(text, "\r\n", "\n", -1), "\n")
	textToArray = append(textToArray, "\n----")

	for _, line := range textToArray {
		if !GetContentBetweenSubjects.MatchString(line) {
			tmp += line + "\n"
		} else {
			sections = append(sections, tmp)
			tmp = ""
		}
	}
	return sections
}

func ParsedData(md string) []structs.Subject {
	var listOfSubjects []structs.Subject

	parsedMD := parseMD(md)

	for i, subjectMatches := range GetSubjectNames.FindAllString(md, -1) {
		var listOfExercises []structs.Exercise
		var nameMatches, descriptionMatches []string

		nameMatches = GetExercisesRawNames.FindAllString(parsedMD[i+1], -1)
		descriptionMatches = GetExercisesDescription.FindAllString(parsedMD[i+1], -1)

		if len(nameMatches) != len(descriptionMatches) {
			panic("Parsing Error...")
		}

		for j := 0; j < len(nameMatches); j++ {
			var exercise structs.Exercise

			exercise.Name = GetExercisesName(nameMatches[j])
			exercise.Description = descriptionMatches[j]

			listOfExercises = append(listOfExercises, exercise)
		}
		subject := structs.Subject{Name: subjectMatches, Exercises: listOfExercises}
		listOfSubjects = append(listOfSubjects, subject)
	}
	return listOfSubjects
}
