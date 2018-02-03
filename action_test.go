package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Action", func() {
	var a *Action

	BeforeEach(func() {
		a = &Action{Response{}}
	})

	It("should fail to Distribute() due to unknown action", func() {
		a.Type = "test"
		_, err := a.Distribute()
		Expect(err).To(HaveOccurred())
	})

	It("should fail to Distribute() to clipboard due to non string body", func() {
		a.Type = "clipboard"
		a.Action = "copy"
		_, err := a.Distribute()
		Expect(err).To(HaveOccurred())
	})

	It("should Distribute() to clipboard successfully", func() {
		a.Type = "clipboard"
		a.Action = "copy"
		a.Body = "test"
		res, err := a.Distribute()
		Expect(err).NotTo(HaveOccurred())
		Expect(res).NotTo(BeNil())
	})

	// It("should Distribute() to media successfully", func() {
	// 	a.Type = "media"
	// 	a.Action = "stop"
	// 	res, err := a.Distribute()
	// 	Expect(err).NotTo(HaveOccurred())
	// 	Expect(res).NotTo(BeNil())
	// })

	It("should fail to Distribute() to notification due to invalid body", func() {
		a.Type = "notification"
		a.Action = "send"
		_, err := a.Distribute()
		Expect(err).To(HaveOccurred())
	})

	It("should Distribute() to notification successfully", func() {
		a.Type = "notification"
		a.Action = "send"
		a.Body = "Test Case"
		a.Message = "If you read this, it is freaking working!"
		res, err := a.Distribute()
		Expect(err).NotTo(HaveOccurred())
		Expect(res).NotTo(BeNil())
	})

	It("should fail to Distribute() to volume due to invalid body", func() {
		a.Type = "volume"
		a.Action = "change"
		a.Body = true
		_, err := a.Distribute()
		Expect(err).To(HaveOccurred())
	})

	It("should Distribute() to volume successfully", func() {
		a.Type = "volume"
		a.Action = "check-mute"
		res, err := a.Distribute()
		Expect(err).NotTo(HaveOccurred())
		Expect(res).NotTo(BeNil())
	})
})
