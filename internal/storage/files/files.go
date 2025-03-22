package files

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"scraper-url/internal/crawler"
)

type Storage struct {
	file *os.File
	path string
}

func New() (*Storage, error) {
	const op = "storage.file.New"

	goPath := os.Getenv("GOPATH")

	log.Printf("gopath:%s", goPath)

	fName := "db.txt"

	p := path.Join(goPath, "scraper-url/internal/storage/files", fName)

	log.Printf("path:%s", p)

	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {

		f, err := os.Create(p)
		if err != nil {
			return nil, fmt.Errorf("%s:%w", op, err)
		}
		defer f.Close()

		fDb := Storage{
			file: f,
			path: p,
		}

		return &fDb, nil

	} else {
		file, err := os.OpenFile(p, os.O_RDONLY, os.ModeDir)
		if err != nil {
			return nil, fmt.Errorf("%s:%w", op, err)
		}

		defer file.Close()

		fDb := Storage{
			file: file,
			path: p,
		}
		return &fDb, nil
	}

}

func (s *Storage) Save(docs []crawler.Document) error {
	const op = "storage.files.Save"

	b, err := json.Marshal(docs)
	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}

	_, err = s.Write([]byte(b))
	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}

	return nil
}

func (s *Storage) Write(data []byte) (n int, err error) {

	const op = "storage.files.Write"

	f, err := os.OpenFile(s.path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return 0, fmt.Errorf("%s:%w", op, err)
	}
	defer f.Close()

	n, err = f.Write(data)
	if err != nil {
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	return
}

func (s *Storage) Read(data []byte) (n int, err error) {
	const op = "storage.files.Read"

	b := []byte{}
	_, err = s.file.Read(b)
	if err != nil {
		return 0, fmt.Errorf("%s:%w", op, err)
	}
	return
}
