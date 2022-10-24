package utils

import (
	"fmt"
	"github.com/gocolly/colly"
	"sort"
	"strings"
)

func Crawl(link string, queries []string) []string {
	host := getHost(link)
	type orderedLink struct {
		order int
		link  string
	}
	orderedLinks := make([]orderedLink, 0, 5)
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", string(r.Body), "\nError:", err)
	})
	for index, query := range queries {
		order := index
		c.OnXML(query, func(e *colly.XMLElement) {
			imLink := e.Text
			if strings.HasPrefix(imLink, "/") {
				imLink = host + imLink
			}
			orderedLinks = append(orderedLinks, orderedLink{order, imLink})
		})
	}
	c.Visit(link)
	sort.SliceStable(orderedLinks, func(i, j int) bool {
		return orderedLinks[i].order < orderedLinks[j].order
	})
	imLinks := make([]string, 0, len(orderedLinks))
	for _, ol := range orderedLinks {
		imLinks = append(imLinks, ol.link)
	}
	return imLinks
}
