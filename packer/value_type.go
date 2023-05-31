package packer

// ValueType the type for JSON element
type ValueType int

const (
	// InvalidValue invalid JSON element
	InvalidValue ValueType = iota
	// StringValue JSON element "string"
	StringValue
	// NumberValue JSON element 100 or 0.10
	NumberValue
	// NullableValue JSON element null
	NullableValue
	// BoolValue JSON element true or false
	BoolValue
	// ArrayValue JSON element []
	ArrayValue
	// ObjectValue JSON element {}
	ObjectValue
)
