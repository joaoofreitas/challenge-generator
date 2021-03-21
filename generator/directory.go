package generator

import (
	"log"
	"os"
)

func GenerateDir(name string) {
	err := os.Mkdir(name, 0777)
	if err != nil {
		log.Fatal(err)
	}
}
