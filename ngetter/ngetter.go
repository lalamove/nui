package ngetter

import (
	"fmt"
	"sync/atomic"

	"github.com/spf13/cast"
)

// Provider is an interface to provide a value in a thread safe manner
// And be able to replace this value in a thread safe manner.
type Provider interface {
	Getter
	Replace(interface{})
}

type provider struct {
	v *atomic.Value
}

func (p *provider) Get() interface{} {
	return p.v.Load()
}

func (p *provider) Replace(v interface{}) {
	p.v.Store(v)
}

// NewProvider returns a new Provider from the given value x
func NewProvider(x interface{}) Provider {
	var v atomic.Value
	v.Store(x)
	return &provider{
		v: &v,
	}
}

// Getter is a generic interface to get a value
type Getter interface {
	Get() interface{}
}

// Inter is an interface to get an int value
type Inter interface {
	Int() int
}

// Inter64 is an interface to get an int64 value
type Inter64 interface {
	Int64() int64
}

// Floater64 is an interface to get a Float64 value
type Floater64 interface {
	Float64() float64
}

// Booler is an interface to get a bool value
type Booler interface {
	Bool() bool
}

// Stringer is an alias to fmt.Stringer
type Stringer fmt.Stringer

// GetterTyped is a getter with methods to cast
type GetterTyped interface {
	Getter
	Inter
	Inter64
	Floater64
	Booler
	Stringer
}

var _ GetterTyped = (GetterTypedFunc)(nil)

// GetterTypedFunc is a function implementing the GetterTyped interface
type GetterTypedFunc func() interface{}

// Get returns the interface{} value
func (f GetterTypedFunc) Get() interface{} {
	return f()
}

// String returns the string value
func (f GetterTypedFunc) String() string {
	return cast.ToString(f())
}

// Int returns the int value
func (f GetterTypedFunc) Int() int {
	return cast.ToInt(f())
}

// Int64 returns the int64 value
func (f GetterTypedFunc) Int64() int64 {
	return cast.ToInt64(f())
}

// Float64 returns the float64 value
func (f GetterTypedFunc) Float64() float64 {
	return cast.ToFloat64(f())
}

// Bool returns the bool value
func (f GetterTypedFunc) Bool() bool {
	return cast.ToBool(f())
}
