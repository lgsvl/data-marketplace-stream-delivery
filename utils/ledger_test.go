package utils_test

import (
	"context"
	"log"
	"os"

	"github.com/jarcoal/httpmock"
	"github.com/lgsvl/data-marketplace-stream-delivery/resources"
	"github.com/lgsvl/data-marketplace-stream-delivery/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hyperledger", func() {

	var (
		logger           *log.Logger
		ctx              context.Context
		endpoint         string
		blockchainHelper utils.BlockchainClient
		auth             map[string]string
	)
	BeforeEach(func() {
		logger = log.New(os.Stdout, "data-stream-delivery-utils-test", log.Lshortfile|log.LstdFlags)
		ctx = context.Background()
		endpoint = utils.FormatURL("http://fakeendpoint.com")
		blockchainHelper = utils.NewHyperledgerClientWithEndpoint(ctx, logger, endpoint)
		auth = map[string]string{"authorization": "fake-auth"}
	})

	Context(".CheckContractID", func() {
		It("should fail when http execute does not return 200 ok", func() {
			httpmock.RegisterResponder("GET", utils.FormatURL(endpoint, resources.CheckContractIDPath, "123"),
				httpmock.NewStringResponder(404, `[{}]`))

			resp, err := blockchainHelper.CheckContractID("123", auth)
			Expect(err).To(HaveOccurred())
			Expect(resp).NotTo(BeTrue())

		})
		It("should succeed when http execute returns 200 ok", func() {
			httpmock.RegisterResponder("GET", utils.FormatURL(endpoint, resources.CheckContractIDPath, "123"),
				httpmock.NewStringResponder(200, `[{}]`))

			resp, err := blockchainHelper.CheckContractID("123", auth)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp).To(BeTrue())

		})

	})
})
