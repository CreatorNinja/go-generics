package azureblob

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

var AzClient *azblob.Client

type Client struct {
	client *azblob.Client
}

func InitializeClient(url string) error {
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return err
	}

	AzClient, err = azblob.NewClient(url, credential, nil)
	return err
}

func NewClient(url string) (*Client, error) {
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	client, err := azblob.NewClient(url, credential, nil)
	return &Client{client: client}, err
}

func (c *Client) WriteFile(body []byte, name, container string) {
	ctx := context.Background()

	// Upload to data to blob storage
	log.Printf("Uploading file %s to blob %s\n", name, container)
	_, err := c.client.UploadBuffer(ctx, container, name, body, &azblob.UploadBufferOptions{})
	if err != nil {
		log.Println("Error while file uploading into blob storage. ", err)
	}
	log.Printf("File %s uploaded\n", name)
}
