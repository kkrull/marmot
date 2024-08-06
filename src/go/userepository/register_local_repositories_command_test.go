package userepository_test

import (
	mock "github.com/kkrull/marmot/corerepositorymock"
	expect "github.com/kkrull/marmot/testsupportexpect"
	"github.com/kkrull/marmot/use"
	"github.com/kkrull/marmot/userepository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegisterLocalRepositoriesCommand", func() {
	var (
		factory use.CommandFactory
		source  *mock.RepositorySource
		subject *userepository.RegisterLocalRepositoriesCommand
	)

	Describe("#Run", func() {
		It("exists", func() {
			source = mock.NewRepositorySource()
			factory = use.NewCommandFactory().WithRepositorySource(source)
			subject = expect.NoError(factory.NewRegisterLocalRepositories())
			Expect(subject).NotTo(BeNil())
		})
	})
})
