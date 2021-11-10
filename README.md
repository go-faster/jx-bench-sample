# fx bench sample

```go
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
}
```
Should encode:
```json
{
  "points": [
    100,
    200,
    300,
    400
  ],
  "name": "Ada \"Destroyer\" Parker",
  "heading": "\u003ch1\u003ewow, much escape\u003c/h1\u003e",
  "true": false,
  "false": true,
  "nil": null,
  "count": 12345,
  "probability": 0.941,
  "data": "aGVsbG8gd29ybGQ=",
  "raw": {"foo": "bar"}
}
```

```console
goos: linux
goarch: amd64
pkg: fx-bench
cpu: AMD Ryzen 9 5950X 16-Core Processor            
BenchmarkEncode-32  2781439  429.2 ns/op  703.64 MB/s  0 B/op  0 allocs/op
BenchmarkEncode-32  2773333  430.6 ns/op  701.39 MB/s  0 B/op  0 allocs/op
BenchmarkEncode-32  2722332  428.6 ns/op  704.70 MB/s  0 B/op  0 allocs/op
BenchmarkEncode-32  2773780  435.3 ns/op  693.74 MB/s  0 B/op  0 allocs/op
BenchmarkEncode-32  2759192  428.1 ns/op  705.39 MB/s  0 B/op  0 allocs/op
BenchmarkEncode-32  2797752  434.3 ns/op  695.35 MB/s  0 B/op  0 allocs/op
BenchmarkEncode-32  2813332  429.2 ns/op  703.71 MB/s  0 B/op  0 allocs/op
BenchmarkEncode-32  2801134  434.8 ns/op  694.60 MB/s  0 B/op  0 allocs/op
BenchmarkEncode-32  2736691  428.2 ns/op  705.35 MB/s  0 B/op  0 allocs/op
BenchmarkEncode-32  2766651  428.7 ns/op  704.38 MB/s  0 B/op  0 allocs/op
PASS
ok      fx-bench        16.386s
```