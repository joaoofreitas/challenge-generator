package main

import (
	"./fetcher"
	"./parsers"
	"./project"
	"log"
)

func main() {
	log.Printf("%+v\n", parsers.ParsedData(fetcher.FetchData()))
	project.GenerateProject(parsers.ParsedData(fetcher.FetchData()))
}

/*
	Todo:
	Add args for naming the main root folder
	Add tree of that path
	Handle possible errors
	Variable Renaming
	Create a nice ASCII art banner
*/
