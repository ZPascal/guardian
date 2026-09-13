package bundlerules

import (
	"strings"

	spec "code.cloudfoundry.org/guardian/gardener/container-spec"
	"code.cloudfoundry.org/guardian/rundmc/goci"
)

// CDIAnnotationKey is the OCI annotation that runc's native CDI resolution
// (runc >= 1.2, invoked with --cdi-spec-dirs) reads to find which CDI
// devices to inject into the container. See the CDI spec's annotation
// convention: https://github.com/cncf-tags/container-device-interface.
const CDIAnnotationKey = "cdi.k8s.io/gpu"

type CDIDevices struct {
}

func (r CDIDevices) Apply(bndl goci.Bndl, spec spec.DesiredContainerSpec) (goci.Bndl, error) {
	if len(spec.CDIDevices) == 0 {
		return bndl, nil
	}

	return bndl.WithAnnotations(map[string]string{
		CDIAnnotationKey: strings.Join(spec.CDIDevices, ","),
	}), nil
}
