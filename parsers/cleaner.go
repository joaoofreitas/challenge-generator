package parsers

func CleanName(rawName string) string {
	return cleanExercisesNames.ReplaceAllString(rawName, "")
}
