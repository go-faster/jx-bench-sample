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
go test -run ^$ -count 10 -bench BenchmarkEncode
goos: linux
goarch: amd64
pkg: fx-bench
cpu: AMD Ryzen 9 5950X 16-Core Processor            
BenchmarkEncode-32      23826493                44.53 ns/op      741.05 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32      23727690                44.09 ns/op      748.42 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32      23706267                44.40 ns/op      743.28 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32      23775981                43.80 ns/op      753.46 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32      22942275                44.14 ns/op      747.61 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32      24156576                44.06 ns/op      748.94 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32      24324625                42.84 ns/op      770.34 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32      26891583                43.94 ns/op      751.11 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32      25446589                43.28 ns/op      762.54 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32      24522786                44.14 ns/op      747.62 MB/s           0 B/op          0 allocs/op
PASS
ok      fx-bench        12.201s
```