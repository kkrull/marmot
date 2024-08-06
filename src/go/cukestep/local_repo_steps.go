package cukestep

import (
	"fmt"
	"path/filepath"

	"github.com/cucumber/godog"
	support "github.com/kkrull/marmot/cukesupport"
)

/* State */

var thoseLocalRepositories *support.LocalRepositories = support.NoLocalRepositories()

/* Configuration */

// Add step definitions related to repositories on the local filesystem.
func AddLocalRepositorySteps(ctx *godog.ScenarioContext) {
	ctx.Given(`^Git repositories on the local filesystem$`, func() error {
		if testDir, pathErr := support.TestDir(); pathErr != nil {
			return pathErr
		} else {
			return createLocalRepo(filepath.Join(testDir, "empty-dir"))
		}
	})

	ctx.Given(`^I have registered those local repositories with a meta repo$`, registerThoseLocals)

	ctx.Then(`^that repository listing should include those local repositories$`, func() error {
		if repoDir, pathErr := support.TestSubDir("empty-dir"); pathErr != nil {
			return pathErr
		} else {
			thatListingShouldHaveLocals(repoDir)
			return nil
		}
	})
}

/* Steps */

func createLocalRepo(repoDir string) error {
	if repo, repoErr := support.InitLocalRepository(repoDir); repoErr != nil {
		return repoErr
	} else {
		thoseLocalRepositories = support.SomeLocalRepositories(repo)
		return nil
	}
}

func registerThoseLocals() error {
	if factory, factoryErr := support.ThatCommandFactory(); factoryErr != nil {
		return fmt.Errorf("local_repo_steps: failed to configure; %w", factoryErr)
	} else if registerCmd, factoryErr := factory.NewRegisterLocalRepositories(); factoryErr != nil {
		return fmt.Errorf("local_repo_steps: failed to initialize; %w", factoryErr)
	} else {
		return registerCmd.Run(thoseLocalRepositories.LocalPaths()...)
	}
}
