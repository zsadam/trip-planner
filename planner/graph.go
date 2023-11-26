package planner

type CyclicGraphError string

func (e CyclicGraphError) Error() string {
	return "Please provide an acyclic graph"
}

type vertex struct {
	name        string
	neighbors   []*vertex
	visited     bool
	cyclicCheck bool
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

func (graph *graph) visit(vertex *vertex, result *[]string) error {
	vertex.visited = true
	vertex.cyclicCheck = true

	for _, neighbor := range vertex.neighbors {
		if !neighbor.visited {
			err := graph.visit(neighbor, result)
			if err != nil {
				return err
			}
		} else if neighbor.cyclicCheck {
			return CyclicGraphError("")
		}
	}

	vertex.cyclicCheck = false

	*result = append(*result, vertex.name)

	return nil
}

func (graph *graph) topologicalSort() ([]string, error) {
	result := make([]string, 0)

	for _, vertex := range graph.vertices {
		if !vertex.visited {
			err := graph.visit(vertex, &result)
			if err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}
