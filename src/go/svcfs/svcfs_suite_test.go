package svcfs_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSvcFs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "svcfs")
}
