# Large File Download

This program downloads a large file from Databricks using the Files API and saves it to disk.

## Features

- Downloads files from Databricks using the Files API
- Supports Unity Catalog volumes
- Saves downloaded content to a local file
- Shows download progress and statistics
- Verifies download by checking file size
- **Content validation**: Compares downloaded file with original local file
- **Clean start**: Automatically deletes existing downloaded file before starting
- Uses efficient streaming download

## Prerequisites

1. **Databricks Configuration**: Make sure you have a Databricks profile configured
2. **File Exists**: The file must exist in the specified remote location
3. **Unity Catalog Volume**: The source volume must be accessible

## Configuration

Update the configuration in `main.go`:

```go
cfg := &databricks.Config{
    Profile: "your-profile-name", // Update this to your profile
}
```

And update the remote file path:

```go
remoteFilePath := "/Volumes/your-catalog/your-schema/your-volume/your-file.bin"
```

You can also customize the local file name:

```go
localFilePath := "your-local-filename.bin"
```

## Usage

```bash
cd examples/large-file-download
go run main.go
```

The program will:
- **Delete existing downloaded file** (if present)
- Connect to your Databricks workspace
- Download the specified file from the remote location
- Save it to the local file system
- Show download statistics (time taken, speed)
- Verify the download by checking file size
- **Compare contents** with the original local file (if available)

## Output

The program will show:
- Confirmation of existing file deletion (if applicable)
- Remote file path being downloaded
- Local file path where it's being saved
- Download progress
- Final download statistics
- Verification results
- Content comparison results (if original file exists)

## Error Handling

The program includes comprehensive error handling for:
- Missing remote files
- Network connectivity issues
- Authentication problems
- File system errors
- Insufficient disk space

## Notes

- The download uses streaming to handle large files efficiently
- The program verifies the download by comparing file sizes
- **Content validation**: If `large_random_file.bin` exists locally, the program will perform a byte-by-byte comparison
- The comparison uses efficient buffered reading (64KB chunks) to handle large files
- **Clean start**: The program automatically deletes any existing `downloaded_large_file.bin` before starting
- Make sure you have sufficient disk space for the downloaded file
- The local file will be created in the current directory

## Example Workflow

1. **Upload a file** (using the upload program):
   ```bash
   cd examples/large-file-upload
   go run main.go
   ```

2. **Download the file**:
   ```bash
   cd examples/large-file-download
   go run main.go
   ```

This creates a complete round-trip workflow for testing large file operations with the Databricks Files API. 