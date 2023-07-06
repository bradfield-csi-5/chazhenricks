package main

import (
	"ch1/ch1/lissajous"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	// server1()
	// server2()
	server3()
}

func simple_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path:  %q\n", r.URL.Path)
}

func server1() {
	http.HandleFunc("/", simple_handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func server2() {
	http.HandleFunc("/", mutex_handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func server3() {
	http.HandleFunc("/", reportHandler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/gif", gifHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mutex_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URL: %s, Raw Path: %s  HostName:%s \n", r.URL, r.URL.RawPath, r.URL.Hostname())
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL Path: %q\n", r.URL.Path)
}

func gifHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	cycles, err := r.Form["cycles"]
	if err != nil {
		log.Print(err)
  }

  
	lissajous.Lissajous(w, strconv.Atoi(cycles))

}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count: %d\n", count)
	mu.Unlock()
}

func reportHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("yarrrr")
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "Remote Addr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
