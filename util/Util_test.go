package util

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)
func TestUtil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestUtilSuite")
}

var _ = Describe("Utility Library", func() {

	It("should generate the url object", func() {
  		Expect(TimeToMillis("1000")).To(BeEquivalentTo(1000))
		Expect(TimeToMillis("1s")).To(BeEquivalentTo(1000))
		Expect(TimeToMillis("1m")).To(BeEquivalentTo(1000*60))
		Expect(TimeToMillis("1ms")).To(BeEquivalentTo(1))
		Expect(TimeToMillis("1h")).To(BeEquivalentTo(1000*60*60))
		Expect(TimeToMillis("5d")).To(BeEquivalentTo(1000*60*60*24*5))
	})
})