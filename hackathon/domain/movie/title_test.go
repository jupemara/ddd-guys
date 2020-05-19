package movie_test

import (
	"testing"

	"github.com/jupemara/ddd-guys/hackathon/domain/movie"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTitleBdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "create Title instance test")
}

var _ = Describe("NewTitle", func() {
	Context("with empty Title", func() {
		It("fails to create Title and returns error", func() {
			titleValue := ""
			title, err := movie.NewTitle(titleValue)
			Expect(title).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
	})
	Context("with valid Title", func() {
		It("creates Title without error", func() {
			titleValue := "valid-title"
			title, err := movie.NewTitle(titleValue)
			Expect(title).NotTo(BeNil())
			Expect(err).To(BeNil())
			Expect(title.Value()).To(Equal(titleValue))
		})
	})

	// MaxTitleLength+1 strings for test
	var longTitle string
	for i := 0; i <= movie.MaxTitleLength; i++ {
		longTitle = longTitle + "A"
	}

	Context("with a title longer than the criteria", func() {
		It("fails to create Title and returns error", func() {
			title, err := movie.NewTitle(longTitle)
			Expect(title).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
	})
})
