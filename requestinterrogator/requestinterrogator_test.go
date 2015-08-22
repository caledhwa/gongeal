package requestinterrogator

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"net/http"
	"github.com/caledhwa/gongeal/config"
	"github.com/caledhwa/gongeal/util"
	"os"
	"encoding/json"
)

var configParams config.Parameters
var configCdn config.Cdn


func TestRequestInterrogator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestInterrogatorSuite")
}

var _ = BeforeSuite(func() {
	configFile, err := os.Open("../test/common/testConfig.json")
	Expect(err).NotTo(HaveOccurred())
	jsonParser := json.NewDecoder(configFile)
	var configuration config.Config
	err = jsonParser.Decode(&configuration)
	configParams = configuration.Parameters
	configCdn = configuration.Cdn
	Expect(err).NotTo(HaveOccurred())
})

var _ = Describe("Request Interrogator", func() {

	It("should generate the url object", func() {
		req, _ := http.NewRequest("GET","/teaching-resource/Queen-Elizabeth-II-Diamond-jubilee-2012-6206420", nil)
		req.Header.Add("host","localhost:5000")
		interrogator := NewRequestInterrogator(&configParams, &configCdn)
		params := interrogator.InterrogateRequest(req)
		expectedPageUrl := "http://localhost:5000/teaching-resource/Queen-Elizabeth-II-Diamond-jubilee-2012-6206420"
		encodedExpectedPageUrl, _ := util.EncodeUrl(expectedPageUrl)
		Expect(params).To(HaveKeyWithValue("url:href",expectedPageUrl))
		Expect(params).To(HaveKeyWithValue("url:href:encoded",encodedExpectedPageUrl))
	})

	It("should extract parameters from the query", func() {
		req, _ := http.NewRequest("GET","/teaching-resource?storyCode=2206421", nil)
		req.Header.Add("host","localhost:5000")
		testConfig := &config.Parameters { Query: []config.Query { config.Query {"storyCode", "resourceId"} }}
		interrogator := NewRequestInterrogator(testConfig, &configCdn)
		params := interrogator.InterrogateRequest(req)
		Expect(params).To(HaveKeyWithValue("param:resourceId","2206421"))
	})

	It("should only extract parameters from the query when they're not empty", func() {
		req, _ := http.NewRequest("GET","/teaching-resource?storycode=2206421", nil)
		req.Header.Add("host","localhost:5000")
		testConfig := &config.Parameters { Query: []config.Query { config.Query {"storyCode", "resourceId"}, config.Query{"storycode", "resourceId"} }}
		interrogator := NewRequestInterrogator(testConfig, &configCdn)
		params := interrogator.InterrogateRequest(req)
		Expect(params).To(HaveKeyWithValue("param:resourceId","2206421"))
	})

	It("should extract parameters from the path", func() {
		req, _ := http.NewRequest("GET","/teaching-resource/Queen-Elizabeth-II-Diamond-jubilee-2012-6206420", nil)
		req.Header.Add("host","localhost:5000")
		testConfig := &config.Parameters { Urls: []config.Url { config.Url { Pattern: "/teaching-resource/(.*)-(\\d+)", Names: []string {"blurb","resourceId"}}}}
		interrogator := NewRequestInterrogator(testConfig, &configCdn)
		params := interrogator.InterrogateRequest(req)
		Expect(params).To(HaveKeyWithValue("param:resourceId","6206420"))
		Expect(params).To(HaveKeyWithValue("param:blurb","Queen-Elizabeth-II-Diamond-jubilee-2012"))
	})

	It("should extract parametrs from the path if multiple paths match", func() {
		req, _ := http.NewRequest("GET","/teaching-resource/Queen-Elizabeth-II-Diamond-jubilee-2012-6206420", nil)
		req.Header.Add("host","localhost:5000")
		testConfig := &config.Parameters { Urls: []config.Url { config.Url { Pattern: "/teaching-resource/.*-(\\d+)", Names: []string {"resourceId"}},
																config.Url { Pattern: "/teaching-resource/(.*)-\\d+", Names: []string {"blurb"}}}}
		interrogator := NewRequestInterrogator(testConfig, &configCdn)
		params := interrogator.InterrogateRequest(req)
		Expect(params).To(HaveKeyWithValue("param:resourceId","6206420"))
		Expect(params).To(HaveKeyWithValue("param:blurb","Queen-Elizabeth-II-Diamond-jubilee-2012"))
	})

	It("should extract parameters if multiple overlap it takes the last one", func() {
		req, _ := http.NewRequest("GET","/teaching-resource/Queen-Elizabeth-II-Diamond-jubilee-2012-6206420", nil)
		req.Header.Add("host","localhost:5000")
		testConfig := &config.Parameters { Urls: []config.Url { config.Url { Pattern: "/teaching-resource/.*-(\\d+)", Names: []string {"resourceId"}},
																config.Url { Pattern: "/teaching-resource/(.*)", Names: []string {"blurb"}},
																config.Url { Pattern: "/teaching-resource/(.*)-\\d+", Names: []string {"blurb"}}}}
		interrogator := NewRequestInterrogator(testConfig, &configCdn)
		params := interrogator.InterrogateRequest(req)
		Expect(params).To(HaveKeyWithValue("param:resourceId","6206420"))
		Expect(params).To(HaveKeyWithValue("param:blurb","Queen-Elizabeth-II-Diamond-jubilee-2012"))
	})

	It("should extract query parameters", func() {
		req, _ := http.NewRequest("GET","/teaching-resource/Queen-Elizabeth-II-Diamond-jubilee-2012-6206420?foo=bar", nil)
		req.Header.Add("host","localhost:5000")
		interrogator := NewRequestInterrogator(&configParams, &configCdn)
		params := interrogator.InterrogateRequest(req)
		Expect(params).To(HaveKeyWithValue("query:foo","bar"))
	})

	It("should extract query parameters with multiple values separated", func() {
		req, _ := http.NewRequest("GET","/teaching-resource/Queen-Elizabeth-II-Diamond-jubilee-2012-6206420?foo=bar&foo=monkey", nil)
		req.Header.Add("host","localhost:5000")
		interrogator := NewRequestInterrogator(&configParams, &configCdn)
		params := interrogator.InterrogateRequest(req)
		Expect(params).To(HaveKeyWithValue("query:foo","bar,monkey"))
	})

	It("should extract headers", func() {
		req, _ := http.NewRequest("GET","/teaching-resource/Queen-Elizabeth-II-Diamond-jubilee-2012-6206420", nil)
		req.Header.Add("host","localhost:5000")
		req.Header.Add("foo","bar")
		interrogator := NewRequestInterrogator(&configParams, &configCdn)
		params := interrogator.InterrogateRequest(req)
		Expect(params).To(HaveKeyWithValue("header:foo","bar"))
	})

	It("should extract headers as lowercase params", func() {
		req, _ := http.NewRequest("GET","/teaching-resource/Queen-Elizabeth-II-Diamond-jubilee-2012-6206420", nil)
		req.Header.Add("host","localhost:5000")
		req.Header.Add("Foo","bar")
		interrogator := NewRequestInterrogator(&configParams, &configCdn)
		params := interrogator.InterrogateRequest(req)
		Expect(params).To(HaveKeyWithValue("header:foo","bar"))
	})

	It("should extract cookies", func() {
		req, _ := http.NewRequest("GET","/teaching-resource/Queen-Elizabeth-II-Diamond-jubilee-2012-6206420", nil)
		req.Header.Add("host","localhost:5000")
		req.AddCookie(&http.Cookie{
			Name: "foo",
			Value: "bar",
		})
		interrogator := NewRequestInterrogator(&configParams, &configCdn)
		params := interrogator.InterrogateRequest(req)
		Expect(params).To(HaveKeyWithValue("cookie:foo","bar"))
	})

	It("should default user:userId if not logged in", func() {
		//TODO: Support user on Request???
	})

	It("should get user from request", func() {
		//TODO: Support user on Request???
	})

	It("should parse cdn url configuration using template variables", func() {
		req, _ := http.NewRequest("GET","/teaching-resource/Queen-Elizabeth-II-Diamond-jubilee-2012-6206420", nil)
		req.Header.Add("host","localhost:5000")
		interrogator := NewRequestInterrogator(&configParams, &configCdn)
		params := interrogator.InterrogateRequest(req)
		Expect(params).To(HaveKeyWithValue("cdn:url","http://localhost:5001/static/"))
	})
})
