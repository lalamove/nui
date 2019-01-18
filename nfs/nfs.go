package nfs

import (
	"io"
	"os"
)

// DefaultFileSystem is a file system wrapping os methods
var DefaultFileSystem = OSFileSystem{}

// FileSystem is an interface to wrap a file system
type FileSystem interface {
	// Open opens a file
	Open(string) (io.ReadCloser, error)
}

// OSFileSystem is the file system implementation wrapping the `os` package
type OSFileSystem struct{}

// Open opens a file, it wraps os.Open
func (osfs OSFileSystem) Open(p string) (io.ReadCloser, error) {
	return os.Open(p)
}
