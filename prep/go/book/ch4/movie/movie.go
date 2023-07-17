package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphery Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve Mcqueen", "Jackqueline Bisset"}},
}

func main() {
	//marshall to JSON
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s\n", err)
	}
	fmt.Printf("%s\n", data)

	//marshall to readable json
	readable, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshalling failed: %s\n", err)
	}
	fmt.Printf("%s\n", readable)

  //titles is a slice of anon structs whos only field is Title
	var titles []struct{ Title string }
  //Unmarshall takes the JSON data and a pointer for where to dump it
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON Unmarshal failed: %s\n", err)
	}
	fmt.Println(titles)

}
