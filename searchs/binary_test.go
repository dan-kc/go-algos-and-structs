package searchs_test

import (
	"example.com/packages/searchs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Binary search", func() {

	var arr []int
	var target int

	Context("Given the target is in the middle of the array", func() {
		It("should return the correct index", func() {
			arr = []int{3, 4, 5, 20, 21}
			target = 4
			Expect(searchs.BinarySearch(arr, target)).To(Equal(1))
		})
	})

	Context("Given the target is at the beginning of the array", func() {
		It("should return the correct index", func() {
			arr = []int{3, 4, 5, 20, 21}
			target = 3
			Expect(searchs.BinarySearch(arr, target)).To(Equal(0))
		})
	})

	Context("Given the target is at the end of the array", func() {
		It("should return the correct index", func() {
			arr = []int{3, 4, 5, 20, 21}
			target = 21
			Expect(searchs.BinarySearch(arr, target)).To(Equal(4))
		})
	})

	Context("Given the target does not exist in the array", func() {
		It("should return -1", func() {
			arr = []int{3, 4, 5, 20, 21}
			target = 10
			Expect(searchs.BinarySearch(arr, target)).To(Equal(-1))
		})
	})

	Context("Given an empty array", func() {
		It("should return -1", func() {
			arr = []int{}
			target = 10
			Expect(searchs.BinarySearch(arr, target)).To(Equal(-1))
		})
	})

	Context("Given the target is the only element in the array", func() {
		It("should return -1", func() {
			arr = []int{2}
			target = 2
			Expect(searchs.BinarySearch(arr, target)).To(Equal(0))
		})
	})

})
