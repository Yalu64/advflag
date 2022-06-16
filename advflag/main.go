package advflag

import (
	"log"
	"os"
	"sync"
)

var once sync.Once

var Config *Flag
var advFlag *internalFlag
var mutex sync.Mutex

func init() {
	Config = &Flag{
		ArgSym:              defaultArgSym,
		ArgBoolSym:          defaultBoolSym,
		ArgMaxCharLen:       defaultArgMaxCharLen,
		ArgPrefixMaxCharLen: defaultArgPrefixMaxCharLen,
	}
	logger := log.New(os.Stdout, "advflag: ", log.Ldate)
	advFlag = &internalFlag{
		logger:        logger,
		args:          make(map[string]*arg),
		localArgs:     make(map[string]arg),
		localArgIndex: make(map[string]int),
	}
}
