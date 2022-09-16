go1.19.1

```
make bench
GOAMD64=v3 go test -bench=. -benchmem 
2022/09/16 12:25:01 in TestMain
goos: darwin
goarch: amd64
pkg: btbench
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkProto-8                   53479             21866 ns/op           73728 B/op          1 allocs/op
BenchmarkProtoUnmarshal-8          31798             37469 ns/op           85304 B/op        309 allocs/op
BenchmarkJSON-8                    10000            101023 ns/op           73789 B/op          1 allocs/op
BenchmarkJSONUnmarshal-8            2622            427579 ns/op           87240 B/op        328 allocs/op
BenchmarkION-8                      4335            254439 ns/op          359440 B/op       2584 allocs/op
BenchmarkIONUnmarshal-8             4071            273780 ns/op          264144 B/op       2969 allocs/op
BenchmarkFFJSON-8                   9409            110289 ns/op          208083 B/op        209 allocs/op
BenchmarkFFJSONUnmarshal-8         14400             83956 ns/op           81831 B/op        315 allocs/op
BenchmarkEasyJSON-8                 2607            443117 ns/op          149835 B/op         14 allocs/op
BenchmarkEasyJSONUnmarshal-8        2310            492302 ns/op           80608 B/op        310 allocs/op
PASS
ok      btbench 13.475s
```


go 1.18.6
```
make bench
GOAMD64=v3 go test -bench=. -benchmem 
2022/09/16 14:22:59 in TestMain
goos: darwin
goarch: amd64
pkg: btbench
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkProto-8                   47460             24877 ns/op           73728 B/op          1 allocs/op
BenchmarkProtoUnmarshal-8          30216             39038 ns/op           85304 B/op        309 allocs/op
BenchmarkJSON-8                    13466             88774 ns/op           73749 B/op          1 allocs/op
BenchmarkJSONUnmarshal-8            2760            411247 ns/op           87240 B/op        328 allocs/op
BenchmarkION-8                      3834            266359 ns/op          378177 B/op       2583 allocs/op
BenchmarkIONUnmarshal-8             4144            279352 ns/op          264144 B/op       2969 allocs/op
BenchmarkFFJSON-8                   8938            112258 ns/op          208126 B/op        209 allocs/op
BenchmarkFFJSONUnmarshal-8         14278             84256 ns/op           81828 B/op        315 allocs/op
BenchmarkEasyJSON-8                 2996            391198 ns/op          150117 B/op         14 allocs/op
BenchmarkEasyJSONUnmarshal-8        2265            505009 ns/op           80608 B/op        310 allocs/op
PASS
ok      btbench 14.315s
```