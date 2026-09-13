package nerd_test

import (
	"code.cloudfoundry.org/guardian/rundmc/runcontainerd/nerd"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

var _ = Describe("CDIDevicesFromSpec", func() {
	It("splits the cdi.k8s.io/gpu annotation on commas", func() {
		spec := &specs.Spec{Annotations: map[string]string{
			"cdi.k8s.io/gpu": "nvidia.com/gpu=0,nvidia.com/gpu=1",
		}}
		Expect(nerd.CDIDevicesFromSpec(spec)).To(Equal([]string{"nvidia.com/gpu=0", "nvidia.com/gpu=1"}))
	})

	It("returns nil when the annotation is absent", func() {
		Expect(nerd.CDIDevicesFromSpec(&specs.Spec{})).To(BeNil())
	})

	It("returns nil when Annotations itself is nil", func() {
		Expect(nerd.CDIDevicesFromSpec(&specs.Spec{Annotations: nil})).To(BeNil())
	})
})
