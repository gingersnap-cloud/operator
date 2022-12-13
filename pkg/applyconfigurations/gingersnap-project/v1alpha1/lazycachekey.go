// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gingersnap-project/operator/api/v1alpha1"
)

// LazyCacheKeyApplyConfiguration represents an declarative configuration of the LazyCacheKey type for use
// with apply.
type LazyCacheKeyApplyConfiguration struct {
	Format       *v1alpha1.KeyFormat `json:"format,omitempty"`
	KeySeparator *string             `json:"keySeparator,omitempty"`
}

// LazyCacheKeyApplyConfiguration constructs an declarative configuration of the LazyCacheKey type for use with
// apply.
func LazyCacheKey() *LazyCacheKeyApplyConfiguration {
	return &LazyCacheKeyApplyConfiguration{}
}

// WithFormat sets the Format field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Format field is set to the value of the last call.
func (b *LazyCacheKeyApplyConfiguration) WithFormat(value v1alpha1.KeyFormat) *LazyCacheKeyApplyConfiguration {
	b.Format = &value
	return b
}

// WithKeySeparator sets the KeySeparator field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeySeparator field is set to the value of the last call.
func (b *LazyCacheKeyApplyConfiguration) WithKeySeparator(value string) *LazyCacheKeyApplyConfiguration {
	b.KeySeparator = &value
	return b
}
