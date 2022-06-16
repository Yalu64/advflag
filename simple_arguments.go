package main

import (
	"fmt"
	"github.com/Yariya/advflag/advflag"
)

var argAdvFlagString = advflag.String("string", "this_is_an_test_value", nil)
var argAdvFlagInt = advflag.Int("int", 256, "Scary number also known as 2^8")
var argAdvFlagIntExpr = advflag.UInt64("uint", 512, "hehe")
var argAdvFlagBool = advflag.Bool("bool", false, nil)

func main() {
	advflag.Parse()
	fmt.Println(*argAdvFlagString)
	fmt.Println(*argAdvFlagInt)
	fmt.Println(*argAdvFlagIntExpr)
	fmt.Println(*argAdvFlagBool)
}
