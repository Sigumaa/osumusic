package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var directory string
	fmt.Print("Enter directory: ")
	s := bufio.NewScanner(os.Stdin)
	if s.Scan() {
		directory = s.Text()
	}

	dirs, err := os.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	var cnt int
	var songs string
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		song, err := SearchSong(directory + "/" + dir.Name())
		if err != nil {
			continue
		}

		songs += song.Title + " - " + song.Artist + "\n"
		cnt++
	}
	fmt.Println(songs)
	fmt.Println("you have", cnt, "songs")
}
