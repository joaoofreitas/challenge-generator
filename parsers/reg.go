package parsers

import "regexp"

var GetContentBetweenSubjects = regexp.MustCompile(`(?m)^-{3,}`)
var GetSubjectNames = regexp.MustCompile(`([\r\n].*?)(:=?\r|\n)(..?(?:------).*)`)

var cleanExercisesNames = regexp.MustCompile("[^a-zA-Z0-9]+")

func GetExercisesNames(rawName string) string {
	return cleanExercisesNames.ReplaceAllString(rawName, "")
}

var GetExercisesRawNames = regexp.MustCompile(`(?m)^\*\*.*\*\*`)
var GetExercisesDescription = regexp.MustCompile(`(?m)^\*\*.*$`)
