package planner

type vertex struct {
	name      string
	neighbors []*vertex
	visited   bool
}

func (vertex *vertex) addNeighbor(neighbor *vertex) {
	vertex.neighbors = append(vertex.neighbors, neighbor)
}

type graph struct {
	vertices []*vertex
}

func (graph *graph) getVertex(name string) *vertex {
	for _, actVertex := range graph.vertices {
		if actVertex.name == name {
			return actVertex
		}
	}

	return nil
}

func (graph *graph) addVertex(name string) *vertex {
	srcVertex := graph.getVertex(name)

	if srcVertex == nil {
		srcVertex = &vertex{
			name: name,
		}
		graph.vertices = append(graph.vertices, srcVertex)
	}

	return srcVertex
}

func (graph *graph) addEdge(srcKey, destKey string) {
	srcVertex := graph.addVertex(srcKey)

	if len(destKey) == 0 {
		return
	}

	destVertex := graph.addVertex(destKey)

	srcVertex.addNeighbor(destVertex)
}

func (graph *graph) visit(vertex *vertex, result *[]string) {
	vertex.visited = true

	for _, neighbor := range vertex.neighbors {
		if !neighbor.visited {
			graph.visit(neighbor, result)
		}
	}

	*result = append(*result, vertex.name)
}

func (graph *graph) topologicalSort() []string {
	result := make([]string, 0)

	for _, vertex := range graph.vertices {
		if !vertex.visited {
			graph.visit(vertex, &result)
		}
	}

	return result
}
