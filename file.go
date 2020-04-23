package mfilefile // import "github.com/voytechnology/mfile-file"

import (
	"fmt"
	"io/ioutil"

	"github.com/voytechnology/mfile"
)

func init() {
	if err := mfile.Register("file", handler{}); err != nil {
		panic(fmt.Sprintf("mfile(file): failed registering: %v", err))
	}
}

type handler struct{}

// ReadFile directly from disk
func (handler) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
