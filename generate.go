// Copyright (C) 2017 ScyllaDB
// Use of this source code is governed by a ALv2-style
// license that can be found in the LICENSE file.

package set

import (
	"github.com/scylladb/go-set/f32set"
	"github.com/scylladb/go-set/f64set"
	"github.com/scylladb/go-set/i16set"
	"github.com/scylladb/go-set/i32set"
	"github.com/scylladb/go-set/i64set"
	"github.com/scylladb/go-set/i8set"
	"github.com/scylladb/go-set/iset"
	"github.com/scylladb/go-set/sset"
	"github.com/scylladb/go-set/u16set"
	"github.com/scylladb/go-set/u32set"
	"github.com/scylladb/go-set/u64set"
	"github.com/scylladb/go-set/u8set"
	"github.com/scylladb/go-set/uset"
)

//go:generate go_generics -i set.tpl -t T=float32 -o f32set/f32set.go -p f32set
//go:generate go_generics -i set_test.tpl -t T=float32 -o f32set/f32set_test.go -p f32set

// NewFloat32Set is a convenience function to create a new f32set.Set
func NewFloat32Set() *f32set.Set {
	return f32set.New()
}

//go:generate go_generics -i set.tpl -t T=float64 -o f64set/f64set.go -p f64set
//go:generate go_generics -i set_test.tpl -t T=float64 -o f64set/f64set_test.go -p f64set

// NewFloat64Set is a convenience function to create a new f64set.Set
func NewFloat64Set() *f64set.Set {
	return f64set.New()
}

//go:generate go_generics -i set.tpl -t T=int -o iset/iset.go -p iset
//go:generate go_generics -i set_test.tpl -t T=int -o iset/iset_test.go -p iset

// NewIntSet is a convenience function to create a new iset.Set
func NewIntSet() *iset.Set {
	return iset.New()
}

//go:generate go_generics -i set.tpl -t T=int8 -o i8set/i8set.go -p i8set
//go:generate go_generics -i set_test.tpl -t T=int8 -o i8set/i8set_test.go -p i8set

// NewInt8Set is a convenience function to create a new i8set.Set
func NewInt8Set() *i8set.Set {
	return i8set.New()
}

//go:generate go_generics -i set.tpl -t T=int16 -o i16set/i16set.go -p i16set
//go:generate go_generics -i set_test.tpl -t T=int16 -o i16set/i16set_test.go -p i16set

// NewInt16Set is a convenience function to create a new i16set.Set
func NewInt16Set() *i16set.Set {
	return i16set.New()
}

//go:generate go_generics -i set.tpl -t T=int32 -o i32set/i32set.go -p i32set
//go:generate go_generics -i set_test.tpl -t T=int32 -o i32set/i32set_test.go -p i32set

// NewInt32Set is a convenience function to create a new i32set.Set
func NewInt32Set() *i32set.Set {
	return i32set.New()
}

//go:generate go_generics -i set.tpl -t T=int64 -o i64set/i64set.go -p i64set
//go:generate go_generics -i set_test.tpl -t T=int64 -o i64set/i64set_test.go -p i64set

// NewInt64Set is a convenience function to create a new i64set.Set
func NewInt64Set() *i64set.Set {
	return i64set.New()
}

//go:generate go_generics -i set.tpl -t T=uint -o uset/uset.go -p uset
//go:generate go_generics -i set_test.tpl -t T=uint -o uset/uset_test.go -p uset

// NewUintSet is a convenience function to create a new uset.Set
func NewUintSet() *uset.Set {
	return uset.New()
}

//go:generate go_generics -i set.tpl -t T=uint8 -o u8set/set.go -p u8set
//go:generate go_generics -i set_test.tpl -t T=uint8 -o u8set/set_test.go -p u8set

// NewUint8Set is a convenience function to create a new u8set.Set
func NewUint8Set() *u8set.Set {
	return u8set.New()
}

//go:generate go_generics -i set.tpl -t T=uint16 -o u16set/u16set.go -p u16set
//go:generate go_generics -i set_test.tpl -t T=uint16 -o u16set/u16set_test.go -p u16set

// NewUint16Set is a convenience function to create a new u16set.Set
func NewUint16Set() *u16set.Set {
	return u16set.New()
}

//go:generate go_generics -i set.tpl -t T=uint32 -o u32set/u32set.go -p u32set
//go:generate go_generics -i set_test.tpl -t T=uint32 -o u32set/u32set_test.go -p u32set

// NewUint32Set is a convenience function to create a new u32set.Set
func NewUint32Set() *u32set.Set {
	return u32set.New()
}

//go:generate go_generics -i set.tpl -t T=uint64 -o u64set/u64set.go -p u64set
//go:generate go_generics -i set_test.tpl -t T=uint64 -o u64set/u64set_test.go -p u64set

// NewUint64Set is a convenience function to create a new u64set.Set
func NewUint64Set() *u64set.Set {
	return u64set.New()
}

//go:generate go_generics -i set.tpl -t T=string -o sset/sset.go -p sset
//go:generate go_generics -i set_test.tpl -t T=string -o sset/sset_test.go -p sset

// NewStringSet is a convenience function to create a new sset.Set
func NewStringSet() *sset.Set {
	return sset.New()
}
