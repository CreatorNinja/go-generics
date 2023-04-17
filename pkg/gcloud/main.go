package gcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type GCloudClient struct {
	Client *storage.Client
}

// SaveFile save file on a bucket
func (c *GCloudClient) SaveFile(bucket, fileName, content string) error {
	ctx := context.Background()
	w := c.Client.Bucket(bucket).Object(fileName).NewWriter(ctx)
	w.Write([]byte(content))
	if err := w.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}
	return nil
}

// ReadFile read file from bucket
func (c *GCloudClient) ReadFile(bucket, fileName string) (string, error) {
	ctx := context.Background()
	r, err := c.Client.Bucket(bucket).Object(fileName).NewReader(ctx)
	if err != nil {
		return "", fmt.Errorf("Object(%q).NewReader: %v", fileName, err)
	}
	defer r.Close()
	b, err := ioutil.ReadAll(r)
	return string(b), err
}

// InitializeGcloudClient initialize the GCloud client
//    gcp: GOOGLE_APPLICATION_CREDENTIALS: receive the URL for the file with the credentials
//    gcc: GCLOUD_CREDENTIALS: receive the content of the file with the credentials
func InitializeGcloudClient(gpc, gcc string) *storage.Client {
	if gpc != "" {
		ctx := context.Background()
		jsonFile, err := os.Open(gpc)
		credentialJSON := make(map[string]interface{})
		// Remove scapes
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &credentialJSON)

		credential, _ := json.Marshal(credentialJSON)
		client, err := storage.NewClient(ctx, option.WithCredentialsJSON(credential))
		if err != nil {
			log.Fatal("Couldn't connect to storage service", err)
		}
		return client
	}

	credentialJSON := make(map[string]interface{})
	// Remove scapes
	json.Unmarshal([]byte(strings.ReplaceAll(gcc, "'", "")), &credentialJSON)
	ctx := context.Background()

	credential, _ := json.Marshal(credentialJSON)
	client, _ := storage.NewClient(ctx, option.WithCredentialsJSON(credential))
	return client

}
