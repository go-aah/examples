// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// aahframework.org/examples source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import (
	"fmt"
	"path/filepath"
	"time"

	"aahframework.org/aah.v0"
	"aahframework.org/essentials.v0"
	"aahframework.org/examples/form-fileupload/app/models"
	"aahframework.org/examples/form-fileupload/app/util"
)

// AppController struct application controller
type AppController struct {
	*aah.Context
}

// Index method is application home page.
func (a *AppController) Index() {
	data := aah.Data{
		"Greet": models.Greet{
			Message: "Welcome to aah framework - File Upload Example",
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

	dstFile := filepath.Join(util.UploadsPath(), filename)

	// Using aah convenient methoed to save uploaded file
	size, err := a.Req.SaveFile("fileUpload", dstFile)
	if err != nil {
		a.Reply().HTML(aah.Data{
			"IsError": true,
		})
		return
	}

	a.Reply().HTML(aah.Data{
		"SavedFileName": filepath.Base(dstFile),
		"FileSize":      size,
	})
}
