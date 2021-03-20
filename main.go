package main

import (
	"./fetcher"
	"./parsers"
	"fmt"
)

func main() {
	fmt.Printf("%+v\n", parsers.ParsedData(fetcher.FetchData()))
}
