package app

import (
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/file"
	"cloud.google.com/go/storage"
	"fmt"
	"io/ioutil"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		client, err := storage.NewClient(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to instanciate the storage client: %v", err), http.StatusInternalServerError)
			return
		}

		bucketName, err := file.DefaultBucketName(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get the default bucket name: %v", err), http.StatusInternalServerError)
			return
		}

		file, err := client.Bucket(bucketName).Object("sample.txt").NewReader(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get reader for file: %v", err), http.StatusInternalServerError)
			return
		}

		b, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to read file bytes: %v", err), http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(b); err != nil {
			http.Error(w, fmt.Sprintf("Failed to write response body: %v", err), http.StatusInternalServerError)
		}
	})
}
