package wc_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"github.com/williammartin/wc"
)

var _ = Describe("Client", func() {

	var server *ghttp.Server
	var client *wc.Client

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = wc.NewClient(server.URL())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("matches", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/matches/"),
				),
			)

			client.MatchService.Fetch()
		})

		It("hits the correct endpoint", func() {
			Expect(server.ReceivedRequests()).Should(HaveLen(1))
		})
	})
})
