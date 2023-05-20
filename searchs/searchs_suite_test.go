package searchs_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSearchs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Searchs Suite")
}
