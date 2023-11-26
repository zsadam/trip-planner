package planner_test

import (
	"testing"
	"trip-planner/planner"

	"github.com/stretchr/testify/suite"
)

type PlannerTestSuite struct {
	suite.Suite
}

func (s *PlannerTestSuite) TestPlan() {
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
	}

	for _, tc := range testCases {
		s.Run(
			tc.description,
			func() {
				sortedDestinations := planner.Plan(tc.dependencies)
				s.Equal(tc.expected, sortedDestinations)
			},
		)
	}
}

func TestPlannerTestSuite(t *testing.T) {
	suite.Run(t, new(PlannerTestSuite))
}
