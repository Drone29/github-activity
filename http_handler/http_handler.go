package http_handler

import (
	"io"
	"log"
	"net/http"
)

func Get(url string) (body []byte, headers map[string][]string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch data from %s\n%v\n", url, err)
		return
	}
	headers = resp.Header
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body %v\n", err)
		return
	}
	return
}
