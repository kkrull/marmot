package usemetarepo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUseMetaRepo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "usemetarepo")
}
