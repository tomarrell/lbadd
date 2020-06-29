package types

// Serializer wraps two basic methods for two-way serializing values. Which
// values can be serialized and deserialized, is up to the implementing object.
// It must be documented, what can and can not be serialized and deserialized.
type Serializer interface {
	Serialize(Value) ([]byte, error)
	Deserialize([]byte) (Value, error)
}
