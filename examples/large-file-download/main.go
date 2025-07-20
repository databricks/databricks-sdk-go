package main

import (
	"context"
	"fmt"
	"io"
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
		Level: logger.LevelDebug,
	}

	// File to download
	remoteFilePath := "/Volumes/parth-testing/default/parth_files_api/large_random_file.bin"
	localFilePath := "downloaded_large_file.bin"

	// Delete the local file if it exists
	if _, err := os.Stat(localFilePath); err == nil {
		fmt.Printf("Deleting existing file: %s\n", localFilePath)
		if err := os.Remove(localFilePath); err != nil {
			log.Fatalf("Failed to delete existing file: %v", err)
		}
	}

	fmt.Printf("Downloading file: %s\n", remoteFilePath)
	fmt.Printf("Saving to: %s\n", localFilePath)
	fmt.Println()

	start := time.Now()

	// Download the file
	response, err := w.Files.Download(ctx, files.DownloadRequest{
		FilePath: remoteFilePath,
	})
	if err != nil {
		log.Fatalf("Failed to download file: %v", err)
	}
	defer response.Contents.Close()

	// Create the local file
	file, err := os.Create(localFilePath)
	if err != nil {
		log.Fatalf("Failed to create local file: %v", err)
	}
	defer file.Close()

	// Copy the downloaded content to the local file
	written, err := io.Copy(file, response.Contents)
	if err != nil {
		log.Fatalf("Failed to write to local file: %v", err)
	}

	duration := time.Since(start)
	fmt.Printf("\nFile downloaded successfully!\n")
	fmt.Printf("Local file: %s\n", localFilePath)
	fmt.Printf("Size: %d bytes (%.2f GB)\n", written, float64(written)/(1024*1024*1024))
	fmt.Printf("Time taken: %v\n", duration)
	fmt.Printf("Download speed: %.2f MB/s\n", float64(written)/(1024*1024)/duration.Seconds())

	// Verify the download by checking local file size
	fmt.Println("\nVerifying download...")
	fileInfo, err := os.Stat(localFilePath)
	if err != nil {
		log.Printf("Warning: Could not verify local file: %v", err)
	} else {
		if fileInfo.Size() == written {
			fmt.Println("✅ Local file size matches downloaded size!")
		} else {
			fmt.Printf("⚠️  File size mismatch! Downloaded: %d, Local: %d\n", written, fileInfo.Size())
		}
	}

	// Compare with original file if it exists
	originalFilePath := "large_random_file.bin"
	if _, err := os.Stat(originalFilePath); err == nil {
		fmt.Println("\nComparing with original file...")
		if compareFiles(originalFilePath, localFilePath) {
			fmt.Println("✅ File contents match the original file!")
		} else {
			fmt.Println("❌ File contents do not match the original file!")
		}
	} else {
		fmt.Printf("\n⚠️  Original file %s not found, skipping content comparison\n", originalFilePath)
	}
}

// compareFiles compares two files byte by byte
func compareFiles(file1, file2 string) bool {
	f1, err := os.Open(file1)
	if err != nil {
		log.Printf("Failed to open file %s: %v", file1, err)
		return false
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		log.Printf("Failed to open file %s: %v", file2, err)
		return false
	}
	defer f2.Close()

	// Compare file sizes first
	stat1, err := f1.Stat()
	if err != nil {
		log.Printf("Failed to get file1 stats: %v", err)
		return false
	}

	stat2, err := f2.Stat()
	if err != nil {
		log.Printf("Failed to get file2 stats: %v", err)
		return false
	}

	if stat1.Size() != stat2.Size() {
		log.Printf("File sizes differ: %s (%d bytes) vs %s (%d bytes)", file1, stat1.Size(), file2, stat2.Size())
		return false
	}

	// Compare contents byte by byte
	const bufferSize = 64 * 1024 // 64KB buffer
	buf1 := make([]byte, bufferSize)
	buf2 := make([]byte, bufferSize)

	for {
		n1, err1 := f1.Read(buf1)
		if err1 != nil && err1 != io.EOF {
			log.Printf("Error reading file1: %v", err1)
			return false
		}

		n2, err2 := f2.Read(buf2)
		if err2 != nil && err2 != io.EOF {
			log.Printf("Error reading file2: %v", err2)
			return false
		}

		if n1 != n2 {
			log.Printf("Read different number of bytes: %d vs %d", n1, n2)
			return false
		}

		if n1 == 0 {
			break // Both files reached EOF
		}

		// Compare the bytes read
		for i := 0; i < n1; i++ {
			if buf1[i] != buf2[i] {
				log.Printf("Byte mismatch at position %d: %d vs %d", i, buf1[i], buf2[i])
				return false
			}
		}
	}

	return true
}
