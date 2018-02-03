package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {
	Context("Helpers", func() {
		It("should return correct value ifZero() has nothing to work with", func() {
			i := ifZero(2, 10)
			Expect(i).To(Equal(2))
		})

		It("should return correct value ifZero() detects 0 value", func() {
			i := ifZero(0, 10)
			Expect(i).To(Equal(10))
		})

		It("should fail to convert actionToVolume() due to non numeric sting provided in body", func() {
			a := &Action{Response: Response{}}
			a.Action = "test"
			a.Body = "test"
			_, err := actionToVolume(a)

			Expect(err).To(HaveOccurred())
		})

		It("should convert actionToVolume() successfully", func() {
			a := &Action{Response: Response{}}
			a.Action = "test"
			v, err := actionToVolume(a)

			Expect(err).NotTo(HaveOccurred())
			Expect(v.Action).To(Equal("test"))
		})

		It("should convert actionToVolume() successfully when string value has been provided", func() {
			a := &Action{Response: Response{}}
			a.Action = "test"
			a.Body = "5"
			v, err := actionToVolume(a)

			Expect(err).NotTo(HaveOccurred())
			Expect(v.Value).To(Equal(5))
		})

		It("should convert actionToVolume() successfully when int value has been provided", func() {
			a := &Action{Response: Response{}}
			a.Action = "test"
			a.Body = 5
			v, err := actionToVolume(a)

			Expect(err).NotTo(HaveOccurred())
			Expect(v.Value).To(Equal(5))
		})
	})

	Context("Volume has been setup", func() {
		var vol int
		v := &Volume{Res: &Response{}}

		It("should fail to AdjustVolume() due to unsupported action", func() {
			v.Action = "test"
			_, err := v.AdjustVolume()
			Expect(err).To(HaveOccurred())
		})

		It("should AdjustVolume() when asked to mute", func() {
			v.Action = "mute"
			res, err := v.AdjustVolume()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
		})

		It("should AdjustVolume() when asked to unmute", func() {
			v.Action = "unmute"
			res, err := v.AdjustVolume()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
		})

		It("should AdjustVolume() when asked to increase", func() {
			v.Action = "increase"
			v.Value = 10
			res, err := v.AdjustVolume()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
		})

		It("should AdjustVolume() when asked to increase", func() {
			v.Action = "decrease"
			v.Value = 10
			res, err := v.AdjustVolume()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
		})

		It("should AdjustVolume() when asked to increase", func() {
			v.Action = "check-mute"
			res, err := v.AdjustVolume()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
		})

		It("should AdjustVolume() when asked to increase", func() {
			v.Action = "check-volume"
			res, err := v.AdjustVolume()
			vol = res.Body.(int)
			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
		})

		It("should AdjustVolume() when asked to increase", func() {
			v.Action = "change"
			v.Value = vol
			res, err := v.AdjustVolume()
			Expect(err).NotTo(HaveOccurred())
			Expect(res).NotTo(BeNil())
		})
	})
})
