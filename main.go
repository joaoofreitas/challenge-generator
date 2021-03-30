package main

import (
	"./fetcher"
	"./parsers"
	"./project"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {

	baseDir := "Project Learning"
	if len(os.Args[1:]) >= 1 {
		if os.Args[1] == "help" || os.Args[1] == "-h" {
			fmt.Printf("HELP: \n\n" +
				"To run this program just simply run it in a terminal. \n" +
				"Eg: ./main \n\n" +
				"If you want to run the program and name a base directory to something of your choice just simply run it " +
				"using a program argument. \n" +
				"Eg: ./main <NameOfBaseDir> \n" +
				"without the \"<, >\" \n\n" +
				"For more information please visit: https://github.com/joaoofreitas and read the README file. \n\n" +
				"This software is licensed by MIT License. \n" +
				"Copyright (c) 2021 João Freitas")
			os.Exit(0)
		} else {
			baseDir = os.Args[1]
		}
	}

	project.GenerateProject(baseDir, parsers.ParsedData(fetcher.FetchData()))

	if runtime.GOOS != "windows" {
		tree, err := exec.Command("tree", baseDir).CombinedOutput()
		if err != nil {
			log.Fatalf("Something happened while trying to print the tree of the generated folder but everything should be working.\n "+
				"If you are having trouble us"+
				"ing this tool please open a issue in:"+
				"http://github.com/joaoofreitas\n "+
				"Error: %s", err.Error())
		}
		fmt.Println(string(tree))
	}
	fmt.Print("\n\n" +
		"________            __                   __\n" +
		"/_  __/ / ___ ____  / /__  __ _____ __ __/ /\n" +
		" / / / _ \\ _ `/ _ \\/  '_/ / // / _ \\ // /_/ \n" +
		"/_/ /_//_\\_,_/_//_/_/\\_\\  \\_, /\\___\\_,_(_)  \n" +
		"                         /___/       " +
		"\n\n")
	fmt.Println("For using free software. \n" +
		"This software is licensed by MIT License. \n" +
		"https://github.com/joaoofreitas\n" +
		"https://joaoofreitas.tech\n" +
		"Copyright (c) 2021 João Freitas")
}
