package cukefeature_test

import (
	"testing"

	"github.com/cucumber/godog"
	step "github.com/kkrull/marmot/cukestep"
	support "github.com/kkrull/marmot/cukesupport"
	. "github.com/onsi/gomega"
)

func TestFeatures(t *testing.T) {
	RegisterTestingT(t)
	suite := godog.TestSuite{
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"."},
			TestingT: t,
		},
		ScenarioInitializer: InitializeScenario,
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned from feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	step.AddLocalRepositorySteps(ctx)
	step.AddMetaRepoSteps(ctx)
	step.AddRepositorySteps(ctx)
	support.AddTo(ctx)
}
