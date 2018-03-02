package utils

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {

	Context("ReverseSlice", func() {
		It("shall return a reversed slice of strings", func() {
			slice := []string{"one", "two", "three", "four"}
			expectedSlice := []string{"four", "three", "two", "one"}
			ReverseSlice(slice)
			Expect(slice).To(Equal(expectedSlice))
		})
	})
})
