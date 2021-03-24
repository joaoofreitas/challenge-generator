package generator

import (
	"bufio"
	"log"
	"os"
)

func GenerateFile(filename string, extension string, content string) {
	file, err := os.Create(filename + extension)
	if err != nil {
		log.Fatalf("Error while trying to create a folder. "+
			"Please try pick other folder name or try to run this program has an administrator. \n"+
			"ERROR: %s", err.Error())
	}

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content + "\n")
	if err != nil {
		log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
	}
	err = writer.Flush()
	if err != nil {
		log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
	}
}
