package acceptance

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"

	"testing"
	"../common"
	"net/url"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

const PAGE_COMPOSITION_PORT = ":5001"

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acceptance Suite")
}

var agoutiDriver *agouti.WebDriver
var client *http.Client

var _ = BeforeSuite(func() {

	agoutiDriver = agouti.PhantomJS()
	client = &http.Client{}

	go common.StartPageCompositionServer(PAGE_COMPOSITION_PORT)
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
		Expect(page.Navigate(GetPageComposerUrl("pageComposerTest.html",""))).To(Succeed())
		Expect(page.All("#faulty").Text()).To(BeEquivalentTo("Faulty service"))
	})

	// I added this test to prove out the ability for goquery
	// to repeat the same activity as above - confirm #faulty value
	// Goquery used to customize the get request with headers
	// WebDriver does not support setting headers (#epicfail)
	It("should deliver a page with a #faulty tag with 'Faulty service' text using goquery", func() {
		response := Get(GetPageComposerUrl("pageComposerTest.html",""), map[string]string{"accept":"*/*"})
		doc, err := goquery.NewDocumentFromResponse(response)
		Expect(err).NotTo(HaveOccurred())
		Expect(doc.Find("#faulty").Text()).To(BeEquivalentTo("Faulty service"))
	})

	It("should silently drop favicon requests", func() {
		response := Get(GetPageComposerUrl("favicon.ico",""), map[string]string{"accept":"image/x-icon"})
		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusOK))
	})

	It("should ignore requests for anything other than html", func() {
		response := Get(GetPageComposerUrl("",""), map[string]string{"accept":"text/plain"})
		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusUnsupportedMediaType))
	})

	It("should process requests for any content type (thanks ie8)", func() {
		response := Get(GetPageComposerUrl("",""), map[string]string{"accept":"*/*"})
		doc, err := goquery.NewDocumentFromResponse(response)
		Expect(err).NotTo(HaveOccurred())
		Expect(doc.Find("#declarative").Text()).To(BeEquivalentTo("Replaced"))
	})

	/*

	it('should not die if given a poisoned url', function(done) {
	var targetUrl = getPageComposerUrl() + '?cid=271014_Primary-103466_email_et_27102014_%%%3dRedirectTo(%40RESOURCEURL1)%3d%%&mid=_&rid=%%External_ID%%&utm_source=ET&utm_medium=email&utm_term=27102014&utm_content=_&utm_campaign=271014_Primary_103466_%%%3dRedirectTo(%40RESOURCEURL1)%3d%%';
	request.get(targetUrl, {headers: {'accept': 'text/html'}}, function(err, response) {
	expect(response.statusCode).to.be(200);
	done();
	});
	});

	it('should return a 404 if any of the fragments return a 404', function(done) {
	var requestUrl = getPageComposerUrl('404backend');
	request.get(requestUrl,{headers: {'accept': 'text/html'}}, function(err, response) {
	expect(response.statusCode).to.be(404);
	done();
	});
	});

	it('should not return a 404 if any of the fragments have ignore-404 or ignore-error', function(done) {
	var requestUrl = getPageComposerUrl('ignore404backend');
	request.get(requestUrl,{headers: {'accept': 'text/html'}}, function(err, response) {
	expect(response.statusCode).to.be(200);
	done();
	});
	});
	 */
})


func GetPageComposerUrl(path string, search string) string {
	var url url.URL
	url.Scheme = "http"
	url.Path = path
	url.Host = "localhost" + PAGE_COMPOSITION_PORT
	url.RawQuery = search
	return url.String()
}

func Get(url string, headers map[string]string) *http.Response {
	request, err := http.NewRequest("GET", url, nil)
	Expect(err).NotTo(HaveOccurred())
	for key,value := range headers {
		request.Header.Add(key,value)
	}
	response, err := client.Do(request)
	Expect(err).NotTo(HaveOccurred())
	return response
}