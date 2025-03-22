package index

import (
	"fmt"
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

func (i *Index) GetIndex() *Index {
	return i
}

func (i *Index) GetUrls(word string) ([]string, error) {
	const op = "internal/index/GetUrls"

	word = strings.ToLower(word)
	word = strings.TrimSpace(word)
	word = strings.Trim(word, ".,!?-\n\r\t")

	var urls []string

	baseIds, ok := i.Words[word]
	if ok == false {
		return urls, nil
	}

	for _, id := range baseIds {
		for _, doc := range i.Docs {
			if doc.ID == id {
				urls = append(urls, doc.URL)
			}
		}

	}

	if len(baseIds) == 0 {
		return nil, fmt.Errorf("%s:%s", op, "url ids not found")
	}

	return urls, nil
}
