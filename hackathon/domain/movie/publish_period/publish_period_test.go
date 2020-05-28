package publish_period_test

import (
	"testing"

	"github.com/jupemara/ddd-guys/hackathon/domain/movie/publish_period"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPublishedPeriodBdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "create published period instance test")
}

var _ = Describe("NewPublishPeriod", func() {
	var (
		contentsProvider *publish_period.ContentsProvider
	)
	BeforeEach(func() {
		var err error
		contentsProvider, err = publish_period.NewContentsProvider("valid-cp")
		Expect(err).To(BeNil())
	})
	Context("with valid startDate, endDate and endDate is after startDate", func() {
		It("creates PublishPeriod without error", func() {
			start := "2020-01-01T00:00:00.00Z"
			end := "2020-01-01T01:00:00.00Z"
			pp, err := publish_period.NewPublishPeriod(*contentsProvider, start, end)
			Expect(pp).NotTo(BeNil())
			Expect(err).To(BeNil())
		})
	})
	Context("with valid startDate, endDate and endDate is before startDate", func() {
		It("fails to create PublishPeriod with error", func() {
			start := "2020-01-01T01:00:00.00Z"
			end := "2020-01-01T00:00:00.00Z"
			pp, err := publish_period.NewPublishPeriod(*contentsProvider, start, end)
			Expect(pp).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
	})
	Context("with no startDate and valid endDate", func() {
		It("creates PublishPeriod with no error", func() {
			start := ""
			end := "2020-01-01T00:00:00.00Z"
			pp, err := publish_period.NewPublishPeriod(*contentsProvider, start, end)
			Expect(pp).NotTo(BeNil())
			Expect(err).To(BeNil())
		})
	})
	Context("with valid startDate and no endDate", func() {
		It("creates PublishPeriod with no error", func() {
			start := "2020-01-01T00:00:00.00Z"
			end := ""
			pp, err := publish_period.NewPublishPeriod(*contentsProvider, start, end)
			Expect(pp).NotTo(BeNil())
			Expect(err).To(BeNil())
		})
	})
	Context("with no startDate and no endDate", func() {
		It("creates PublishPeriod with no error", func() {
			start := ""
			end := ""
			pp, err := publish_period.NewPublishPeriod(*contentsProvider, start, end)
			Expect(pp).NotTo(BeNil())
			Expect(err).To(BeNil())
		})
	})
})
