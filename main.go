package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/bogem/id3v2"
	"github.com/mmcdole/gofeed"
	"gopkg.in/cheggaaa/pb.v1"
)

func main() {
	// Parse CLI args
	downloadDirectory := flag.String("dir", "musicforprogramming", "Directory into which the tracks will be downloaded")
	flag.Parse()
	if *downloadDirectory == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	createDir(*downloadDirectory)
	getCover(*downloadDirectory)
	tracks := getTracks()
	for trackIndex, URL := range tracks.URL {
		getTrack(*downloadDirectory, tracks.Title[trackIndex], URL)
	}
}

type tracks struct {
	Title []string
	URL   []string
}

// Borrowed from https://stackoverflow.com/questions/34816489/reverse-slice-of-strings
func reverseStringSlice(ss []string) []string {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
	return ss
}

// getTracks gets titles and URLs to tracks from RSS feed
func getTracks() (tracks tracks) {
	log.Println("Downloading and parsing tracks index...")
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
	// Invert both slices
	// Having index 0 be the first track eases testing and makes more sense to the user
	tracks.Title = reverseStringSlice(tracks.Title)
	tracks.URL = reverseStringSlice(tracks.URL)
	return tracks
}

// getCover downloads folder.jpg (album cover image) if it doesn't exist
func getCover(destDirectory string) {
	coverPath := filepath.Join(destDirectory, "folder.jpg")
	if _, err := os.Stat(coverPath); os.IsNotExist(err) {
		log.Printf("Album cover does not exist, downloading...")

		file, err := os.Create(coverPath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		resp, err := http.Get("https://musicforprogramming.net/img/folder.jpg")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		// See https://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
		// Stream HTTP response into file
		n, err := io.Copy(file, resp.Body)
		if err != nil {
			panic(err)
		}
		// We don't use n
		_ = n
		log.Println("Album cover downloaded!")
	}
}

// createDir checks whether the directory exists and creates it if it doesn't
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

// Add album metadata to the track supplied.
// This is needed because the downloaded files don't contain this data.
func setAlbum(file string) {
	mp3File, err := id3v2.Open(file, id3v2.Options{Parse: true})
	if err != nil {
		panic(err)
	}
	defer mp3File.Close()
	mp3File.SetAlbum("Music For Programming")
}

// Download supplied track to supplied directory if it doesn't exist on disk
func getTrack(destDirectory string, title string, URL string) {
	trackPath := filepath.Join(destDirectory, title+".mp3")
	if _, err := os.Stat(trackPath); os.IsNotExist(err) {
		log.Printf("Track '" + title + "' does not exist on disk, downloading... ")
		file, err := os.Create(trackPath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		resp, err := http.Get(URL)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		// Prepare to show a progress bar
		// Start by getting the size of the file we'll be downloading
		size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
		// Set up our progress bar
		// See https://github.com/cheggaaa/pb/#progress-bar-for-io-operations
		bar := pb.New(size).SetUnits(pb.U_BYTES)
		bar.Start()
		writer := io.MultiWriter(file, bar)
		// See https://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go
		// Stream HTTP response into file
		n, err := io.Copy(writer, resp.Body)
		if err != nil {
			panic(err)
		}
		bar.Finish()
		// We don't use n
		_ = n
		setAlbum(trackPath)
		log.Printf("Track '" + title + "' downloaded!")
	}
}
