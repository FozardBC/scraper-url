package database

import "io"

type Database interface {
	io.Writer
	io.Reader
}
