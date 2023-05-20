package sorts_test

import (
	"example.com/packages/sorts"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bubble sort", func() {

	var test, sortedTest *[]int

	Context("Given a sorted array", func() {
		It("should return a sorted array", func() {
			test = &[]int{3, 5, 10, 22, 64}
			sortedTest = &[]int{3, 5, 10, 22, 64}
			sorts.BubbleSort(test)
			Expect(test).To(Equal(sortedTest))
		})
	})

	Context("Given a reverse sorted array", func() {
		It("should return an ordered array", func() {
			test = &[]int{64, 22, 10, 5, 3}
			sortedTest = &[]int{3, 5, 10, 22, 64}
			sorts.BubbleSort(test)
			Expect(test).To(Equal(sortedTest))
		})
	})

	Context("Given an empty array", func() {
		It("should return an empty array", func() {
			test = &[]int{}
			sortedTest = &[]int{}
			sorts.BubbleSort(test)
			Expect(test).To(Equal(sortedTest))
		})
	})

	Context("Given an array with duplicate elements", func() {
		It("should return an ordered array", func() {
			test = &[]int{64, 22, 22, 22, 10, 5, 3}
			sortedTest = &[]int{3, 5, 10, 22, 22, 22, 64}
			sorts.BubbleSort(test)
			Expect(test).To(Equal(sortedTest))
		})
	})

	Context("Given an array with negative elements", func() {
		It("should return an ordered array", func() {
			test = &[]int{-64, 22, 10, -5, 3}
			sortedTest = &[]int{-64, -5, 3, 10, 22}
			sorts.BubbleSort(test)
			Expect(test).To(Equal(sortedTest))
		})
	})

	Context("Given an array with single element", func() {
		It("should return a single element array", func() {
			test = &[]int{3}
			sortedTest = &[]int{3}
			sorts.BubbleSort(test)
			Expect(test).To(Equal(sortedTest))
		})
	})

})
