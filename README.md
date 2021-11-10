# jx bench sample

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
$ go test -run ^$ -count 10 -bench BenchmarkEncode
goos: linux
goarch: amd64
pkg: jx-bench
cpu: AMD Ryzen 9 5950X 16-Core Processor            
BenchmarkEncode-32       2768792               429.6 ns/op       703.03 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32       2794426               425.6 ns/op       709.53 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32       2575532               425.1 ns/op       710.35 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32       2719108               423.5 ns/op       713.12 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32       2798019               424.3 ns/op       711.73 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32       2804731               432.5 ns/op       698.27 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32       2806100               417.4 ns/op       723.58 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32       2774716               424.1 ns/op       712.04 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32       2791240               424.4 ns/op       711.63 MB/s           0 B/op          0 allocs/op
BenchmarkEncode-32       2761339               424.4 ns/op       711.57 MB/s           0 B/op          0 allocs/op
PASS
ok      jx-bench        16.191s
```