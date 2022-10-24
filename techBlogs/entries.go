package techBlogs

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"spymaster.com/utils"
	"time"
)

type BlogEntry struct {
	Blog    string
	Title   string
	Link    string
	Time    *time.Time
	Tags    []string
	Summary string
	Images  []utils.File
}

func (b *Blog) initialize() {
	if b.summaryProcessor == nil {
		b.summaryProcessor = func(summary string) string {
			return summary
		}
	}
}

func (b *Blog) GetEntries() (entries []BlogEntry, err error) {
	start := time.Now()
	defer func() { end := time.Now(); fmt.Println(end.Sub(start)) }()
	b.initialize()
	entries = make([]BlogEntry, 0, 10)
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(b.Url)
	if err != nil {
		return
	}
	tc := make(chan *BlogEntry)
	for _, item := range feed.Items {
		go createEntry(b, item, tc)
	}
	for range feed.Items {
		tbp := <-tc
		if tbp != nil {
			entries = append(entries, *tbp)
		}
	}
	return
}

func createEntry(b *Blog, item *gofeed.Item, ch chan<- *BlogEntry) {
	ch <- &BlogEntry{
		Blog:    b.Name,
		Title:   item.Title,
		Link:    item.Link,
		Time:    item.PublishedParsed,
		Tags:    item.Categories,
		Summary: b.summaryProcessor(item.Description),
		Images: utils.DownloadFiles(
			utils.Crawl(item.Link, b.imageQueries),
		),
	}
}
