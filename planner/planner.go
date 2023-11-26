package planner

import "errors"

type CyclicDependencyError string

func (e CyclicDependencyError) Error() string {
	return "There is a cycle in the destination dependencies"
}

type destination string

type Dependency struct {
	Destination destination
	Dependency  destination
}

func Plan(dependencies []Dependency) ([]string, error) {
	graph := graph{}
	for _, item := range dependencies {
		graph.addEdge(string(item.Destination), string(item.Dependency))
	}

	sortedDependencies, err := graph.topologicalSort()

	var cyclicGraphError CyclicGraphError
	if errors.As(err, &cyclicGraphError) {
		return nil, CyclicDependencyError("")
	}

	return sortedDependencies, nil
}
