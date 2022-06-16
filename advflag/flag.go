package advflag

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type arg struct {
	dataType reflect.Value
	usage    interface{}
	value    interface{}
}

type Flag struct {
	ArgSym              string
	ArgBoolSym          string
	ArgMaxCharLen       int
	ArgPrefixMaxCharLen int
}

type internalFlag struct {
	logger *log.Logger
	// arguments get added here and will be modified with localArgs
	args map[string]*arg
	// [Parsing] index where an argument (defaultArgSym) starts
	localArgIndex map[string]int
	// [Parsing] all arguments
	localArgs map[string]arg
}

func parseVal(val string, arg *arg) {
	switch arg.dataType.Interface().(type) {
	case int:
		{
			res, err := strconv.Atoi(val)
			if err != nil {
				advFlag.logger.Printf("Could not parse '%s' to int\n", val)
				break
			}
			*arg.value.(*int) = res
		}
	case int8:
		{
			res, err := strconv.ParseInt(val, 10, 8)
			if err != nil {
				advFlag.logger.Printf("Could not parse '%s' to int8\n", val)
				break
			}
			*arg.value.(*int8) = int8(res)
		}
	case int16:
		{
			res, err := strconv.ParseInt(val, 10, 16)
			if err != nil {
				advFlag.logger.Printf("Could not parse '%s' to int16\n", val)
				break
			}
			*arg.value.(*int16) = int16(res)
		}
	case int32:
		{
			res, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				advFlag.logger.Printf("Could not parse '%s' to int32\n", val)
				break
			}
			*arg.value.(*int32) = int32(res)
		}
	case int64:
		{
			res, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				advFlag.logger.Printf("Could not parse '%s' to int64\n", val)
				break
			}
			*arg.value.(*int64) = res
		}
	case uint:
		{
			res, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				advFlag.logger.Printf("Could not parse '%s' to uint\n", val)
				break
			}
			*arg.value.(*uint) = uint(res)
		}
	case uint8:
		{
			res, err := strconv.ParseUint(val, 10, 8)
			if err != nil {
				advFlag.logger.Printf("Could not parse '%s' to uint8\n", val)
				break
			}
			*arg.value.(*uint8) = uint8(res)
		}
	case uint32:
		{
			res, err := strconv.ParseUint(val, 10, 32)
			if err != nil {
				advFlag.logger.Printf("Could not parse '%s' to uint32\n", val)
				break
			}
			*arg.value.(*uint32) = uint32(res)
		}
	case uint64:
		{
			res, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				advFlag.logger.Printf("Could not parse '%s' to uint64\n", val)
				break
			}
			*arg.value.(*uint64) = res
		}
	case float32:
		{
			res, err := strconv.ParseFloat(val, 32)
			if err != nil {
				advFlag.logger.Printf("Could not parse '%s' to float32\n", val)
				break
			}
			*arg.value.(*float32) = float32(res)
		}
	case float64:
		{
			res, err := strconv.ParseFloat(val, 64)
			if err != nil {
				advFlag.logger.Printf("Could not parse '%s' to float64\n", val)
				break
			}
			*arg.value.(*float64) = res
		}
	case complex64:
		{
			res, err := strconv.ParseComplex(val, 64)
			if err != nil {
				advFlag.logger.Printf("Could not parse '%s' to complex64\n", val)
				break
			}
			*arg.value.(*complex64) = complex64(res)
		}
	case complex128:
		{
			res, err := strconv.ParseComplex(val, 128)
			if err != nil {
				advFlag.logger.Printf("Could not parse '%s' to complex64\n", val)
				break
			}
			*arg.value.(*complex128) = res
		}

	case string:
		{
			*arg.value.(*string) = val
		}
	default:
		break
	}
}

func isBool(r reflect.Value) bool {
	switch r.Interface().(type) {
	case bool:
		{
			return true
		}
	}
	return false
}

func Parse() {
	l := Len()
	if l == 0 {
		return
	}
	args := Args()

	if PrintSyntax {
		if l == 1 {
			if strings.ToLower(args[0]) == "help" || strings.ToLower(args[0]) == Config.ArgSym+"help" || args[0] == "?" {
				PrintHelp()
			}
		}
	}

	for i, field := range args {
		if len(field) > 2 && field[:2] == Config.ArgBoolSym {
			ptr := advFlag.args[field[2:]]
			if ptr != nil && isBool(ptr.dataType) {
				*advFlag.args[field[2:]].value.(*bool) = true
			}
			continue
		} else if len(field) > 1 && field[:1] == Config.ArgSym {
			ptr := advFlag.args[field[1:]]
			if ptr != nil {
				advFlag.localArgIndex[field[1:]] = i
			}
		}
	}

	for a, i := range advFlag.localArgIndex {
		if i >= len(args)-1 {
			continue
		} else if args[i+1][:1] == Config.ArgSym {
			continue
		}

		ptr := advFlag.args[a]
		if ptr != nil {
			parseVal(args[i+1], ptr)
		}
	}
}

func PrintHelp() {
	fmt.Println("Syntax of", AppName)
	var boolArgs string
	for name, argument := range advFlag.args {
		if isBool(argument.dataType) {
			boolArgs += fmt.Sprintf("\t%s (usage: %v)\n\t\t%s%s <%s> [default: %v]\n", name, argument.usage, Config.ArgBoolSym, name, argument.dataType.Type(), argument.dataType.Interface())
			continue
		}
		fmt.Printf("\t%s (usage: %v)\n\t\t%s%s <%s> [default: %v]\n", name, argument.usage, Config.ArgSym, name, argument.dataType.Type(), argument.dataType.Interface())
	}
	fmt.Print(boolArgs)
}
