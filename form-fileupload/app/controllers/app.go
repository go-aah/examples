// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// go-aah/tutorials source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import (
	"fmt"
	"path/filepath"
	"time"

	"aahframework.org/aah.v0"
	"aahframework.org/essentials.v0"
	"aahframework.org/log.v0"
	"github.com/go-aah/tutorials/form-fileupload/app/models"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Index method is application home page.
func (a *AppController) Index() {
	data := aah.Data{
		"Greet": models.Greet{
			Message: "Welcome to aah framework - File Upload Tutorial",
		},
	}

	a.Reply().Ok().HTML(data)
}

// FileUpload method is to save uploaded multipart into destination path.
func (a *AppController) FileUpload(fileName string) {
	filename := fmt.Sprintf("%s-%s%s",
		ess.StripExt(fileName),
		time.Now().Format("2006-01-02-15-04-05"),
		filepath.Ext(fileName))

	dstFile := filepath.Join(dstPath(), filename)

	// Using aah convenient methoed to save uploaded file
	size, err := a.Req.SaveFile("fileUpload", dstFile)
	if err != nil {
		a.Reply().HTML(aah.Data{
			"IsError": true,
		})
		return
	}

	// OR
	// You can use traditional way using
	// a.Req.FormFile("fileUpload")

	a.Reply().HTML(aah.Data{
		"SavedFileName": filepath.Base(dstFile),
		"FileSize":      size,
	})
}

// Helper methods to create upload directory
// and register aah server on start event

func dstPath() string {
	return filepath.Join(aah.AppBaseDir(), "uploads")
}

func createDirectory(_ *aah.Event) {
	if err := ess.MkDirAll(dstPath(), 0766); err != nil {
		log.Error(err)
	}
}

func init() {
	aah.OnStart(createDirectory)
}
