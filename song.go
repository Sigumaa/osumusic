package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type Song struct {
	Title     string
	Artist    string
	AudioPath string
}

func NewSong() *Song {
	song := &Song{}
	return song
}

func SearchSong(folder string) (*Song, error) {
	song := NewSong()

	osuFiles, err := filepath.Glob(filepath.Join(folder, "*.osu"))
	if err != nil {
		return nil, err
	}

	if len(osuFiles) < 1 {
		err = errors.New("no osu file found")
		return nil, err
	}

	properties := make(map[string]string)
	fileData, err := os.ReadFile(osuFiles[0])
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(fileData), "\n")
	for _, line := range lines {
		if !strings.Contains(line, ":") {
			continue
		}

		splitted := strings.SplitN(line, ":", 2)
		key := strings.TrimSpace(splitted[0])
		value := strings.TrimSpace(splitted[1])
		properties[key] = value
	}

	if t, ok := properties["TitleUnicode"]; ok {
		song.Title = t
	} else {
		song.Title = properties["Title"]
	}

	if a, ok := properties["ArtistUnicode"]; ok {
		song.Artist = a
	} else {
		song.Artist = properties["Artist"]
	}

	song.AudioPath = filepath.Join(folder, properties["AudioFilename"])

	return song, nil
}
