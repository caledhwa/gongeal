package requestinterrogator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestRequestInterrogator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestInterrogatorSuite")
}

var _ = Describe("Request Interrogator", func() {

	It("should generate the url object", func() {

	})

	It("should extract parameters from the query", func() {

	})

	It("should only extract parameters from the query when they're not empty", func() {

	})

	It("should generate the url object", func() {

	})

	It("should extract parameters from the path", func() {

	})

	It("should extract parametrs from the path if multiple paths match", func() {

	})

	It("should extract parameters if multiple overlap it takes the last one", func() {

	})

	It("should extract query parameters", func() {

	})

	It("should extract headers", func() {

	})

	It("should default user:userId if not logged in", func() {

	})

	It("should get user from request", func() {

	})

	It("should parse cdn url configuration using template variables", func() {

	})
})
