package logger

// LogObj object that holds data for each field inserted to each log message.
// How Log implementer is treating this object should read the field typ
// and follow the guideline from LogObjType and each of the supported types.
type LogObj struct {
	typ LogObjType
	key string
	str string
	num int
	flt float64
	b   bool
	any interface{}
	err error
}

// LogObjType indicates how the Log implementer should treat each
// LogObj.
type LogObjType uint8

const (
	// StringType use field str string of LogObj as the value.
	StringType LogObjType = iota
	// NumType use field num int of LogObj as the value.
	NumType
	// FloatType use field flt float64 of LogObj as the value.
	FloatType
	// BoolType use field b bool of LogObj as the value.
	BoolType
	// AnyType use field any interface of LogObj as the value.
	AnyType
	// ErrorType use field err from error interface of LogObj as the value.
	ErrorType
)

// String constructs a LogObj with the given key and value. This set the type
// to StringType.
func String(k, v string) LogObj {
	return LogObj{typ: StringType, key: k, str: v}
}

// Num constructs a LogObj with the given key and value. This set the type
// to NumType.
func Num(k string, num int) LogObj {
	return LogObj{typ: NumType, key: k, num: num}
}

// Float constructs a LogObj with the given key and value. This set the type
// to FloatType.
func Float(k string, f float64) LogObj {
	return LogObj{typ: FloatType, key: k, flt: f}
}

// Bool constructs a LogObj with the given key and value. This set the type
// to BoolType.
func Bool(k string, b bool) LogObj {
	return LogObj{typ: BoolType, key: k, b: b}
}

// Any constructs a LogObj with the given key and value. This set the type
// to AnyType.
func Any(k string, any interface{}) LogObj {
	return LogObj{typ: AnyType, key: k, any: any}
}

// Error constructs a LogObj with the given err and 'error' as the key. This set
// the type to ErrorType.
func Error(err error) LogObj {
	return LogObj{typ: ErrorType, key: "error", err: err}
}
