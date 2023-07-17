package main

import (
	xkcd "book/ch4/xkcd/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// TODO
// take in int as arg
// decode json file into structs
// make map out of comics with num as key
func main() {
	// xkcd.FetchComics()

	comicId := os.Args[1]
	fmt.Printf("COMIC ID: %s\n", comicId)
  idInt, _ := strconv.Atoi(comicId)
	jsonFile, err := os.Open("comics.json")
	if err != nil {
		fmt.Printf("DUDE WHAT %s\n", err)
	}
	fmt.Print("SUCCESSFULLYY OPENED JSON FILE\n")
	defer jsonFile.Close()

	//Read Opened Json FIle
	bytes, _ := ioutil.ReadAll(jsonFile)
	fmt.Print("SUCCESSFULLYY READ JSON FILE\n")

	var comics xkcd.ComicResult
	json.Unmarshal(bytes, &comics)
	fmt.Printf("COUNT: %d\n", comics.TotalCount)

	// var foundComic xkcd.Comic
	for _, comic := range comics.Comics {
		if comic.Num == idInt {
			// foundComic = comic
		fmt.Printf("LOOPING DAWG: %d\n", comic.Num)
			fmt.Printf("YAY I FOIND IT - %s\n", comic.Transcript)
      break
		}
	}

}
