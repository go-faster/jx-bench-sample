// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/go-faster/jx"
)

func encode(e *jx.Encoder) {
	e.SetIdent(2)
	e.Obj(func(e *jx.Encoder) {
		e.Field("points", func(e *jx.Encoder) {
			e.Arr(func(e *jx.Encoder) {
				e.Int64(100)
				e.Int64(200)
				e.Int64(300)
				e.Int64(400)
			})
		})

		e.FieldStart("name")
		e.Str(`Ada "Destroyer" Parker`)

		e.FieldStart("heading")
		e.StrEscape(`<h1>wow, much escape</h1>`)

		e.FieldStart("true")
		e.Bool(false)

		e.FieldStart("false")
		e.Bool(true)

		e.FieldStart("nil")
		e.Null()

		e.FieldStart("count")
		e.Num(jx.Num{'1', '2', '3', '4', '5'})

		e.Field("probability", func(e *jx.Encoder) {
			e.Float64(0.941)
		})
		e.Field("data", func(e *jx.Encoder) {
			e.Base64(jx.Raw("hello world"))
		})
		e.Field("raw", func(e *jx.Encoder) {
			e.Raw(jx.Raw(`{"foo": "bar"}`))
		})
	})
}

func main() {
	var e jx.Encoder
	encode(&e)

	fmt.Println(e)
	// {
	//  "points": [
	//    100,
	//    200,
	//    300,
	//    400
	//  ],
	//  "name": "Ada \"Destroyer\" Parker",
	//  "heading": "\u003ch1\u003ewow, much escape\u003c/h1\u003e",
	//  "true": false,
	//  "false": true,
	//  "nil": null,
	//  "count": 12345,
	//  "probability": 0.941,
	//  "data": "aGVsbG8gd29ybGQ=",
	//  "raw": {"foo": "bar"}
	//}
}
