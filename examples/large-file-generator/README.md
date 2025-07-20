# Large File Generator

This program creates a 4.5 GB file filled with cryptographically secure random data.

## Features

- Generates exactly 4.5 GB of random data
- Uses cryptographically secure random number generation (`crypto/rand`)
- Writes data in 1MB chunks for efficient memory usage
- Shows progress every 100MB
- Displays final statistics including write speed

## Usage

```bash
cd examples/large-file-generator
go run main.go
```

The program will create a file named `large_random_file.bin` in the current directory.

## Output

The program will show:
- Progress updates every 100MB
- Final file size and creation time
- Write speed in MB/s

## Requirements

- Go 1.16 or later
- Sufficient disk space (at least 4.5 GB free)

## Notes

- The file contains cryptographically secure random data, not pseudo-random data
- The program uses efficient buffered I/O to minimize memory usage
- All data is synced to disk before completion 