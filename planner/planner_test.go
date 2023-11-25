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
	sortedDestinations := planner.Plan(
		[]planner.Dependency{
			{"Paris", ""},
		})
	s.Equal([]string{"Paris"}, sortedDestinations)
}

func TestPlannerTestSuite(t *testing.T) {
	suite.Run(t, new(PlannerTestSuite))
}
