package main_test

import (
	"testing"

	"github.com/cucumber/godog"
	"github.com/kkrull/marmot/hooks"
	"github.com/kkrull/marmot/step_definitions"
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
	hooks.AddTo(ctx)
	step_definitions.AddMetaRepoSteps(ctx)
	step_definitions.AddRepositorySteps(ctx)
}
