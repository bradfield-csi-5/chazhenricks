package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}


func main(){
  addEdge("shit", "fuck")

  addEdge("shit", "balls")
  ok := hasEdge("shit", "fuck")
  fmt.Println(ok)
  ok = hasEdge("shit", "balls")
  fmt.Println(ok)
}
