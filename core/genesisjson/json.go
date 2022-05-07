package genesisjson

import "io"

type Source interface {
	OpenReader(resourceName string) (io.ReadCloser, error)
}

type Target interface {
	OpenWriter(resourceName string) (io.WriteCloser, error)
}
