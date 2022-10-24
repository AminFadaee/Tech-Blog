package techBlogs

import (
	"github.com/k3a/html2text"
	"strings"
)

type Blog struct {
	Name             string
	Url              string
	imageQueries     []string
	summaryProcessor func(string) string
}

var Blogs = []Blog{
	{
		Name: "facebook",
		Url:  "https://engineering.fb.com/feed/",
		imageQueries: []string{
			"//figure[@id=\"post-feat-image-container\"]/img/@src",
			"//img[contains(@class, \"wp-image\")]/@src",
		},
		summaryProcessor: func(summary string) string {
			if summary != "" {
				summary = html2text.HTML2Text(strings.Split(summary, "\n")[0])
			}
			return summary
		},
	},
	{
		Name:         "linkedin",
		Url:          "https://engineering.linkedin.com/blog.rss.html",
		imageQueries: []string{"//li[@class=\"resource-image\"]/img/@src"},
	},
	{
		Name: "github",
		Url:  "https://github.blog/category/engineering/feed/",
		imageQueries: []string{
			"//img[contains(@class, \"wp-image\")]/@src",
			"//div[contains(@class, \"position-relative\") and contains(@class, \"container-xl\") and contains(@class, \"mx-auto\") and contains(@class, \"p-responsive\")]/img/@src",
		},
	},
}
