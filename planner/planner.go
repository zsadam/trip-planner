package planner

type destination string

type Dependency struct {
	Destination destination
	Dependency  destination
}

func Plan(dependencies []Dependency) []string {
	graph := graph{}
	for _, item := range dependencies {
		graph.addEdge(string(item.Destination), string(item.Dependency))
	}

	return graph.topologicalSort()
}
