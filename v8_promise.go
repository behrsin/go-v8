package v8

// #include "v8_c_bridge.h"
// #cgo CXXFLAGS: -I${SRCDIR} -I${SRCDIR}/include -g3 -fno-rtti -fpic -std=c++11
// #cgo LDFLAGS: -pthread -L${SRCDIR}/libv8 -lv8_base -lv8_init -lv8_initializers -lv8_libbase -lv8_libplatform -lv8_libsampler -lv8_nosnapshot
import "C"

import (
	"fmt"
	"runtime"
)

type Resolver struct {
	context *Context
	pointer C.ResolverPtr
}

func (c *Context) NewResolver() (*Resolver, error) {
	pr := C.v8_Promise_NewResolver(c.pointer)
	if pr == nil {
		return nil, fmt.Errorf("cannot create resolver for context")
	}
	r := &Resolver{
		context: c,
		pointer: pr,
	}
	runtime.SetFinalizer(r, (*Resolver).release)
	return r, nil
}

func (r *Resolver) Resolve(v *Value) error {
	err := C.v8_Resolver_Resolve(r.context.pointer, r.pointer, v.pointer)
	return r.context.isolate.newError(err)
}

func (r *Resolver) Reject(v *Value) error {
	err := C.v8_Resolver_Reject(r.context.pointer, r.pointer, v.pointer)
	return r.context.isolate.newError(err)
}

func (r *Resolver) Promise() *Value {
	pv := C.v8_Resolver_GetPromise(r.context.pointer, r.pointer)
	v := r.context.newValue(pv, unionKindPromise)
	v.created = true
	return v
}

func (r *Resolver) release() {
	if r.pointer != nil {
		C.v8_Resolver_Release(r.context.pointer, r.pointer)
	}
	r.context = nil
	r.pointer = nil
	runtime.SetFinalizer(r, nil)
}