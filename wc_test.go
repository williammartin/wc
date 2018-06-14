package wc_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"github.com/williammartin/wc"
)

var _ = Describe("World Cup Client", func() {

	var server *ghttp.Server
	var client *wc.Client

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = wc.NewClient(server.URL())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("Matches", func() {

		Context("when fetching all matches", func() {
			var (
				matches wc.Matches
				err     error
			)

			BeforeEach(func() {
				returnedMatches := wc.Matches{
					{Venue: "Moscow"},
					{Venue: "My House"},
				}
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/matches/"),
						ghttp.RespondWithJSONEncoded(http.StatusOK, returnedMatches),
					),
				)

				matches, err = client.MatchService.Fetch()
			})

			It("doesn't error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("hits the correct endpoint", func() {
				Expect(server.ReceivedRequests()).Should(HaveLen(1))
			})

			It("returns all the matches available", func() {
				Expect(matches).To(HaveLen(2))
				Expect(matches[0].Venue).To(Equal("Moscow"))
				Expect(matches[1].Venue).To(Equal("My House"))
			})
		})
	})

})
