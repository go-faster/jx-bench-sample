# fx bench sample

```go
package main

import "github.com/go-faster/jx"

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
}
```
Should encode:
```json
[100,"hello",200,300,"world",400]
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