package main

import (
	xkcd "book/ch4/xkcd/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// TODO
// take in int as arg
// decode json file into structs
// make map out of comics with num as key
func main() {
	// xkcd.FetchComics()

	argCount := len(os.Args[1:])
	if argCount > 1 {
		log.Fatal("Too Many Args")
	}
	comicId, _ := strconv.Atoi(os.Args[1])
	jsonFile, err := os.Open("comics.json")
	if err != nil {
		log.Fatalf("DUDE WHAT %s\n", err)
	}
	defer jsonFile.Close()

	//Read Opened Json FIle
	bytes, _ := ioutil.ReadAll(jsonFile)

	var comics xkcd.ComicResult
	json.Unmarshal(bytes, &comics)

	// var foundComic xkcd.Comic
	if comicId > comics.TotalCount {
		log.Fatalf("Id %d not found\n", comicId)
	}
	for _, comic := range comics.Comics {
		if comic.Num == comicId {
			// foundComic = comic
			fmt.Printf("%s\n", comic.Transcript)
			break
		}
	}

}
