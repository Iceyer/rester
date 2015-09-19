package rester_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRester(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rester Suite")
}
