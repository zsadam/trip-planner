package planner

type destination string

type Dependency struct {
	Destination destination
	Dependency  destination
}

func Plan(dependencies []Dependency) []string {
	result := make([]string, len(dependencies))

	for i, dependency := range dependencies {
		result[i] = string(dependency.Destination)
	}

	return result
}
