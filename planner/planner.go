package planner

import (
	"errors"
	"trip-planner/planner/topsort"
)

type CyclicDependencyError string

func (e CyclicDependencyError) Error() string {
	return "There is a cycle in the destination dependencies"
}

type EmptyDestinationError string

func (e EmptyDestinationError) Error() string {
	return "Empty destination"
}

var (
	cyclicGraphError topsort.CyclicGraphError
	emptyVertexError topsort.EmptyVertexKeyError
)

type destination string

type Dependency struct {
	Destination destination
	Dependency  destination
}

func Plan(dependencies []Dependency) ([]string, error) {
	graph := topsort.Graph{}

	for _, item := range dependencies {
		err := graph.AddEdge(string(item.Destination), string(item.Dependency))

		if errors.As(err, &emptyVertexError) {
			return nil, EmptyDestinationError("")
		}
	}

	sortedDependencies, err := graph.TopologicalSort()

	if errors.As(err, &cyclicGraphError) {
		return nil, CyclicDependencyError("")
	}

	return sortedDependencies, nil
}
