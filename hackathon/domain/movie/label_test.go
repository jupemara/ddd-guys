package movie_test

import (
	"testing"

	"github.com/jupemara/ddd-guys/hackathon/domain/movie"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLabelBdd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "create Label instance test")
}

var _ = Describe("NewLabel", func() {
	Context("with empty Label", func() {
		It("fails to create Label and returns error", func() {
			labelValue := ""
			label, err := movie.NewLabel(idValue)
			Expect(label).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
	})
	Context("with valid Label", func() {
		It("creates Label without error", func() {
			labelValue := "valid-label"
			label, err := movie.NewLabel(labelValue)
			Expect(label).NotTo(BeNil())
			Expect(err).To(BeNil())
			Expect(label.Value()).To(Equal(labelValue))
		})
	})
	Context("with a label shorter than the criteria", func() {
		It("fails to create Label and returns error", func() {
			labelValue := "l"
			label, err := movie.NewLabel(labelValue)
			Expect(id).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
	})
	Context("with a label longer than the criteria", func() {
		It("fails to create Label and returns error", func() {
			labelValue := "label-label-label-label"
			label, err := movie.NewLabel(labelValue)
			Expect(id).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
	})
})
