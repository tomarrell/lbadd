package types

// Caster wraps the Cast method, which is used to transform the input value
// into an output value. Types can implement this interface. E.g. if the
// type String implements Caster, any value passed into the Cast method
// should be attempted to be cast to String, or an error should be returned.
type Caster interface {
	Cast(Value) (Value, error)
}
