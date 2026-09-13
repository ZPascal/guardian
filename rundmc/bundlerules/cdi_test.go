package bundlerules_test

import (
	spec "code.cloudfoundry.org/guardian/gardener/container-spec"
	"code.cloudfoundry.org/guardian/rundmc/bundlerules"
	"code.cloudfoundry.org/guardian/rundmc/goci"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CDIDevices Rule", func() {
	var (
		cdiDevices []string
		newBndl    goci.Bndl
		rule       bundlerules.CDIDevices
	)

	JustBeforeEach(func() {
		var err error
		rule = bundlerules.CDIDevices{}
		newBndl, err = rule.Apply(goci.Bundle(), spec.DesiredContainerSpec{
			CDIDevices: cdiDevices,
		})
		Expect(err).NotTo(HaveOccurred())
	})

	Context("with one CDI device requested", func() {
		BeforeEach(func() {
			cdiDevices = []string{"nvidia.com/gpu=0"}
		})

		It("sets the cdi.k8s.io/gpu annotation to the device name", func() {
			Expect(newBndl.Spec.Annotations).To(HaveKeyWithValue("cdi.k8s.io/gpu", "nvidia.com/gpu=0"))
		})
	})

	Context("with multiple CDI devices requested", func() {
		BeforeEach(func() {
			cdiDevices = []string{"nvidia.com/gpu=0", "nvidia.com/gpu=1"}
		})

		It("joins them with a comma", func() {
			Expect(newBndl.Spec.Annotations).To(HaveKeyWithValue("cdi.k8s.io/gpu", "nvidia.com/gpu=0,nvidia.com/gpu=1"))
		})
	})

	Context("with no CDI devices requested", func() {
		BeforeEach(func() {
			cdiDevices = nil
		})

		It("does not set any annotation", func() {
			Expect(newBndl.Spec.Annotations).To(BeEmpty())
		})
	})
})
