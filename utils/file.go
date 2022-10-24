package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"spymaster.com/settings"
	"strings"
)

type File struct {
	Order int
	Link  string
	Path  string
}

func DownloadFiles(links []string) (files []File) {
	fc := make(chan *File)
	for i, link := range links {
		go downloadFile(i+1, link, fc)
	}
	for range links {
		if f := <-fc; f != nil {
			files = append(files, *f)
		}
	}
	return
}

func downloadFile(n int, link string, fc chan *File) {
	path := generateFilePathFromLink(link)
	err := downloadFileToPath(path, link)
	if err == nil {
		fc <- &File{n, link, path}
	} else {
		fmt.Println("Error while downloading the file:", err)
		fc <- nil
	}
}

func getPureLink(link string) string {
	index := strings.Index(link, "?")
	if index == -1 {
		return link
	}
	return link[:index]
}

func generateFilePathFromLink(link string) (path string) {
	ext := filepath.Ext(getPureLink(link))
	name := fmt.Sprintf("%x", md5.Sum([]byte(link))) + ext
	return filepath.Join(settings.FilesDir, name)
}

func downloadFileToPath(path string, url string) error {
	if _, err := os.Stat(path); err == nil {
		return nil // No need to download
	} else if os.IsNotExist(err) {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		out, err := os.Create(path)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		return err
	} else {
		fmt.Println("Downloading failed:", err)
		return err
	}
}
