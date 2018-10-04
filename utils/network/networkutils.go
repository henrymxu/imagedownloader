package network

import (
	"io/ioutil"
	"net/http"
)

func MakeSimpleRequest(url string) []byte {
	request, err := http.Get(url)
	check(err)
	defer request.Body.Close()
	body, _ := ioutil.ReadAll(request.Body)
	return body
}

func MakeRequestWithQuery(url string, queries map[string]string) []byte {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	query := request.URL.Query()
	for key, value := range queries {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()
	response, err := client.Do(request)
	check(err)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}