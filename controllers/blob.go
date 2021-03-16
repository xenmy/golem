package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/go-chi/chi/v5"

	"github.com/xenmy/golem/config"
)

var blobHandler *BlobHandler

type BlobHandler struct {
	Account      string
	Accountkey   string
	Container    string
	ContainerURL azblob.ContainerURL
}

func BlobInit() {
	config := config.GetConfig()
	blobHandler = &BlobHandler {
		Account : config.GetString("azure.storage.account"),
		Accountkey : config.GetString("azure.storage.accountkey"),
		Container : config.GetString("azure.storage.container"),
	}

	containerUrlTempl := "https://%s.blob.core.windows.net/%s"
	cred, err := azblob.NewSharedKeyCredential(blobHandler.Account, blobHandler.Accountkey)
	if err != nil {
		log.Fatal(err)
	}
	containerUrlText := fmt.Sprintf(containerUrlTempl, blobHandler.Account, blobHandler.Container)
	pl := azblob.NewPipeline(cred, azblob.PipelineOptions{})
	cURL, _ := url.Parse(containerUrlText)
	blobHandler.ContainerURL = azblob.NewContainerURL(*cURL, pl)
}

func GetBlobHandler() *BlobHandler {
	return blobHandler
}

func (b BlobHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	containerPropRes, err := b.ContainerURL.GetProperties(ctx, azblob.LeaseAccessConditions{})
	if err != nil {
		log.Println(err)
		http.Error(w, "Could not access container.", 500)
	} else {
		res := containerPropRes.Response()
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(res.Status))
	}
}

func (b BlobHandler) DonwloadBlobData(w http.ResponseWriter, r *http.Request) {
	path := chi.URLParam(r, "*")
	blobURL := b.ContainerURL.NewBlobURL(path)
	ctx := context.Background()

	var blobSize int64 = 1024
	var blobType string = "text/plain"
	blobPropRes, err := blobURL.GetProperties(ctx, azblob.BlobAccessConditions{}, azblob.ClientProvidedKeyOptions{})
	if err != nil {
		log.Printf("%#v\n", err)
		http.Error(w, http.StatusText(404), 404)
	} else {
		blobSize = blobPropRes.ContentLength()
		blobType = blobPropRes.ContentType()
	}

	data := make([]byte, blobSize)
	err = azblob.DownloadBlobToBuffer(ctx, blobURL, 0, azblob.CountToEnd, data, azblob.DownloadFromBlobOptions{})
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	} else {
		w.Header().Set("Content-Type", blobType)
		w.Write([]byte(data))
	}
}
