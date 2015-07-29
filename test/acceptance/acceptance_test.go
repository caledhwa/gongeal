package acceptance

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"

	"testing"
	"../common"
)

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acceptance Suite")
}

var agoutiDriver *agouti.WebDriver

var _ = BeforeSuite(func() {

	agoutiDriver = agouti.PhantomJS()

	go common.StartPageCompositionServer(":5001")
	go common.StartStubServer(":5002")

	Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
	Expect(agoutiDriver.Stop()).To(Succeed())
})

var _ = Describe("Page Composer", func() {

	var page *agouti.Page

	BeforeEach(func() {
		var err error
		page, err = agoutiDriver.NewPage(agouti.Browser("chrome"))
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	// I added this test to simply prove out the agouti framework
	// on one of the static pages in the test/common folder
	It("should deliver a page with a #faulty tag with 'Faulty service' text", func() {
		By("navigating to it", func() {
			Expect(page.Navigate("http://localhost:5001/pageComposerTest.html")).To(Succeed())
			Expect(page.All("#faulty").Text()).To(BeEquivalentTo("Faulty service"))
		})
	})
})
