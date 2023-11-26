package planner

import (
	"errors"
	"trip-planner/planner/topsort"
)

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
	graph := topsort.Graph{}
	for _, item := range dependencies {
		graph.AddEdge(string(item.Destination), string(item.Dependency))
	}

	sortedDependencies, err := graph.TopologicalSort()

	var cyclicGraphError topsort.CyclicGraphError
	if errors.As(err, &cyclicGraphError) {
		return nil, CyclicDependencyError("")
	}

	return sortedDependencies, nil
}
