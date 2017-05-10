// Copyright 2016 Albert Nigmatzianov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package id3v2 is the ID3 parsing and writing library for Go.
package id3v2

import (
	"os"

	"github.com/bogem/id3v2/util"
)

// Available picture types for picture frame.
const (
	PTOther = iota
	PTFileIcon
	PTOtherFileIcon
	PTFrontCover
	PTBackCover
	PTLeafletPage
	PTMedia
	PTLeadArtistSoloist
	PTArtistPerformer
	PTConductor
	PTBandOrchestra
	PTComposer
	PTLyricistTextWriter
	PTRecordingLocation
	PTDuringRecording
	PTDuringPerformance
	PTMovieScreenCapture
	PTBrightColouredFish
	PTIllustration
	PTBandArtistLogotype
	PTPublisherStudioLogotype
)

// Available encodings.
var (
	// ISO-8859-1.
	ENISO = util.Encoding{
		Key:              0,
		TerminationBytes: []byte{0},
	}

	// UTF-16 encoded Unicode with BOM.
	ENUTF16 = util.Encoding{
		Key:              1,
		TerminationBytes: []byte{0, 0},
	}

	// UTF-16BE encoded Unicode without BOM.
	ENUTF16BE = util.Encoding{
		Key:              2,
		TerminationBytes: []byte{0, 0},
	}

	// UTF-8 encoded Unicode.
	ENUTF8 = util.Encoding{
		Key:              3,
		TerminationBytes: []byte{0},
	}

	Encodings = []util.Encoding{ENISO, ENUTF16, ENUTF16BE, ENUTF8}
)

// Open opens file with name and passes it to OpenFile.
func Open(name string, opts Options) (*Tag, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return OpenFile(file, opts)
}

// OpenFile parses opened file and finds tag in it considering opts.
// If there is no tag in file, OpenFile will create new one with version ID3v2.4.
func OpenFile(file *os.File, opts Options) (*Tag, error) {
	return parseTag(file, opts)
}