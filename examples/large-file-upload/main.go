package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/service/files"
)

func main() {
	// Configuration - update this with your profile
	cfg := &databricks.Config{
		Profile: "dbc-1232e87d-9384", // Update this to your profile
	}

	// Create workspace client
	w := databricks.Must(databricks.NewWorkspaceClient(cfg))
	ctx := context.Background()

	// Set up logging
	logger.DefaultLogger = &logger.SimpleLogger{
		Level: logger.LevelInfo,
	}

	// File to upload
	localFilePath := "large_random_file.bin"
	remoteFilePath := "/Volumes/parth-testing/default/parth_files_api/large_random_file.bin"

	// Check if local file exists
	if _, err := os.Stat(localFilePath); os.IsNotExist(err) {
		log.Fatalf("Local file %s does not exist. Please run the file generator first.", localFilePath)
	}

	// Get file info
	fileInfo, err := os.Stat(localFilePath)
	if err != nil {
		log.Fatalf("Failed to get file info: %v", err)
	}

	fileSize := fileInfo.Size()
	fmt.Printf("Uploading file: %s\n", localFilePath)
	fmt.Printf("File size: %d bytes (%.2f GB)\n", fileSize, float64(fileSize)/(1024*1024*1024))
	fmt.Printf("Remote path: %s\n", remoteFilePath)
	fmt.Println()

	start := time.Now()

	// Open the local file
	file, err := os.Open(localFilePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Upload the file
	err = w.Files.Upload(ctx, files.UploadRequest{
		FilePath: remoteFilePath,
		Contents: file,
	})
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	duration := time.Since(start)
	fmt.Printf("\nFile uploaded successfully!\n")
	fmt.Printf("Remote path: %s\n", remoteFilePath)
	fmt.Printf("Time taken: %v\n", duration)
	fmt.Printf("Upload speed: %.2f MB/s\n", float64(fileSize)/(1024*1024)/duration.Seconds())

	// Verify the upload by getting file metadata
	fmt.Println("\nVerifying upload...")
	fileInfoResponse, err := w.Files.GetMetadataByFilePath(ctx, remoteFilePath)
	if err != nil {
		log.Printf("Warning: Could not verify file info: %v", err)
	} else {
		fmt.Printf("Verified file size: %d bytes\n", fileInfoResponse.ContentLength)
		if fileInfoResponse.ContentLength == fileSize {
			fmt.Println("✅ File size matches!")
		} else {
			fmt.Printf("⚠️  File size mismatch! Expected: %d, Got: %d\n", fileSize, fileInfoResponse.ContentLength)
		}
	}
}
