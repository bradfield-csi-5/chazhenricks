package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	ComicURL   = "https://xkcd.com/"
	URLPostFix = "/info.0.json"
)

type Comic struct {
	Month      string
	Num        int
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
}

type ComicResult struct {
	TotalCount int `json:"total_count"`
	Comics     []*Comic
}

// TODO
// loop through and fetch all comics
// Make an struct that holds a slice of comics
// for each response, add new comic to the slice
//encode as JSON and store result in a .JSON file

func FetchComics() {
	i := 1

	var result ComicResult
	for true {
		strint := strconv.Itoa(i)
		resp, err := http.Get(ComicURL + strint + URLPostFix)
		if err != nil {
			fmt.Printf("Reached end of comics - found %d - %s\n", i, err)
			break
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK && i != 404 {
			fmt.Printf("Reached end of comics - found %d\n", i)
			break
		}

		var comic Comic
		if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
			fmt.Printf("idk we got fucked up somewhere: %s\n", err)
		}

		fmt.Printf("%s - %d\n", comic.SafeTitle, comic.Num)
		i++

		result.Comics = append(result.Comics, &comic)
		result.TotalCount = i
	}

  file,err := json.Marshal(result)
  if err != nil{
    fmt.Printf("Error saving to file: %q\n", err)
  }
  ioutil.WriteFile("comics.json", file, 0777)
}

