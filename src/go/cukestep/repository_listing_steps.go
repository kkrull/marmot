package cukestep

import (
	"context"

	"github.com/cucumber/godog"
	core "github.com/kkrull/marmot/corerepository"
	. "github.com/onsi/gomega"
)

// State to clear between scenarios
var thatListing core.Repositories

// Add step definitions related to listing repositories.
func AddRepositorySteps(ctx *godog.ScenarioContext) {
	ctx.After(func(ctx context.Context, _ *godog.Scenario, err error) (context.Context, error) {
		thatListing = nil
		return ctx, err
	})

	ctx.Then(`^that repository listing should be empty$`, thatListingShouldBeEmpty)
}

func thatListingShouldBeEmpty() {
	Expect(thatListing.Count()).To(Equal(0))
}
