package parsers

import "regexp"

var GetContentBetweenSubjects = regexp.MustCompile(`(?m)^-{3,}`)
var cleanExercisesNames = regexp.MustCompile("[^a-zA-Z0-9]+")

var GetSubjectNames = regexp.MustCompile(`([\r\n].*?)(:=?\r|\n)(..?(?:------).*)`)

var GetExercisesRawNames = regexp.MustCompile(`(?m)^\*\*.*\*\*`)
var GetExercisesDescription = regexp.MustCompile(`(?m)^\*\*.*$`)
