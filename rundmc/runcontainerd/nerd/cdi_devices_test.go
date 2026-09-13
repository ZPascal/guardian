package nerd

import (
	"reflect"
	"testing"

	specs "github.com/opencontainers/runtime-spec/specs-go"
)

func TestCDIDevicesFromSpec(t *testing.T) {
	t.Run("splits the cdi.k8s.io/gpu annotation on commas", func(t *testing.T) {
		spec := &specs.Spec{Annotations: map[string]string{
			"cdi.k8s.io/gpu": "nvidia.com/gpu=0,nvidia.com/gpu=1",
		}}
		got := CDIDevicesFromSpec(spec)
		want := []string{"nvidia.com/gpu=0", "nvidia.com/gpu=1"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("CDIDevicesFromSpec(%v) = %v, want %v", spec, got, want)
		}
	})

	t.Run("returns nil when the annotation is absent", func(t *testing.T) {
		spec := &specs.Spec{}
		got := CDIDevicesFromSpec(spec)
		if got != nil {
			t.Errorf("CDIDevicesFromSpec(%v) = %v, want nil", spec, got)
		}
	})

	t.Run("returns nil when Annotations itself is nil", func(t *testing.T) {
		spec := &specs.Spec{Annotations: nil}
		got := CDIDevicesFromSpec(spec)
		if got != nil {
			t.Errorf("CDIDevicesFromSpec(%v) = %v, want nil", spec, got)
		}
	})
}
