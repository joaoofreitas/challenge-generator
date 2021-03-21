package generator

import (
	"bufio"
	"log"
	"os"
)

func GenerateFile(filename string, extension string, content string) int {
	// Add validator for file generator...

	file, err := os.Create(filename + extension)
	if err != nil {
		log.Fatal(err)
		return 1
	}

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content + "\n")
	if err != nil {
		log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
		return 1
	}
	writer.Flush()
	return 0
}
