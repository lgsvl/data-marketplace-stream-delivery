package resources_test

import (
	"context"
	"log"
	"os"

	"github.com/lgsvl/data-marketplace-stream-delivery/resources"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("pull-push-controller", func() {
	var (
		logger   *log.Logger
		ctx      context.Context
		registry resources.Registry
		stream   resources.DataStream
	)
	BeforeEach(func() {
		logger = log.New(os.Stdout, "data-stream-delivery-registry-test: ", log.Lshortfile|log.LstdFlags)
		ctx = context.Background()

		registry = resources.NewRegistry(ctx, logger)
		stream = resources.DataStream{DataContractTypeID: "123",
			DataStreamSourceURL:      "fakeSRC",
			DataStreamDestinationURL: "fakeDest",
			DataStreamProtocol:       "fakeProtocl",
			DataStreamTopic:          "fakeTopic"}

	})

	Context(".AddStream", func() {
		It("should succeed when streamID does not exist", func() {

			err := registry.AddStream(ctx, stream)
			Expect(err).NotTo(HaveOccurred())

		})
		It("should fail when streamID exists", func() {

			err := registry.AddStream(ctx, stream)
			Expect(err).NotTo(HaveOccurred())
			err = registry.AddStream(ctx, stream)
			Expect(err).To(HaveOccurred())

		})
	})

	Context(".UpdateStream", func() {
		It("should fail when streamID does not exist", func() {

			err := registry.UpdateStream(ctx, stream)
			Expect(err).To(HaveOccurred())

		})
		It("should succeed when streamID exists", func() {
			err := registry.AddStream(ctx, stream)
			Expect(err).NotTo(HaveOccurred())
			stream = resources.DataStream{DataContractTypeID: "123",
				DataStreamSourceURL:      "fakeSRC2",
				DataStreamDestinationURL: "fakeDest2",
				DataStreamProtocol:       "fakeProtocl2",
				DataStreamTopic:          "fakeTopic2"}
			err = registry.UpdateStream(ctx, stream)
			Expect(err).NotTo(HaveOccurred())
			str, err := registry.GetStream(ctx, "123")
			Expect(err).NotTo(HaveOccurred())
			Expect(str.DataStreamDestinationURL).To(Equal("fakeDest2"))

		})
	})

	Context(".HasStream", func() {
		It("should return false when streamID does not exist", func() {

			exists := registry.HasStream(ctx, "123")
			Expect(exists).NotTo(BeTrue())

		})
		It("should return true when streamID exists", func() {
			err := registry.AddStream(ctx, stream)
			Expect(err).NotTo(HaveOccurred())
			exists := registry.HasStream(ctx, "123")
			Expect(exists).To(BeTrue())

		})
	})

	Context(".GetStream", func() {
		It("should fail when streamID does not exist", func() {

			_, err := registry.GetStream(ctx, "123")
			Expect(err).To(HaveOccurred())

		})
		It("should return the stream when streamID exists", func() {
			err := registry.AddStream(ctx, stream)
			Expect(err).NotTo(HaveOccurred())
			str, err := registry.GetStream(ctx, "123")
			Expect(err).NotTo(HaveOccurred())
			Expect(str.DataContractTypeID).To(Equal("123"))

		})
	})

	Context(".ListStream", func() {
		It("should return an empty list when the registry is empty", func() {

			res := registry.ListStreams()
			Expect(len(res)).To(Equal(0))

		})
		It("should return the list of streams when registry is not empty", func() {
			err := registry.AddStream(ctx, stream)
			Expect(err).NotTo(HaveOccurred())
			stream = resources.DataStream{DataContractTypeID: "321",
				DataStreamSourceURL:      "fakeSRC2",
				DataStreamDestinationURL: "fakeDest2",
				DataStreamProtocol:       "fakeProtocl2",
				DataStreamTopic:          "fakeTopic2"}
			err = registry.AddStream(ctx, stream)
			Expect(err).NotTo(HaveOccurred())
			res := registry.ListStreams()
			Expect(len(res)).To(Equal(2))

		})
	})

	Context(".DeleteStream", func() {
		It("should fail when streamID does not exist", func() {

			err := registry.DeleteStream(ctx, "123")
			Expect(err).To(HaveOccurred())

		})
		It("should succeed when streamID exists", func() {
			err := registry.AddStream(ctx, stream)
			Expect(err).NotTo(HaveOccurred())
			err = registry.DeleteStream(ctx, "123")
			Expect(err).NotTo(HaveOccurred())
			Expect(registry.HasStream(ctx, "123")).NotTo(BeTrue())

		})
	})

	Context(".AddChannel", func() {
		It("should fail when streamID does not exist in streams", func() {
			channel := make(chan string)
			err := registry.AddChannel(ctx, "123", channel)
			Expect(err).To(HaveOccurred())

		})
		It("should succeed when streamID does not exist in channels map", func() {
			err := registry.AddStream(ctx, stream)
			Expect(err).NotTo(HaveOccurred())
			channel := make(chan string)
			err = registry.AddChannel(ctx, "123", channel)
			Expect(err).NotTo(HaveOccurred())

		})
		It("should fail when streamID exists in channels map", func() {
			err := registry.AddStream(ctx, stream)
			Expect(err).NotTo(HaveOccurred())
			channel := make(chan string)
			err = registry.AddChannel(ctx, "123", channel)
			Expect(err).NotTo(HaveOccurred())
			err = registry.AddChannel(ctx, "123", channel)
			Expect(err).To(HaveOccurred())

		})
	})

	Context(".GetChannel", func() {
		It("should fail when streamID does not exist in channels map", func() {
			err := registry.DeleteChannel(ctx, "123")
			Expect(err).To(HaveOccurred())

		})
		It("should succeed when streamID exists in channels map", func() {
			err := registry.AddStream(ctx, stream)
			Expect(err).NotTo(HaveOccurred())
			channel := make(chan string)
			err = registry.AddChannel(ctx, "123", channel)
			Expect(err).NotTo(HaveOccurred())
			channel, err = registry.GetChannel(ctx, "123")
			Expect(err).NotTo(HaveOccurred())
			Expect(channel).NotTo(BeNil())

		})
	})

	Context(".DeleteChannel", func() {
		It("should fail when streamID does not exist in channels map", func() {
			err := registry.DeleteChannel(ctx, "123")
			Expect(err).To(HaveOccurred())

		})
		It("should succeed when streamID exists in channels map", func() {
			err := registry.AddStream(ctx, stream)
			Expect(err).NotTo(HaveOccurred())
			channel := make(chan string)
			err = registry.AddChannel(ctx, "123", channel)
			Expect(err).NotTo(HaveOccurred())
			err = registry.DeleteChannel(ctx, "123")
			Expect(err).NotTo(HaveOccurred())
			channel, err = registry.GetChannel(ctx, "123")
			Expect(err).To(HaveOccurred())
			Expect(channel).To(BeNil())

		})
	})
})
