package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/madgeniusblink/parallel-downloader/internal/helper"
)

// Core downloading logic

// DownloadConfig represents the configuration for a download request.
type DownloadConfig struct {
	URL      string
	Chunks   int
	FilePath string
}

// Download initiates the file download based on the given configuration.
func Download(config DownloadConfig) error {
	// TODO: Implement the core downloading logic here

	contentLength, err := helper.GetFileSizeFromHeader(config.URL)

	if err != nil {
		log.Printf("Failed to get file size: %v", err)
		return err
	}
	// Calculating the chunk size
	chunkSize := contentLength / int64(config.Chunks)

	// Creating empty file of given size
	file, err := helper.CreateEmptyFile(config.FilePath, contentLength)
	if err != nil {
		log.Printf("Failed to create file: %v", err)
		return err
	}
	defer file.Close()

	// channel to capture any error from goroutines
	errChan := make(chan error, config.Chunks)

	// channel to signal main routine that all chunks are downloaded
	doneChan := make(chan bool, config.Chunks)

	for i := 0; i < config.Chunks; i++ {
		go func(i int) {
			start := int64(i) * chunkSize
			end := start + chunkSize - 1
			if i == config.Chunks-1 {
				end = contentLength - 1
			}

			req, err := http.NewRequest("GET", config.URL, nil)
			if err != nil {
				errChan <- err
				return
			}

			req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", start, end))
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				errChan <- err
				return
			}
			defer resp.Body.Close()

			// Writing to file
			file.Seek(int64(start), 0)
			_, err = io.Copy(file, resp.Body)
			if err != nil {
				errChan <- err
				return
			}
			doneChan <- true
		}(i)

	}

	// Waiting for all goroutines to complete
	for i := 0; i < config.Chunks; i++ {
		select {
		case err := <-errChan:
			return err
		case <-doneChan:
		}
	}

	return nil
}
