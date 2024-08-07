package cukesupport

import "github.com/cucumber/godog"

// Make hooks available to a scenario and add fixture methods to avoid shared state among scenarios.
func AddFixtures(ctx *godog.ScenarioContext) {
	addLocalDirHook(ctx)
	addMetaRepoFixture(ctx)
}
