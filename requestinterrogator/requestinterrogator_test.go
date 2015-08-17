package requestinterrogator

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"net/http"
	"net/url"
	"github.com/caledhwa/gongeal/config"
	"os"
	"encoding/json"
)

var configuration config.Config



func TestRequestInterrogator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestInterrogatorSuite")
}

var _ = BeforeSuite(func() {
	configFile, err := os.Open("../test/common/testConfig.json")
	Expect(err).NotTo(HaveOccurred())
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&configuration)
	Expect(err).NotTo(HaveOccurred())
})

var _ = Describe("Request Interrogator", func() {

	It("should generate the url object", func() {
		req, _ := http.NewRequest("GET","/teaching-resource/Queen-Elizabeth-II-Diamond-jubilee-2012-6206420", nil)
		req.Header.Add("host","localhost:5000")
		interrogator := NewRequestInterrogator(&configuration) //TODO {name:'test'}
		params := interrogator.InterrogateRequest(req)
		expectedPageUrl := "http://localhost:5000/teaching-resource/Queen-Elizabeth-II-Diamond-jubilee-2012-6206420"
		encodedExpectedPageUrl, _ := UrlEncoded(expectedPageUrl)
		Expect(params).To(HaveKeyWithValue("url:href",expectedPageUrl))
		Expect(params).To(HaveKeyWithValue("url:href:encoded",encodedExpectedPageUrl))
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

// UrlEncoded encodes a string like Javascript's encodeURIComponent()
func UrlEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
