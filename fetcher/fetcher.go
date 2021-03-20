package fetcher

import (
	"io/ioutil"
	"net/http"
)

func FetchData() string {
	URL := "https://raw.githubusercontent.com/karan/Projects/master/README.md"
	res, err := http.Get(URL)
	if err != nil {
		print(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		print(err)
	}

	return string(body)
}
