package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Og Suite")
}
