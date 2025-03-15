package index

import "scraper-url/internal/crawler"

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

func (i *Index) DocsID(word string) []int {
	return i.Words[word]
}
