package index

import (
	"scraper-url/internal/crawler"
	"strings"
)

type Index struct {
	Words map[string][]int
	Docs  []crawler.Document
}

func New() *Index {
	return &Index{
		Words: make(map[string][]int),
		Docs:  make([]crawler.Document, 0),
	}
}

func (i *Index) AddWord(word string, docId int) {
	i.Words[word] = append(i.Words[word], docId)
}

func (i *Index) GetUrls(word string) []string {
	word = strings.ToLower(word)
	word = strings.TrimSpace(word)
	word = strings.Trim(word, ".,!?-\n\r\t")

	var urls []string

	baseIds, ok := i.Words[word]
	if ok == false {
		return urls
	}

	for _, id := range baseIds {
		for _, doc := range i.Docs {
			if doc.ID == id {
				urls = append(urls, doc.URL)
			}
		}

	}
	return urls
}
