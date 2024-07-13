package core_metarepo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCoreMetaRepo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "core_metarepo")
}
