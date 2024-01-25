package internal

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

func storeGCS(content io.ReadCloser, bucketName, fileName string) error {
	// Create GCS connection
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	// Connect to bucket
	bucket := client.Bucket(bucketName)

	// Setup the GCS object with the filename to write to
	obj := bucket.Object(fileName)

	// w implements io.Writer.
	w := obj.NewWriter(ctx)

	// Copy file into GCS
	if _, err := io.Copy(w, content); err != nil {
		return fmt.Errorf("failed to copy to bucket: %v", err)
	}

	// Close, just like writing a file. File appears in GCS after
	if err := w.Close(); err != nil {
		return fmt.Errorf("failed to close: %v", err)
	}

	return nil
}
