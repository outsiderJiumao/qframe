package types

// DataType represents any of the data types valid in a QFrame.
type DataType string

const (
	// None represents an unknown data type.
	// This is mainly used to indicate that the type of a column should be auto detected.
	None DataType = ""

	// Int translates into the Go int type. Missing values cannot be represented explicitly.
	Int = "int"

	// String translates into the Go *string type. nil represents a missing value.
	// Internally a string currently has an overhead of eight bytes (64 bits) in
	// addition to the bytes actually used to hold the string.
	String = "string"

	// Float translates into the Go float64 type. NaN represents a missing value.
	Float = "float"

	// Bool translates into the Go bool type. Missing values cannot be represented explicitly.
	Bool = "bool"

	// Enum translates into the Go *string type. nil represents a missing value.
	// An enum column can, at most, have 254 distinct values.
	Enum = "enum"

	// Datetime translates into Go time.Time type
	Datetime = "datetime"

	// Undefined represents an unspecified data type.
	// This is used for zero length columns where the datatype could not be identified.
	Undefined DataType = "Undefined"
)

// FunctionType represents the different types of input that functions operating on columns can take.
type FunctionType byte

const (
	// FunctionTypeUndefined bla
	FunctionTypeUndefined FunctionType = iota
	FunctionTypeInt
	FunctionTypeFloat
	FunctionTypeBool
	FunctionTypeString
	FunctionTypeDatetime
)

func (t FunctionType) String() string {
	switch t {
	case FunctionTypeInt:
		return "Int function"
	case FunctionTypeBool:
		return "Bool function"
	case FunctionTypeString:
		return "String function"
	case FunctionTypeFloat:
		return "Float function"
	case FunctionTypeUndefined:
		return "Undefined type function"
	case FunctionTypeDatetime:
		return "Datetime function"
	default:
		return "Unknown function"
	}
}
