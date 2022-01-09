//A minimum spanning tree is a spanning tree in which the sum of the weight of the edges is as minimum as possible.
package minimumspanningtree

import "fmt"

const INF = 9999999 //infinity ?
func Prim(graph [][]int) {

	//choose random vertex
	//choose the shortest edge ->add it
	//keep adding shortest not yet added
	//visited = visisted + vertex
	//tree = //add to tree

	//TODO convert between adjacency matrix and adjacency list
	numVertexes := len(graph)

	selected := []bool{false, false, false, false, false} //TODO this should be numvertexes long
	numberEdges := 0
	selected[0] = true //select the first vertex
	for numberEdges < numVertexes-1 {
		minimum := INF
		x := 0
		y := 0
		for i := 0; i < numVertexes; i++ { //go through the selected vertexes
			if selected[i] {
				for j := 0; j < numVertexes; j++ {
					if !selected[j] && graph[i][j] != 0 { //vertexes with an edge with the current one, which aren't selected
						if minimum > graph[i][j] {
							minimum = graph[i][j]
							x = i //store minimum so far
							y = j
						}

					}
				}
			}
		}
		fmt.Printf("%d-%d:%d\n", x, y, graph[x][y])
		selected[y] = true
		numberEdges++
	}

}
