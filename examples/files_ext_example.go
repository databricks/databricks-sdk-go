package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/files"
)

func main() {
	// Example 1: Basic setup and usage
	fmt.Println("=== Enhanced Files API Example ===")

	// Create configuration
	cfg := &config.Config{
		Host:  os.Getenv("DATABRICKS_HOST"),
		Token: os.Getenv("DATABRICKS_TOKEN"),
	}

	// Create client
	databricksClient, err := client.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create enhanced Files API
	filesExt := files.NewFilesExt(databricksClient)

	// Example 2: Upload a small file (one-shot upload)
	fmt.Println("\n--- Uploading small file ---")
	smallContent := strings.NewReader("Hello, World! This is a small file.")
	err = filesExt.Upload(context.Background(), files.UploadRequest{
		FilePath:  "/Volumes/example-catalog/example-schema/example-volume/small-file.txt",
		Contents:  io.NopCloser(smallContent),
		Overwrite: true,
	})
	if err != nil {
		log.Printf("Small file upload failed: %v", err)
	} else {
		fmt.Println("✓ Small file uploaded successfully")
	}

	// Example 3: Upload a large file (multipart upload)
	fmt.Println("\n--- Uploading large file ---")
	largeContent := strings.NewReader(strings.Repeat("Large file content for multipart upload demonstration. ", 100000))
	err = filesExt.Upload(context.Background(), files.UploadRequest{
		FilePath:  "/Volumes/example-catalog/example-schema/example-volume/large-file.txt",
		Contents:  io.NopCloser(largeContent),
		Overwrite: true,
	})
	if err != nil {
		log.Printf("Large file upload failed: %v", err)
	} else {
		fmt.Println("✓ Large file uploaded successfully")
	}

	// Example 4: Download a file with resilient download
	fmt.Println("\n--- Downloading file ---")
	response, err := filesExt.Download(context.Background(), files.DownloadRequest{
		FilePath: "/Volumes/example-catalog/example-schema/example-volume/small-file.txt",
	})
	if err != nil {
		log.Printf("Download failed: %v", err)
	} else {
		defer response.Contents.Close()

		content, err := io.ReadAll(response.Contents)
		if err != nil {
			log.Printf("Failed to read content: %v", err)
		} else {
			fmt.Printf("✓ Downloaded file successfully\n")
			fmt.Printf("  Content length: %d bytes\n", len(content))
			fmt.Printf("  Content: %s\n", string(content))
		}
	}

	// Example 5: Streaming download with recovery
	fmt.Println("\n--- Streaming download ---")
	streamResponse, err := filesExt.Download(context.Background(), files.DownloadRequest{
		FilePath: "/Volumes/example-catalog/example-schema/example-volume/large-file.txt",
	})
	if err != nil {
		log.Printf("Streaming download failed: %v", err)
	} else {
		defer streamResponse.Contents.Close()

		buffer := make([]byte, 1024)
		totalBytes := 0
		chunks := 0

		for {
			n, err := streamResponse.Contents.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Streaming read error: %v", err)
				break
			}

			totalBytes += n
			chunks++

			// Process the chunk (in this example, just count bytes)
			_ = buffer[:n]
		}

		fmt.Printf("✓ Streaming download completed\n")
		fmt.Printf("  Total bytes: %d\n", totalBytes)
		fmt.Printf("  Chunks processed: %d\n", chunks)
	}

	// Example 6: Configuration demonstration
	fmt.Println("\n--- Configuration ---")
	config := files.DefaultUploadConfig()
	fmt.Printf("Default configuration:\n")
	fmt.Printf("  Min stream size: %d bytes (%d MB)\n",
		config.MultipartUploadMinStreamSize,
		config.MultipartUploadMinStreamSize/(1024*1024))
	fmt.Printf("  Chunk size: %d bytes (%d MB)\n",
		config.MultipartUploadChunkSize,
		config.MultipartUploadChunkSize/(1024*1024))
	fmt.Printf("  Batch URL count: %d\n", config.MultipartUploadBatchURLCount)
	fmt.Printf("  Max retries: %d\n", config.MultipartUploadMaxRetries)
	fmt.Printf("  Download max recovers: %d\n", config.FilesAPIClientDownloadMaxTotalRecovers)

	// Example 7: Custom configuration
	demonstrateCustomConfig()

	fmt.Println("\n=== Example completed ===")
}

// Helper function to demonstrate custom configuration
func demonstrateCustomConfig() {
	fmt.Println("\n--- Custom Configuration Example ---")

	customConfig := &files.UploadConfig{
		MultipartUploadMinStreamSize:                             50 * 1024 * 1024, // 50MB
		MultipartUploadChunkSize:                                 50 * 1024 * 1024, // 50MB
		MultipartUploadBatchURLCount:                             5,
		MultipartUploadMaxRetries:                                5,
		MultipartUploadSingleChunkUploadTimeoutSeconds:           600,
		MultipartUploadURLExpirationDuration:                     time.Hour * 2,
		FilesAPIClientDownloadMaxTotalRecovers:                   15,
		FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 5,
	}

	fmt.Printf("Custom configuration:\n")
	fmt.Printf("  Min stream size: %d bytes (%d MB)\n",
		customConfig.MultipartUploadMinStreamSize,
		customConfig.MultipartUploadMinStreamSize/(1024*1024))
	fmt.Printf("  Chunk size: %d bytes (%d MB)\n",
		customConfig.MultipartUploadChunkSize,
		customConfig.MultipartUploadChunkSize/(1024*1024))
	fmt.Printf("  Batch URL count: %d\n", customConfig.MultipartUploadBatchURLCount)
	fmt.Printf("  Max retries: %d\n", customConfig.MultipartUploadMaxRetries)
	fmt.Printf("  Download max recovers: %d\n", customConfig.FilesAPIClientDownloadMaxTotalRecovers)
}
