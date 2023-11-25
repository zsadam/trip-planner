package planner

type destination string

type Dependency struct {
	Destination destination
	Dependency  destination
}

func Plan(dependencies []Dependency) []string {
	return []string{string(dependencies[0].Destination)}
}
