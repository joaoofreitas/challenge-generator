package fetcher

import (
	"io/ioutil"
	"log"
	"net/http"
)

func FetchData() string {
	URL := "https://raw.githubusercontent.com/karan/Projects/master/README.md"
	res, err := http.Get(URL)
	if err != nil {
		log.Fatalf("Some error occurred fetching data from the web. Please check your internet connection. \n"+
			"Error: %s", err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Some error occurred while parsing the body of the request. Please restart the program.\n"+
			"Error: %s", err.Error())
	}

	return string(body)
}
