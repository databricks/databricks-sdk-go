package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	// 4.5 GB in bytes
	fileSize = 4.5 * 1024 * 1024 * 1024
	// Buffer size for writing (1MB chunks)
	bufferSize = 1024 * 1024
)

func main() {
	start := time.Now()

	// Create the output file
	outputFile := "large_random_file.bin"
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	fmt.Printf("Creating %s (%d bytes)...\n", outputFile, fileSize)

	// Create a buffer for random data
	buffer := make([]byte, bufferSize)
	bytesWritten := int64(0)

	// Write random data in chunks
	for bytesWritten < fileSize {
		// Calculate how many bytes to write in this iteration
		remaining := fileSize - bytesWritten
		chunkSize := bufferSize
		if remaining < int64(bufferSize) {
			chunkSize = int(remaining)
		}

		// Generate random data
		_, err := rand.Read(buffer[:chunkSize])
		if err != nil {
			log.Fatalf("Failed to generate random data: %v", err)
		}

		// Write to file
		written, err := file.Write(buffer[:chunkSize])
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}

		bytesWritten += int64(written)

		// Progress indicator every 100MB
		if bytesWritten%(100*1024*1024) == 0 {
			progress := float64(bytesWritten) / float64(fileSize) * 100
			fmt.Printf("Progress: %.1f%% (%d MB written)\n", progress, bytesWritten/(1024*1024))
		}
	}

	// Ensure all data is written to disk
	err = file.Sync()
	if err != nil {
		log.Fatalf("Failed to sync file: %v", err)
	}

	duration := time.Since(start)
	fmt.Printf("\nFile created successfully!\n")
	fmt.Printf("File: %s\n", outputFile)
	fmt.Printf("Size: %d bytes (%.2f GB)\n", bytesWritten, float64(bytesWritten)/(1024*1024*1024))
	fmt.Printf("Time taken: %v\n", duration)
	fmt.Printf("Write speed: %.2f MB/s\n", float64(bytesWritten)/(1024*1024)/duration.Seconds())
}
