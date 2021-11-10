// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/go-faster/jx"
)

func BenchmarkEncode(b *testing.B) {
	b.ReportAllocs()
	var e jx.Encoder
	encode(&e)
	b.SetBytes(int64(len(e.Bytes())))

	for i := 0; i < b.N; i++ {
		e.Reset()
		encode(&e)
	}
}
