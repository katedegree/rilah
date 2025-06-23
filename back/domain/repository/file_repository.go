package repository

import "io"

type FileRepository interface {
	Upload(file io.ReadSeeker, contentType string) (url string, err error)
}
