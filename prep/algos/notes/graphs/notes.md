# Graphs (sections 1-3, 5-6)

## Introduction 
-   a more general tree (or a tree is a specific kind of graph)

- **vertex** - aka "node". Can have a name (called "key") and additional info ("payload")
- **edge** - connects two vertices. may be one way (like a tree) or two ways 
- **weight** - edges can be weighted to show cost from one vertex to another. Think "traffic" in a graph of two destinations. 
- **path** - sequence of vertices that are connected by edges. 
- **cycle** - in a directed graph, a path that can start and end on the same vertex. 


#### Abstract Data Type 
- graph() - creates a new, empty graph
- add_vertex(vertex) - adds an instance of Vertex to the graph
- add_edge(from, to) adds a new directed graph edge to the graph that connects two vertices 
- get_vertex(key)- finds vertex in the graph named "key"
- get_vertices - returns list of all vertices 
- in - returns True or False from expressions like `vertex in Graph` 


## Representing A Graph

- the most accurate way to represent a graph would be with a matrix of each vertex in the graph across the X and Y axis. 
- a weight is filled in at each cell where there is an edge between the x vertex and y vertex
- problem with this is the vast majority of graphs would be "sparse" meaning theres probs more empty space than edges. 
- better way is with Adjacency List 

### Adjacency List 
- Main dict that holds all the vertices as keys
- values are another dict that holds all edges that vertex has.
 

#### Class Based Graph 
- see `adjacency_list.py` for the code version of a class based list 

### Using Dictionaries Directly 
- you can also just us a straight up dictionary of dictionaries. 
- we'll do this for the rest of the shit. lol. 

```python 

{
    0: {1: 5, 5: 2},
    1: {2: 4},
    2: {3: 9},
    3: {4: 7, 5: 3},
    4: {0: 1},
    5: {4: 8}
}
```
