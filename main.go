package main

import (
	"net/http"
	"log"
	"path/filepath"
	"os"
	"io"

	"github.com/mmcdole/gofeed"
)
type tracks struct {
	Title []string
	URL []string
}

// Get titles and URLs to tracks from RSS feed
func getTracks() (tracks tracks) {
	log.Printf("Downloading and parsing tracks index...")
	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL("https://musicforprogramming.net/rss.php")
	if err != nil {
		panic(err)
	}
	// Read and append track title and URL to tracks
	for _, item := range feed.Items {
		tracks.Title = append(tracks.Title, item.Title)
		for _, enclosure := range item.Enclosures {
			tracks.URL = append(tracks.URL, enclosure.URL)
		}
	}
	return tracks
}

// Download folder.jpg (album cover image) if it doesn't exist
func getCover(destDirectory string) {
	coverPath := filepath.Join(destDirectory, "folder.jpg") 
	if _, err := os.Stat(coverPath); os.IsNotExist(err) {
		log.Printf("Album cover does not exist, downloading...")
		
		file, err :=os.Create(coverPath)
		defer file.Close()
		if err != nil {
			panic(err)
		}
		
		resp, err := http.Get("https://musicforprogramming.net/img/folder.jpg")
		defer resp.Body.Close()
		if err != nil {
			panic(err)
		}
		// See https://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
		// Stream HTTP response into file
		n, err := io.Copy(file, resp.Body)
		// We don't use n
		_ = n
		log.Printf("Album cover downloaded!")
	}
}

// Check if target directory exists, create if it doesn't
func createDir(destDirectory string) {
	if _, err := os.Stat(destDirectory); os.IsNotExist(err) {
		log.Printf("Directory " + destDirectory + " does not exist, creating...")
		err := os.MkdirAll(destDirectory, 0775)
		if err != nil {
			panic(err)
		}
		log.Printf("Directory " + destDirectory + " created!")
	}
}

// Download supplied track to supplied directory if it doesn't exist on disk
func getTrack(destDirectory string, title string, URL string) {
	
}
func main() {
	createDir("test")
}
