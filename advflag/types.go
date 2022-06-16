package advflag

import (
	"reflect"
)

func Int(prefix string, def int, usage interface{}) *int {
	a := &arg{
		dataType: reflect.ValueOf(def),
		value:    &def,
		usage:    usage,
	}
	mutex.Lock()
	advFlag.args[prefix] = a
	mutex.Unlock()
	return a.value.(*int)
}

func Int8(prefix string, def int8, usage interface{}) *int8 {
	a := &arg{
		dataType: reflect.ValueOf(def),
		value:    &def,
		usage:    usage,
	}
	mutex.Lock()
	advFlag.args[prefix] = a
	mutex.Unlock()
	return a.value.(*int8)
}

func Int16(prefix string, def int16, usage interface{}) *int16 {
	a := &arg{
		dataType: reflect.ValueOf(def),
		value:    &def,
		usage:    usage,
	}
	mutex.Lock()
	advFlag.args[prefix] = a
	mutex.Unlock()
	return a.value.(*int16)
}

func Int32(prefix string, def int32, usage interface{}) *int32 {
	a := &arg{
		dataType: reflect.ValueOf(def),
		value:    &def,
		usage:    usage,
	}
	mutex.Lock()
	advFlag.args[prefix] = a
	mutex.Unlock()
	return a.value.(*int32)
}

func Int64(prefix string, def int64, usage interface{}) *int64 {
	a := &arg{
		dataType: reflect.ValueOf(def),
		value:    &def,
		usage:    usage,
	}
	mutex.Lock()
	advFlag.args[prefix] = a
	mutex.Unlock()
	return a.value.(*int64)
}

func UInt(prefix string, def uint, usage interface{}) *uint {
	a := &arg{
		dataType: reflect.ValueOf(def),
		value:    &def,
		usage:    usage,
	}
	mutex.Lock()
	advFlag.args[prefix] = a
	mutex.Unlock()
	return a.value.(*uint)
}

func UInt8(prefix string, def uint8, usage interface{}) *uint8 {
	a := &arg{
		dataType: reflect.ValueOf(def),
		value:    &def,
		usage:    usage,
	}
	mutex.Lock()
	advFlag.args[prefix] = a
	mutex.Unlock()
	return a.value.(*uint8)
}

func UInt16(prefix string, def uint16, usage interface{}) *uint16 {
	a := &arg{
		dataType: reflect.ValueOf(def),
		value:    &def,
		usage:    usage,
	}
	mutex.Lock()
	advFlag.args[prefix] = a
	mutex.Unlock()
	return a.value.(*uint16)
}

func UInt32(prefix string, def uint32, usage interface{}) *uint32 {
	a := &arg{
		dataType: reflect.ValueOf(def),
		value:    &def,
		usage:    usage,
	}
	mutex.Lock()
	advFlag.args[prefix] = a
	mutex.Unlock()
	return a.value.(*uint32)
}

func UInt64(prefix string, def uint64, usage interface{}) *uint64 {
	a := &arg{
		dataType: reflect.ValueOf(def),
		value:    &def,
		usage:    usage,
	}
	mutex.Lock()
	advFlag.args[prefix] = a
	mutex.Unlock()
	return a.value.(*uint64)
}

func Float32(prefix string, def float32, usage interface{}) *float32 {
	a := &arg{
		dataType: reflect.ValueOf(def),
		value:    &def,
		usage:    usage,
	}
	mutex.Lock()
	advFlag.args[prefix] = a
	mutex.Unlock()
	return a.value.(*float32)
}

func Float64(prefix string, def float64, usage interface{}) *float64 {
	a := &arg{
		dataType: reflect.ValueOf(def),
		value:    &def,
		usage:    usage,
	}
	mutex.Lock()
	advFlag.args[prefix] = a
	mutex.Unlock()
	return a.value.(*float64)
}

func Complex64(prefix string, def complex64, usage interface{}) *complex64 {
	a := &arg{
		dataType: reflect.ValueOf(def),
		value:    &def,
		usage:    usage,
	}
	mutex.Lock()
	advFlag.args[prefix] = a
	mutex.Unlock()
	return a.value.(*complex64)
}

func Complex128(prefix string, def complex128, usage interface{}) *complex128 {
	a := &arg{
		dataType: reflect.ValueOf(def),
		value:    &def,
		usage:    usage,
	}
	mutex.Lock()
	advFlag.args[prefix] = a
	mutex.Unlock()
	return a.value.(*complex128)
}

func Bool(prefix string, def bool, usage interface{}) *bool {
	a := &arg{
		dataType: reflect.ValueOf(def),
		value:    &def,
		usage:    usage,
	}
	mutex.Lock()
	advFlag.args[prefix] = a
	mutex.Unlock()
	return a.value.(*bool)
}

func String(prefix string, def string, usage interface{}) *string {
	a := &arg{
		dataType: reflect.ValueOf(def),
		value:    &def,
		usage:    usage,
	}
	mutex.Lock()
	advFlag.args[prefix] = a
	mutex.Unlock()
	return a.value.(*string)
}
