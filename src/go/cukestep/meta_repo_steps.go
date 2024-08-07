package cukestep

import (
	"github.com/cucumber/godog"
	support "github.com/kkrull/marmot/cukesupport"
)

// Add step definitions to manage the life cycle of a meta repo.
func AddMetaRepoSteps(ctx *godog.ScenarioContext) {
	initWithArbitraryVersion := func() error { return support.InitNewMetaRepo("42") }
	ctx.Given(`^I have initialized a new meta repo$`, initWithArbitraryVersion)
}
