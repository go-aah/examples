package util

import (
	"path/filepath"

	"aahframe.work"
	"aahframe.work/essentials"
	"aahframe.work/log"
)

func UploadsPath() string {
	return filepath.Join(aah.App().BaseDir(), "uploads")
}

func CreateUploadsDirectory(_ *aah.Event) {
	if err := ess.MkDirAll(UploadsPath(), 0766); err != nil {
		log.Error(err)
	}
}
