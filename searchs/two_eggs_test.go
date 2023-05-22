package searchs_test

import (
	"example.com/packages/searchs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Two eggs search", func() {

	var breakArr []bool

	Context("If the egg never breaks", func() {
		It("should return -1", func() {
			breakArr = []bool{false, false, false, false, false}
			Expect(searchs.TwoEggsSearch(breakArr)).To(Equal(-1))
		})
	})

	Context("If the egg breaks on the ground floor", func() {
		It("should return 0", func() {
			breakArr = []bool{true, true, true, true, true}
			Expect(searchs.TwoEggsSearch(breakArr)).To(Equal(0))
		})
	})

	Context("If the egg breaks on the third floor", func() {
		It("should return 3", func() {
			breakArr = []bool{false, false, false, true, true}
			Expect(searchs.TwoEggsSearch(breakArr)).To(Equal(3))
		})
	})

	Context("If the egg breaks on the last floor", func() {
		It("should return 4", func() {
			breakArr = []bool{false, false, false, false, true}
			Expect(searchs.TwoEggsSearch(breakArr)).To(Equal(4))
		})
	})

	Context("If the eggs break array isn't sorted", func() {
		It("should return -1", func() {
			breakArr = []bool{false, false, true, false, true}
			Expect(searchs.TwoEggsSearch(breakArr)).To(Equal(-1))
		})
	})

})
