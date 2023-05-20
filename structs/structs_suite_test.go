package structs_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStructs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Structs Suite")
}
