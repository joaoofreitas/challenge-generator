package project

import (
	"../generator"
	"../types"
	"path/filepath"
)

func GenerateProject(baseDir string, model []types.Subject) {
	generator.GenerateDir(baseDir)

	for _, subject := range model {
		subjectPath := filepath.Join(baseDir, subject.Name)
		generator.GenerateDir(subjectPath)
		for _, exercise := range subject.Exercises {
			exercisePath := filepath.Join(subjectPath, exercise.Name)
			generator.GenerateDir(exercisePath)
			generator.GenerateFile(filepath.Join(exercisePath, "README"), ".md", exercise.Description)
		}
	}
}
