package services_test

import (
	"net/http"

	"github.com/dghubble/sling"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"github.com/williammartin/wc/services"
)

var _ = Describe("MatchService", func() {

	var server *ghttp.Server
	var matchService *services.MatchService

	BeforeEach(func() {
		server = ghttp.NewServer()
		matchService = services.NewMatchService(sling.New().Base(server.URL()))
	})

	AfterEach(func() {
		server.Close()
	})

	Context("when fetching all matches", func() {
		var (
			matches services.Matches
			err     error
		)

		BeforeEach(func() {
			returnedMatches := services.Matches{
				{Venue: "Moscow"},
				{Venue: "My House"},
			}
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/matches/"),
					ghttp.RespondWithJSONEncoded(http.StatusOK, returnedMatches),
				),
			)

			matches, err = matchService.Fetch()
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
