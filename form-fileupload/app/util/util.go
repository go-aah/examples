package util

import (
	"path/filepath"

	"aahframework.org/aah.v0"
	"aahframework.org/essentials.v0"
	"aahframework.org/log.v0"
)

func UploadsPath() string {
	return filepath.Join(aah.AppBaseDir(), "uploads")
}

func CreateUploadsDirectory(_ *aah.Event) {
	if err := ess.MkDirAll(UploadsPath(), 0766); err != nil {
		log.Error(err)
	}
}
