package main

import (
	"fmt"
	"github.com/Yariya/advflag/advflag"
)

var c128 = advflag.Complex128("number", 2+0i, nil)

func main() {
	advflag.Parse()
	fmt.Println(*c128)
}
