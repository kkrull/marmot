package main_test

import (
	"testing"

	"github.com/cucumber/godog"
	steps "github.com/kkrull/marmot/feature-steps"
	support "github.com/kkrull/marmot/feature-support"
	. "github.com/onsi/gomega"
)

func TestFeatures(t *testing.T) {
	RegisterTestingT(t)
	suite := godog.TestSuite{
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
		ScenarioInitializer: InitializeScenario,
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned from feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	steps.AddMetaRepoSteps(ctx)
	steps.AddRepositorySteps(ctx)
	support.AddTo(ctx)
}
