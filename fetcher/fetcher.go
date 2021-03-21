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
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
