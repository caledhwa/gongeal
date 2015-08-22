//
//package acceptance
//
//
//import (
//	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
//	"testing"
//	"../common"
//	"net/url"
//	"net/http"
//	"github.com/PuerkitoBio/goquery"
//	"time"
//	"io/ioutil"
//	"strconv"
//)
//
//const PAGE_COMPOSITION_PORT int = 5001
//var client *http.Client
//
//func TestAcceptance(t *testing.T) {
//	RegisterFailHandler(Fail)
//	RunSpecs(t, "Acceptance Suite")
//}
//
//var _ = BeforeSuite(func() {
//
//	client = &http.Client{}
//
//	go common.StartPageCompositionServer(PAGE_COMPOSITION_PORT,"localhost",func(){},"")
//	go common.StartStubServer(5002)
//})
//
//var _ = Describe("Page Composer", func() {
//
//	// I added this test to prove out the ability for goquery
//	// to repeat the same activity as above - confirm #faulty value
//	// Goquery used to customize the get request with headers
//	// WebDriver does not support setting headers (#epicfail)
//	It("should deliver a page with a #faulty tag with 'Faulty service' text using goquery", func() {
//		response := Get(GetPageComposerUrl("pageComposerTest.html",""), map[string]string{"accept":"*/*"})
//		doc, err := goquery.NewDocumentFromResponse(response)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(doc.Find("#faulty").Text()).To(BeEquivalentTo("Faulty service"))
//	})
//
//	It("should silently drop favicon requests", func() {
//		response := Get(GetPageComposerUrl("favicon.ico",""), map[string]string{"accept":"image/x-icon"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusOK))
//	})
//
//	It("should ignore requests for anything other than html", func() {
//		response := Get(GetPageComposerUrl("",""), map[string]string{"accept":"text/plain"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusUnsupportedMediaType))
//	})
//
//	It("should process requests for any content type (thanks ie8)", func() {
//		response := Get(GetPageComposerUrl("",""), map[string]string{"accept":"*/*"})
//		doc, err := goquery.NewDocumentFromResponse(response)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(doc.Find("#declarative").Text()).To(BeEquivalentTo("Replaced"))
//	})
//
//	It("should not die if given a poisoned url", func() {
//		response := Get(GetPageComposerUrl("","cid=271014_Primary-103466_email_et_27102014_%%%3dRedirectTo(%40RESOURCEURL1)%3d%%&mid=_&rid=%%External_ID%%&utm_source=ET&utm_medium=email&utm_term=27102014&utm_content=_&utm_campaign=271014_Primary_103466_%%%3dRedirectTo(%40RESOURCEURL1)%3d%%"), map[string]string{"accept":"text/html"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusOK))
//	})
//
//	It("should return a 404 if any of the fragments return a 404", func() {
//		response := Get(GetPageComposerUrl("404backend",""), map[string]string{"accept":"text/html"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusNotFound))
//	})
//
//	It("should not return a 404 if any of the fragments have ignore-404 or ignore-error", func() {
//		response := Get(GetPageComposerUrl("ignore404backend",""), map[string]string{"accept":"text/html"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusOK))
//	})
//
//	It("should return a 404 if the backend template returns a 404", func() {
//		response := Get(GetPageComposerUrl("404",""), map[string]string{"accept":"text/html"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusNotFound))
//	})
//
//	It("should return a 500 if the backend template returns a 500", func() {
//		response := Get(GetPageComposerUrl("500",""), map[string]string{"accept":"text/html"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusServiceUnavailable))
//	})
//
//	It("should return a 500 if the backend template returns a 500", func() {
//		response := Get(GetPageComposerUrl("broken",""), map[string]string{"accept":"text/html"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusServiceUnavailable))
//	})
//
//	It("should add no-store cache-control header if any fragments use cx-no-cache", func() {
//		response := Get(GetPageComposerUrl("noCacheBackend",""), map[string]string{"accept":"text/html"})
//		Expect(response.Header.Get("cache-control")).To(BeEquivalentTo("no-store"))
//	})
//
//	It("should fail quietly if the backend is configured to do so", func() {
//		response := Get(GetPageComposerUrl("quiet",""), map[string]string{"accept":"text/html"})
//		doc, err := goquery.NewDocumentFromResponse(response)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(doc.Find("#faulty").Text()).To(BeEquivalentTo("Faulty service"))
//	})
//
//	It("should fail loudly if the backend is configured to do so", func() {
//		response := Get(GetPageComposerUrl("",""), map[string]string{"accept":"text/html"})
//		doc, err := goquery.NewDocumentFromResponse(response)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(doc.Find("#faulty").Text()).To(BeEquivalentTo("Error: Service http://localhost:5001/500 FAILED due to status code 500"))
//	})
//
//	It("should leave the content that was originally in the element if it is configured to do so", func() {
//		response := Get(GetPageComposerUrl("leave",""), map[string]string{"accept":"text/html"})
//		doc, err := goquery.NewDocumentFromResponse(response)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(doc.Find("#faulty").Text()).To(BeEquivalentTo("Faulty service"))
//	})
//
//	It("should leave the HTML content that was originally in the element if it is configured to do so", func() {
//		response := Get(GetPageComposerUrl("leave",""), map[string]string{"accept":"text/html"})
//		doc, err := goquery.NewDocumentFromResponse(response)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(doc.Find("#faultyhtml h1").Text()).To(BeEquivalentTo("Bob"))
//		Expect(doc.Find("#faultyhtml span").Text()).To(BeEquivalentTo("The builder"))
//	})
//
//
//	It("should fail gracefully if the service returns no response at all", func() {
//		GetSection("","","#broken",func(text string){
//			Expect(text).To(BeEquivalentTo("Error: Service http://localhost:5001/broken FAILED due to socket hang up"))
//		})
//	})
//
//	It("should remove the element if cx-replace-outer is set", func() {
//		response := Get(GetPageComposerUrl("",""), map[string]string{"accept":"text/html"})
//		doc, err := goquery.NewDocumentFromResponse(response)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(doc.Find("#replace-outer-content").Length()).To(BeEquivalentTo(0))
//		Expect(doc.Find("#replace-outer").Text()).To(BeEquivalentTo("wrapping Replaced content"))
//	})
////
//	It("should ignore a cx-url that is invalid", func() {
//		GetSection("","","#invalidurl",func(text string){
//			Expect(text).To(BeEquivalentTo("Error: Service invalid FAILED due to Invalid URL invalid"))
//		})
//	})
//
//
//	It("should ignore a cx-url that is invalid unless it is cache", func() {
//		GetSection("","","#cacheurl1",func(text string){
//			Expect(text).To(BeEquivalentTo("No content in cache at key: cache"))
//		})
//	})
//
//	It("should ignore a cx-url that is invalid unless it is cache, and get the content if cache is primed", func() {
//		GetSection("","","#cacheurl2",func(text string){
//			Expect(text).To(BeEquivalentTo("Replaced"))
//		})
//	})
//
//	It("should allow simple mustache logic", func() {
//		GetSection("","logic=yes","#testlogic",func(text string) {
//			Expect(text).To(BeEquivalentTo("Logic ftw!"))
//		})
//	})
//
//	It("should have access to current environment", func() {
//		GetSection("","","#environment",func(text string) {
//			Expect(text).To(BeEquivalentTo("test"))
//		})
//	})
//
//	It("should not cache segments that return no-store in Cache-control header", func() {
//		GetSection("","","#no-store",func(text string){
//			before := text
//			time.AfterFunc(1*time.Millisecond, func() {
//				GetSection("","","#no-store",func(text string) {
//					Expect(text).NotTo(BeEquivalentTo(before))
//				})
//			})
//		})
//	})
//
//	It("should pass no-store in Cache-control header from fragment response to client response", func() {
//		response := Get(GetPageComposerUrl("",""), map[string]string{})
//		Expect(response.Header.Get("cache-control")).To(BeEquivalentTo("no-store"))
//	})
//
//	It("should honor max-age when sent through in fragments", func() {
//		time.AfterFunc(1 * time.Second, func() {
//			GetSection("", "", "#max-age", func(text string) {
//				time.AfterFunc(50 * time.Millisecond, func() {
//					GetSection("", "", "#max-age", func(text2 string) {
//						Expect(text2).To(BeEquivalentTo(text))
//						time.AfterFunc(1 * time.Second, func() {
//							GetSection("", "", "#max-age", func(text3 string) {
//								Expect(text3).NotTo(BeEquivalentTo(text))
//							})
//						})
//					})
//				})
//			})
//		}) // Allow previous test cache to clear
//	})
//
//	It("should pass through non GET requests directly to the backend service along with headers and cookies", func() {
//		request, err := http.NewRequest("POST", GetPageComposerUrl("post",""), nil)
//		Expect(err).NotTo(HaveOccurred())
//		request.AddCookie(&http.Cookie{
//			Name: "PostCookie",
//			Value: "Hello",
//		})
//		request.Header.Add("accept","text/html")
//		response, err := client.Do(request)
//		Expect(err).NotTo(HaveOccurred())
//		content, err := ioutil.ReadAll(response.Body)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(string(content)).To(BeEquivalentTo("POST Hello"))
//	})
//
//	It("should pass a cookie to a url when using a template", func() {
//		request, err := http.NewRequest("GET", GetPageComposerUrl("",""), nil)
//		Expect(err).NotTo(HaveOccurred())
//		request.AddCookie(&http.Cookie{
//			Name: "geo",
//			Value: "US",
//		})
//		request.Header.Add("accept","text/html")
//		response, err := client.Do(request)
//		Expect(err).NotTo(HaveOccurred())
//		doc, err := goquery.NewDocumentFromResponse(response)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(doc.Find("#country").Text()).To(BeEquivalentTo("US"))
//	})
//
//	It("should NOT pass through GET requests that have text/html content type to a backend service", func() {
//		response := Get(GetPageComposerUrl("post",""), map[string]string{"accept":"text/html"})
//		content, err := ioutil.ReadAll(response.Body)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(string(content)).To(BeEquivalentTo("GET /post"))
//	})
//
//	It("should select the correct backend if a selectorFn is invoked", func() {
//		response := Get(GetPageComposerUrl("","selectFn=true"), map[string]string{"accept":"text/html"})
//		doc, err := goquery.NewDocumentFromResponse(response)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(doc.Find("#select").Text()).To(BeEquivalentTo("This is the backend selected by a selector fn"))
//	})
//
//	It("should use the handler functions to respond to a 403 status code", func() {
//		response := Get(GetPageComposerUrl("403backend",""), map[string]string{"accept":"text/html"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusForbidden))
//	})
//
//	It("should use the handler functions to respond to a 403 status code of the backend template", func() {
//		response := Get(GetPageComposerUrl("403",""), map[string]string{"accept":"text/html"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusForbidden))
//	})
//
//	It("should use the handler functions to respond to a 302 status code in a fragment", func() {
//		response := Get(GetPageComposerUrl("302backend",""), map[string]string{"accept":"text/html"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusFound))
//		Expect(response.Header.Get("location")).To(BeEquivalentTo("/replaced"))
//	})
//
//	It("should use the handler functions to respond to a 302 status code in a backend template", func() {
//		response := Get(GetPageComposerUrl("302",""), map[string]string{"accept":"text/html"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusFound))
//		Expect(response.Header.Get("location")).To(BeEquivalentTo("/replaced"))
//	})
//
//	It("should use the handler functions to respond to a 302 status code in a backend template", func() {
//		response := Get(GetPageComposerUrl("302",""), map[string]string{"accept":"text/html"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusFound))
//		Expect(response.Header.Get("location")).To(BeEquivalentTo("/replaced"))
//	})
//
//	It("should pass x-tracer to downstreams", func() {
//		response := Get(GetPageComposerUrl("tracer",""), map[string]string{"accept":"text/html","x-tracer":"willie wonka"})
//		content, err := ioutil.ReadAll(response.Body)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(string(content)).To(BeEquivalentTo("willie wonka"))
//	})
//
//	It("should pass accept-language to downstreams", func() {
//		response := Get(GetPageComposerUrl("lang",""), map[string]string{"accept":"text/html","Accept-Language":"es"})
//		content, err := ioutil.ReadAll(response.Body)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(string(content)).To(BeEquivalentTo("es"))
//	})
//
//	It("should forward specified headers to downstreams", func() {
//		response := Get(GetPageComposerUrl("header/x-geoip-country-code",""), map[string]string{"accept":"text/html","x-geoip-country-code":"GB"})
//		content, err := ioutil.ReadAll(response.Body)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(string(content)).To(BeEquivalentTo("GB"))
//	})
//
//	It("should pass a default accept header of text/html", func() {
//		response := Get(GetPageComposerUrl("header/accept",""), map[string]string{"accept":"text/html"})
//		content, err := ioutil.ReadAll(response.Body)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(string(content)).To(BeEquivalentTo("text/html"))
//	})
//
//	It("should allow fragments to over ride the accept header", func() {
//		response := Get(GetPageComposerUrl("",""), map[string]string{"accept":"text/html"})
//		doc, err := goquery.NewDocumentFromResponse(response)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(doc.Find("#accept").Text()).To(BeEquivalentTo("application/json"))
//	})
//
//	It("should pass set-cookie headers upstream from a backend", func() {
//		response := Get(GetPageComposerUrl("set-cookie",""), map[string]string{"accept":"text/html"})
//		Expect(response.Header.Get("set-cookie")).To(ContainSubstring("hello=world; Path=/"))
//		Expect(response.Header.Get("set-cookie")).To(ContainSubstring("another=cookie; Path=/"))
//		Expect(response.Header.Get("set-cookie")).To(ContainSubstring("hello=again; Path=/"))
//	})
//
//	It("should retrieve bundles via the cx-bundle directive and cdn configuration using service supplied version numbers if appropriate", func() {
//		response := Get(GetPageComposerUrl("bundles",""), map[string]string{"accept":"text/html"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusOK))
//		doc, err := goquery.NewDocumentFromResponse(response)
//		Expect(err).NotTo(HaveOccurred())
//		bundles := doc.Find(".bundle")
//		Expect(bundles.First().Text()).To(BeEquivalentTo("service-one >> 100 >> top.js.htmlservice-two >> YOU_SPECIFIED_A_BUNDLE_THAT_ISNT_AVAILABLE_TO_THIS_PAGE >> top.js.html"))
//	})
//
//	It("should use allow you to specify a host over-ride to use instead of the target host", func() {
//		response := Get(GetPageComposerUrl("differenthost",""), map[string]string{"accept":"text/html"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusOK))
//		content, err := ioutil.ReadAll(response.Body)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(string(content)).To(BeEquivalentTo("tes.co.uk"))
//	})
//
//	It("should only allow cookies to pass through that are whitelisted", func() {
//		request, err := http.NewRequest("GET", GetPageComposerUrl("",""), nil)
//		Expect(err).NotTo(HaveOccurred())
//		request.AddCookie(&http.Cookie{
//			Name: "CompoxureCookie",
//			Value: "Test",
//		})
//		request.AddCookie(&http.Cookie{
//			Name: "AnotherCookie",
//			Value: "Test",
//		})
//		request.AddCookie(&http.Cookie{
//			Name: "TSLCookie",
//			Value: "Test",
//		})
//		request.Header.Add("accept","text/html")
//		response, err := client.Do(request)
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusOK))
//		Expect(err).NotTo(HaveOccurred())
//		doc, err := goquery.NewDocumentFromResponse(response)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(doc.Find("#cookie").Text()).To(BeEquivalentTo("CompoxureCookie=Test; TSLCookie=Test"))
//	})
//
//
//	It("should be able to whitelist even with broken cookies", func() {
//		brokenCookie := "__gads=ID=5217e5ce98e5a5f6:T=1413059056:S=ALNI_MZDmTo6sr27tzMt9RUR65K4xSUWzw; s_fid=79BC0100183D81BE-2708D64605382DEA; TSLCookie=585108831577993685E2ADCF228581BE11AD0DA8B9E378FB8C33DF9B01E21E48C8991D75B61F24E8D7CA2A6A04B2F64B67A6D53A6A375B00EEE705EEADB6ED3FBE04E19D385F5DC89793ADB6978BC6EC17D52A7ED4740D3266C3EDDFCAC2AD881762439AD0485C24B5511984A9D21387921B85193D2689CF6A9B3CCA8CEA4E8939D187CC7327ABC47111A1840C251B1C49DB823713CB866BE0D9958BAAD8CF06D05762525DAD7741272E479BC07CA3D2B35DA1EC2FF8C9284C2996811D4E704573AF8A9E1D4BE609B50A6AC5B29FDC31DCA8460164A44EAB83B730BE565DCC7470EA6C66; TESCookie=XynqF84fIQqO6TMaKPbxsVTGdTQ48cl3KrcYfm0DYZX6eVdcjL9ySX0YHGtk4pqaIJG7TqCiS0%2b6J0bUJgfQR2B7b4AfikEDSl6lrxOdFL9jZQ0vNZuHz9f3Gzr%2f5wu6FSvssSUjGS1paLLxB1UH0idMUHD6RqydZQDVxWpo0BeYg6ZsuSv9XeksslbTqs7FbMetUqSC0JwIRkXsFb6tve7YkunuEg%2fYvrW%2fcsNb1p%2bHXQTWXCKFEa10PMCpXo%2fNw5fV5ofp4svALCnLWUlpO4TDMopHrADRfS3FezOIgQWqES2VQQGBD8lRYWn7ijS%2bUxTzYWBF1b1NWAlGbRORyOAUaq7uS0zvlQ6VuHPca98%3d; TESCookieUser=4241009; tp_ex=0; s_campaign=031114_Secondary-124726_email_et_3112014_%25%25%3DRedirectTo%28%40TOPTENURL7%29%3D%25%25; ET:recipientid=%25%25External_ID%25%25; ET:messageid=_; .TesApplication=9DA9A85E2E258EE23C0537C87F7D4F0DDD37CB5FDDFB44DD230E5CC584B58586EA35644839CA7F75DF6EC079ECFE5B99BE7C3E36EE93A651BA365EE935D7A16EE08793AB021FC95537FD5079CD75BB56EE5A2D438CB8B2F47C3AA3C4EE0C9B2DBE361889F1DD75E0D2F967193449D61191A2F75BEF3D2608CC75620EAE313938BA52495555F785ED8B8FA393FC84D7360D19507576B1BDB0A999B31835360C84B8F023AED31CCA8910BC13FDEF3476006C9FD16C11FBC133E67F1EC958332DF86447EDEFDC3AD59EDC4CB183B49D1F081AC586178FD3D2BCD9BDB16E561F70BD94E73EE404024542DD2DAFA317DCD5B310A79ABC441B01B44A8E3D5FFE922BE389AE91E41FDCB5F2A4FFBC6994812E769BC657007A26414CC2BD7EE68AC3EDD630D076B28048B428ECF42598DEDE9427CA3CAA856CDD46ACE57B85E8846A8674E37D75BCB29ABAAEB227F8EE6C996D994E0B06DF; __utmt_UA-13200995-3=1; s_cc=true; s_sq=%5B%5BB%5D%5D; __utma=233401627.2136099593.1404067931.1416513050.1416513139.14; __utmb=233401627.5.10.1416513139; __utmc=233401627; __utmz=233401627.1416513139.14.12.utmcsr=ET|utmccn=031114_Secondary_124726_%%=RedirectTo(@TOPTENURL7)=%%|utmcmd=email|utmctr=3112014|utmcct=_; __atuvc=0%7C43%2C0%7C44%2C7%7C45%2C0%7C46%2C3%7C47; __atuvs=546e4672cc74592b002"
//		tslCookie := "TSLCookie=585108831577993685E2ADCF228581BE11AD0DA8B9E378FB8C33DF9B01E21E48C8991D75B61F24E8D7CA2A6A04B2F64B67A6D53A6A375B00EEE705EEADB6ED3FBE04E19D385F5DC89793ADB6978BC6EC17D52A7ED4740D3266C3EDDFCAC2AD881762439AD0485C24B5511984A9D21387921B85193D2689CF6A9B3CCA8CEA4E8939D187CC7327ABC47111A1840C251B1C49DB823713CB866BE0D9958BAAD8CF06D05762525DAD7741272E479BC07CA3D2B35DA1EC2FF8C9284C2996811D4E704573AF8A9E1D4BE609B50A6AC5B29FDC31DCA8460164A44EAB83B730BE565DCC7470EA6C66"
//
//		response := Get(GetPageComposerUrl("",""), map[string]string{"accept":"text/html","cookie":brokenCookie})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusOK))
//		doc, err := goquery.NewDocumentFromResponse(response)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(doc.Find("#cookie").Text()).To(BeEquivalentTo(tslCookie))
//	})
//
//	It("should create a default handler if none provided", func() {
//		go common.StartPageCompositionServer(5003,"localhost", func() {},"")
//	})
//
//	It("should allow use of variables in a backend target", func() {
//		response := Get(GetPageComposerUrl("variabletarget",""), map[string]string{"accept":"text/html"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusOK))
//		doc, err := goquery.NewDocumentFromResponse(response)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(doc.Find("#declarative").Text()).To(BeEquivalentTo("Replaced"))
//	})
//
//
//	It("should use cached headers when a backend 500s", func() {
//		response := Get(GetPageComposerUrl("alternate500",""), map[string]string{"accept":"text/html"})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusOK))
//
//		time.AfterFunc(50 * time.Millisecond, func() {
//			response := Get(GetPageComposerUrl("alternate500",""), map[string]string{"accept":"text/html"})
//			Expect(response.StatusCode).To(BeEquivalentTo(http.StatusOK))
//			doc, err := goquery.NewDocumentFromResponse(response)
//			Expect(err).NotTo(HaveOccurred())
//			bundles := doc.Find(".bundle")
//			Expect(bundles.First().Text()).To(BeEquivalentTo("service-one >> 100 >> top.js.html"))
//
//		})
//	})
//
//	It("should not completely die with broken cookies", func() {
//		brokenCookie := "__gads=ID=5217e5ce98e5a5f6:T=1413059056:S=ALNI_MZDmTo6sr27tzMt9RUR65K4xSUWzw; s_fid=79BC0100183D81BE-2708D64605382DEA; TSLCookie=585108831577993685E2ADCF228581BE11AD0DA8B9E378FB8C33DF9B01E21E48C8991D75B61F24E8D7CA2A6A04B2F64B67A6D53A6A375B00EEE705EEADB6ED3FBE04E19D385F5DC89793ADB6978BC6EC17D52A7ED4740D3266C3EDDFCAC2AD881762439AD0485C24B5511984A9D21387921B85193D2689CF6A9B3CCA8CEA4E8939D187CC7327ABC47111A1840C251B1C49DB823713CB866BE0D9958BAAD8CF06D05762525DAD7741272E479BC07CA3D2B35DA1EC2FF8C9284C2996811D4E704573AF8A9E1D4BE609B50A6AC5B29FDC31DCA8460164A44EAB83B730BE565DCC7470EA6C66; TESCookie=XynqF84fIQqO6TMaKPbxsVTGdTQ48cl3KrcYfm0DYZX6eVdcjL9ySX0YHGtk4pqaIJG7TqCiS0%2b6J0bUJgfQR2B7b4AfikEDSl6lrxOdFL9jZQ0vNZuHz9f3Gzr%2f5wu6FSvssSUjGS1paLLxB1UH0idMUHD6RqydZQDVxWpo0BeYg6ZsuSv9XeksslbTqs7FbMetUqSC0JwIRkXsFb6tve7YkunuEg%2fYvrW%2fcsNb1p%2bHXQTWXCKFEa10PMCpXo%2fNw5fV5ofp4svALCnLWUlpO4TDMopHrADRfS3FezOIgQWqES2VQQGBD8lRYWn7ijS%2bUxTzYWBF1b1NWAlGbRORyOAUaq7uS0zvlQ6VuHPca98%3d; TESCookieUser=4241009; tp_ex=0; s_campaign=031114_Secondary-124726_email_et_3112014_%25%25%3DRedirectTo%28%40TOPTENURL7%29%3D%25%25; ET:recipientid=%25%25External_ID%25%25; ET:messageid=_; .TesApplication=9DA9A85E2E258EE23C0537C87F7D4F0DDD37CB5FDDFB44DD230E5CC584B58586EA35644839CA7F75DF6EC079ECFE5B99BE7C3E36EE93A651BA365EE935D7A16EE08793AB021FC95537FD5079CD75BB56EE5A2D438CB8B2F47C3AA3C4EE0C9B2DBE361889F1DD75E0D2F967193449D61191A2F75BEF3D2608CC75620EAE313938BA52495555F785ED8B8FA393FC84D7360D19507576B1BDB0A999B31835360C84B8F023AED31CCA8910BC13FDEF3476006C9FD16C11FBC133E67F1EC958332DF86447EDEFDC3AD59EDC4CB183B49D1F081AC586178FD3D2BCD9BDB16E561F70BD94E73EE404024542DD2DAFA317DCD5B310A79ABC441B01B44A8E3D5FFE922BE389AE91E41FDCB5F2A4FFBC6994812E769BC657007A26414CC2BD7EE68AC3EDD630D076B28048B428ECF42598DEDE9427CA3CAA856CDD46ACE57B85E8846A8674E37D75BCB29ABAAEB227F8EE6C996D994E0B06DF; __utmt_UA-13200995-3=1; s_cc=true; s_sq=%5B%5BB%5D%5D; __utma=233401627.2136099593.1404067931.1416513050.1416513139.14; __utmb=233401627.5.10.1416513139; __utmc=233401627; __utmz=233401627.1416513139.14.12.utmcsr=ET|utmccn=031114_Secondary_124726_%%=RedirectTo(@TOPTENURL7)=%%|utmcmd=email|utmctr=3112014|utmcct=_; __atuvc=0%7C43%2C0%7C44%2C7%7C45%2C0%7C46%2C3%7C47; __atuvs=546e4672cc74592b002"
//		go common.StartPageCompositionServer(5003,"localhost",func() {}, "noWhitelist" )
//
//		response := Get("http://localhost:5003/", map[string]string{"accept":"text/html","cookie":brokenCookie})
//		Expect(response.StatusCode).To(BeEquivalentTo(http.StatusOK))
//		doc, err := goquery.NewDocumentFromResponse(response)
//		Expect(err).NotTo(HaveOccurred())
//		Expect(doc.Find("#cookie").Text()).To(BeEquivalentTo(brokenCookie))
//	})
//})
//
//
//func GetPageComposerUrl(path string, search string) string {
//	var url url.URL
//	url.Scheme = "http"
//	url.Path = path
//	url.Host = "localhost:" + strconv.Itoa(PAGE_COMPOSITION_PORT)
//	url.RawQuery = search
//	return url.String()
//}
//
//func Get(url string, headers map[string]string) *http.Response {
//	request, err := http.NewRequest("GET", url, nil)
//	Expect(err).NotTo(HaveOccurred())
//	for key,value := range headers {
//		request.Header.Add(key,value)
//	}
//	response, err := client.Do(request)
//	Expect(err).NotTo(HaveOccurred())
//	return response
//}
//
//func GetSection(path string, search string, query string, next func(string)) {
//	response := Get(GetPageComposerUrl(path,search), map[string]string{"accept":"text/html"})
//	doc, err := goquery.NewDocumentFromResponse(response)
//	Expect(err).NotTo(HaveOccurred())
//	next(doc.Find(query).Text())
//}
//
