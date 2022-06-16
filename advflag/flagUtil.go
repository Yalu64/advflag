package advflag

import (
	"os"
)

func Len() int {
	return len(os.Args) - 1
}

func Args() []string {
	return os.Args[1:]
}

func RawLen() int {
	return len(os.Args)
}

func FilePath() string {
	return os.Args[0]
}
