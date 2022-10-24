package utils

import (
	"log"
	"net/url"
)

func getHost(input string) (host string) {
	url, err := url.Parse(input)
	if err != nil {
		log.Fatal(err)
	}
	host = url.Scheme + "://" + url.Hostname()
	return
}
