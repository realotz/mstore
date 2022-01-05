package tests

import (
	"fmt"
	"github.com/eventials/go-tus"
	"os"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) {
	dir,_ :=filepath.Split("/Users/real/tmp/2425a39406d9cba71241e4913680a968.info")
	fmt.Println(dir)
}

func TestBigFileUpload(t *testing.T) {
	f, err := os.Open("cloudwonder-atms-beat-client-1.0.2-mac-amd64-upgrade.dmg")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	// create the tus client.
	client, _ := tus.NewClient("http://127.0.0.1:8000/api/v1/big-upload/37a67144-3f68-4bb6-b555-252b887172f4", nil)

	// create an upload from a file.
	upload, _ := tus.NewUploadFromFile(f)
	// create the uploader.
	uploader, er := client.CreateUpload(upload)
	fmt.Println("---",er)
	//fmt.Println(string(er.(tus.ClientError).Body))
	// start the uploading process.
	err = uploader.Upload()
	//fmt.Println(string(err.(tus.ClientError).Body))
}