package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

func MakeSimpleRequest(url string) []byte {
	request, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
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
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.134 Safari/537.36")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body
}

func MakeRequestAndCreateDocument(url string) *goquery.Document {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.134 Safari/537.36")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic("Error")
	}
	return doc
}