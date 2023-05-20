package structs_test

import (
	"example.com/packages/structs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Linked list", func() {

	ll := structs.NewLinkedList()

	Context("Given the target is the only element in the array", func() {
		It("should return -1", func() {
			arr = []int{2}
			target = 2
			Expect(searchs.BinarySearch(arr, target)).To(Equal(0))
		})
	})

})
