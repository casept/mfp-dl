package main

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestGetCover(t *testing.T) {
	// TODO: Mock storage so we don't have to use a temp dir
	tempDataDir := filepath.Join("testdata", "tmp")
	os.MkdirAll(tempDataDir, 0775)
	getCover(tempDataDir)
	coverPath := filepath.Join(tempDataDir, "folder.jpg")
	// TODO: Fix this
	useless, err := os.Stat(coverPath)
	if err != nil {
		t.Errorf("Test failed, folder.jpg could not be found")
	}
	_ = useless
}

func TestGetTracks(t *testing.T) {
	tracks := getTracks()
	trackIndex := 0
	trackTitle := "Episode 01: Datassette"
	if tracks.Title[trackIndex] != trackTitle {
		t.Errorf("Title does not match! Expected " + trackTitle + ", got " + tracks.Title[trackIndex])
	}
	// Make sure the URL we retrieved points to a valid file
	resp, _ := http.Get(tracks.URL[trackIndex])
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Retrieved URL for track " + tracks.Title[trackIndex] + " is invalid! Got status " + string(resp.StatusCode) + ", expected " + string(http.StatusOK))
	}
}
