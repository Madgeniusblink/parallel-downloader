package main

import (
	"flag"
	"fmt"

	"github.com/madgeniusblink/parallel-downloader/internal/downloader"
)

func main() {
	url := flag.String("url", "", "URL to download")
	flag.Parse()

	if *url == "" {
		fmt.Println("Please provide a URL using the -url flag.")
		return
	}

	fmt.Println("Downloader Daemon is running...")
	fmt.Printf("Downloading %s...\n", *url)

	// TODO: Argument parsing and actual downloading logic
	config := downloader.DownloadConfig{
		URL:      *url,
		Chunks:   5,
		FilePath: "downloads/downloaded.file",
	}
	err := downloader.Download(config)
	if err != nil {
		fmt.Printf("Error downloading %s: %s\n", *url, err)
		return
	}

	fmt.Println("Downloader Daemon is shutting down...")
}
