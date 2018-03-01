package core

import (
	"net/http"
	"io/ioutil"
)

func HttpGet(url string) (content string, error error) {
	response, e := http.Get(url)
	if e != nil {
		return "", e
	}
	defer response.Body.Close()
	body, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return "", e
	}
	return string(body), nil
}
