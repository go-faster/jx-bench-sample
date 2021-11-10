// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/go-faster/jx"
)

func main() {
	var e jx.Encoder
	e.ArrStart()
	e.Int64(100)
	e.Str("hello")
	e.Int64(200)
	e.Int64(300)
	e.Str("world")
	e.Int64(400)
	e.ArrEnd()

	fmt.Println(e)
	// [100,200,300,400,"hello"]
}
