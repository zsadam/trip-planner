package planner

type destination string

type Dependency struct {
	Destination destination
	Dependency  destination
}

func Plan(dependencies []Dependency) []string {
	dependencyMap := make(map[string]string, len(dependencies))

	for _, dependency := range dependencies {
		if len(dependency.Dependency) > 0 {
			dependencyMap[string(dependency.Dependency)] = string(dependency.Dependency)
			dependencyMap[string(dependency.Destination)] = string(dependency.Destination)

			continue
		}

		dependencyMap[string(dependency.Destination)] = string(dependency.Destination)
	}

	result := make([]string, 0, len(dependencies))
	for _, value := range dependencyMap {
		result = append(result, value)
	}

	return result
}
