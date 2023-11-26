package planner_test

import (
	"testing"
	"trip-planner/planner"

	"github.com/stretchr/testify/suite"
)

type PlannerTestSuite struct {
	suite.Suite
}

func (s *PlannerTestSuite) TestPlanShouldSortGivenDestinations() {
	testCases := []struct {
		description  string
		dependencies []planner.Dependency
		expected     []string
	}{
		{
			"single destination without dependency",
			[]planner.Dependency{
				{"Paris", ""},
			},
			[]string{"Paris"},
		},
		{
			"multiple destinations without dependency",
			[]planner.Dependency{
				{"Paris", ""},
				{"London", ""},
				{"Madrid", ""},
			},
			[]string{"Paris", "London", "Madrid"},
		},
		{
			"multiple destinations with single dependency",
			[]planner.Dependency{
				{"Paris", ""},
				{"London", "Madrid"},
				{"Madrid", ""},
			},
			[]string{"Paris", "Madrid", "London"},
		},
		{
			"multiple destinations with multiple dependency",
			[]planner.Dependency{
				{"Paris", ""},
				{"Berlin", "Madrid"},
				{"Madrid", "London"},
				{"Amsterdam", "Paris"},
				{"Roma", "Berlin"},
				{"London", ""},
			},
			[]string{"Paris", "London", "Madrid", "Berlin", "Amsterdam", "Roma"},
		},
	}

	for _, tc := range testCases {
		s.Run(
			tc.description,
			func() {
				sortedDestinations, err := planner.Plan(tc.dependencies)
				s.Nil(err)
				s.Equal(tc.expected, sortedDestinations)
			},
		)
	}
}

func (s *PlannerTestSuite) TestPlanShouldReturnErrorWhenDestinationsAreCyclic() {
	dependencies := []planner.Dependency{
		{"Paris", "Berlin"},
		{"Berlin", "Paris"},
	}
	sortedDestinations, err := planner.Plan(dependencies)
	s.Nil(sortedDestinations)
	s.ErrorIs(err, planner.CyclicDependencyError(""))
}

func TestPlannerTestSuite(t *testing.T) {
	suite.Run(t, new(PlannerTestSuite))
}
