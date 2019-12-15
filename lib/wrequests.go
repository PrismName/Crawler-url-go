package lib

import (
	"github.com/parnurzeal/gorequest"
	"log"
	"net/http"
)

// WRequest get http response body
func WRequest(targetURL string) (*http.Response, error) {
	request := gorequest.New()
	resp, _, err := request.Get(targetURL).AppendHeader("User-Agent", "__spider_man__").End()
	if err != nil {
		log.Fatal(err)
		return nil, nil
	}

	return resp, nil
}
