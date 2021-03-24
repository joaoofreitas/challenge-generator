package generator

import (
	"log"
	"os"
)

func GenerateDir(name string) {
	err := os.Mkdir(name, 0777)
	if err != nil {
		log.Fatalf("Error while trying to create a folder. \n"+
			"Please try pick other folder name or try to run this program has an administrator. \n"+
			"Error: %s", err.Error())
	}
}
