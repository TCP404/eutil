package eutil

import (
	"io/ioutil"

	"github.com/parnurzeal/gorequest"
)

func DownloadBinary(url, path string) []error {
	_, body, errs := gorequest.New().Get(url).EndBytes()
	if len(errs) != 0 {
		return errs
	}
	err := ioutil.WriteFile(path, body, 0755)
	if err != nil {
		return []error{err}
	}
	return nil
}
