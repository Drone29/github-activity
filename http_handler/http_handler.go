package http_handler

import (
	"errors"
	"io"
	"net/http"
)

func Get(url string) (body []byte, headers map[string][]string, err error) {
	resp, err := http.Get(url)
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	}
	if err != nil {
		return
	}
	headers = resp.Header
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
