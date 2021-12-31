package tests

import (
	"os"
	"testing"
	"github.com/eventials/go-tus"
)

func TestBigFileUpload(t *testing.T) {
	f, err := os.Open("my-file.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	// create the tus client.
	client, _ := tus.NewClient("http://127.0.0.1:8000/files", nil)

	// create an upload from a file.
	upload, _ := tus.NewUploadFromFile(f)

	// create the uploader.
	uploader, _ := client.CreateUpload(upload)

	// start the uploading process.
	uploader.Upload()
}