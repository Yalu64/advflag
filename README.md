# advflag
A fast and advanced and simple command line argument parser written in pure go and up to 4x faster than build in flag package.

# Benchmark
*GOMAXPROCS=1*
- go test bench_test.go -bench=. -cpu=1 -benchmem
```
goos: linux
goarch: amd64
cpu: AMD Ryzen 7 5800X 8-Core Processor             
BenchmarkAdvFlag        21550642                55.83 ns/op            0 B/op          0 allocs/op
BenchmarkGoFlag          6084282               201.7 ns/op             0 B/op          0 allocs/op
PASS
ok      command-line-arguments  2.690s
```

# Examples
*Simple String and Bool argument parsed*
- go run simple_arguments.go -string hello_world -int 727 --bool
```
hello_world
727
true
```
*Complex Type parsed*
- go run complex_arguments.go -number 10+0i
```
(10+0i)
```


# Syntax
* Message shows up when -help, ? is typed as first argument. Displays all added arguments with type, default value and usage
```
Syntax of AppName
        string (usage: <nil>)
                -string <string> [default: this_is_an_test_value]
        int (usage: Scary number also known as 2^8)
                -int <int> [default: 256]
        uint (usage: hehe)
                -uint <uint64> [default: 512]
        bool (usage: <nil>)
                --bool <bool> [default: false]

```                
