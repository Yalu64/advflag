package main

import (
	"flag"
	"github.com/Yariya/advflag/advflag"
	"testing"
)

var argAdvFlag1 = advflag.String("arg1", "this_is_an_test_value", nil)
var argAdvFlag2 = advflag.String("arg2", "this_is_an_test_value", nil)
var argAdvFlag3 = advflag.String("arg3", "this_is_an_test_value", nil)

var argGoFlag1 = flag.String("arg1", "this_is_an_test_value", "")
var argGoFlag2 = flag.String("arg2", "this_is_an_test_value", "")
var argGoFlag3 = flag.String("arg3", "this_is_an_test_value", "")

func BenchmarkAdvFlag(b *testing.B) {
	for x := 0; x < b.N; x++ {
		advflag.Parse()
		_ = *argAdvFlag1
		_ = *argAdvFlag2
		_ = *argAdvFlag3
	}
}

func BenchmarkGoFlag(b *testing.B) {
	for x := 0; x < b.N; x++ {
		flag.Parse()
		_ = *argGoFlag1
		_ = *argGoFlag2
		_ = *argGoFlag3
	}
}
