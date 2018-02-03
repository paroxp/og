package main

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {
	It("should compose error response successfully", func() {
		e := NewErrorResponse(fmt.Errorf("test"))
		Expect(e).NotTo(BeNil())
		Expect(e.Message).To(Equal("test"))
	})
})
