package planner

type Graph struct {
	adjList  map[string][]string
	vertices []string
}

func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[string][]string),
	}
}

func (g *Graph) addEdge(u, v string) {
	g.vertices = append(g.vertices, u)
	g.adjList[u] = append(g.adjList[u], v)
}

func visit(v string, visited map[string]bool, stack *[]string, g *Graph) {
	visited[v] = true

	for _, neighbor := range g.adjList[v] {
		if len(neighbor) > 0 && !visited[neighbor] {
			visit(neighbor, visited, stack, g)
		}
	}

	*stack = append(*stack, v)
}

func topologicalSort(g *Graph) []string {
	visited := make(map[string]bool)
	stack := make([]string, 0)

	for _, vertex := range g.vertices {
		if !visited[vertex] {
			visit(vertex, visited, &stack, g)
		}
	}

	return stack
}

type destination string

type Dependency struct {
	Destination destination
	Dependency  destination
}

func Plan(dependencies []Dependency) []string {
	graph := NewGraph()

	for _, item := range dependencies {
		graph.addEdge(string(item.Destination), string(item.Dependency))
	}

	return topologicalSort(graph)
}
