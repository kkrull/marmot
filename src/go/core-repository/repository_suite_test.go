package core_repository_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCoreRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "core_repository")
}
