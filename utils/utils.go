package utils

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

// NewStorageClient creates a new storage client.
// if the function is running locally, it will use the creds.json file.
// if the function is running in the cloud, it will use the default credentials.
func NewStorageClient(ctx context.Context) (*storage.Client, error) {
	var client *storage.Client
	var err error
	if os.Getenv("USE_CREDS_FILE") != "true" {
		client, err = storage.NewClient(ctx)
	} else {
		client, err = storage.NewClient(ctx, option.WithCredentialsFile("./creds.json"))
	}

	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %w", err)
	}
	return client, nil
}
