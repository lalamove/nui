package nmetrics

import (
	"context"
	"errors"

	"github.com/prometheus/client_golang/prometheus"
)

type registryKey string

// RegistryKey is the key to access the registry in context
const RegistryKey registryKey = "prometheus.Registry"

// ErrRegistryNotFoundInContext is an error indicating registry not found in context
var ErrRegistryNotFoundInContext = errors.New("Registry not found in context")

// SetRegistryInContext returns a new context with a new registry set
func SetRegistryInContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, RegistryKey, prometheus.NewRegistry())
}

// RegistryFromContext return a registry stored in the context
// if the registry is not found, nil will be sent
func RegistryFromContext(ctx context.Context) *prometheus.Registry {
	r := ctx.Value(RegistryKey)
	if r == nil {
		return nil
	}
	return r.(*prometheus.Registry)
}

// MustRegistryFromContext return a registry stored in the context
// if the registry is not found, panic with ErrRegistryNotFoundInContext will be sent
func MustRegistryFromContext(ctx context.Context) *prometheus.Registry {
	r := ctx.Value(RegistryKey)
	if r == nil {
		panic(ErrRegistryNotFoundInContext)
	}
	return r.(*prometheus.Registry)
}
