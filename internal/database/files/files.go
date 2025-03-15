package files

import (
	"errors"
	"log"
	"os"
	"path"
)

type FilesDB struct {
	f *os.File
}

func (df FilesDB) Write(data []byte) (n int, err error) {

	f, err := os.Open(df.f.Name())
	if err != nil {
		log.Print("can't open file:%w", err)
		return
	}
	defer f.Close()

	n, err = f.WriteString(string(data))
	if err != nil {
		log.Printf("can't write string: %w", err)
		return
	}
	log.Print("Data has writed in file")
	return
}

func (df FilesDB) Read(data []byte) (n int, err error) {
	b := []byte{}
	_, err = df.f.Read(b)
	if err != nil {
		log.Fatal("can't read file")
	}
	return
}

func New() (*FilesDB, error) {

	goPath := os.Getenv("GOPATH")

	fName := "db.txt"

	p := path.Join(goPath, "searcher-urls/internal/database/files", fName)

	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {

		f, err := os.Create(p)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		fDb := FilesDB{
			f: f,
		}

		return &fDb, nil

	} else {
		file, err := os.Open(p)
		if err != nil {
			return nil, err
		}

		fDb := FilesDB{
			f: file,
		}
		return &fDb, nil
	}

}
