package structs_test

import (
	"example.com/packages/structs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Linked list", func() {
	var ll structs.LinkedList
	Context("Given the list is empty", func() {
		It("should have an undefined head", func() {
			ll = structs.NewLinkedList()
			Expect(ll.Head).To(Equal(nil))
		})
	})
})
