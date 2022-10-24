package main

import (
	"encoding/json"
	"fmt"
	"spymaster.com/techBlogs"
	"strings"
)

func main() {
	for _, b := range techBlogs.Blogs {
		tbs, err := b.GetEntries()
		if err == nil {
			for _, tb := range tbs {
				f, _ := json.Marshal(tb)
				fmt.Println(string(f))
				fmt.Println(strings.Repeat("_", 100))
			}
		} else {
			fmt.Println(err)
		}
	}
}
