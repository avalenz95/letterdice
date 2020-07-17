Letter Dice is my CS 2.2 Final Project

Letter Dice is simply a small program meant to demonstrate a use of a Network Flow Algorithm using the Edmonds-Karp Algorithm which is a derivation of the Ford-Fulkerson algorthm that uses BFS to find paths.

To run the program simply execute the given file with the command below

```bash
 ./letterdice
```

If you would like to add more inputs go into the main.go file and add or remove elements to the input struct in the main function then simply run type in the terminal command.
```bash
go run main.go
```


The current solution I have isn't very scalable in fact it takes a few seconds with a couple of inputs already. If I wanted to make it slightly more scalable I could change how I build the graph for each subsequent word. Instead of complete demolishing the graph and rebuilding it again I could simply just replace the Letter nodes with the new word nodes.

The run time for this algorithm is O(V E^2) since we have to traverse each edge and for each edge we traverse it's backedge as well while doing the bfs call.