package types

type (
	// ArithmeticAdder wraps the arithmetic add operation, usually represented
	// by a '+'. The actual addition is defined and must be documented by the
	// implementing type.
	ArithmeticAdder interface {
		Add(Value, Value) (Value, error)
	}

	// ArithmeticSubtractor wraps the arithmetic sub operation, usually
	// represented by a '-'. The actual subtraction is defined and must be
	// documented by the implementing type.
	ArithmeticSubtractor interface {
		Sub(Value, Value) (Value, error)
	}

	// ArithmeticMultiplicator wraps the arithmetic mul operation, usually
	// represented by a '*'. The actual multiplication is defined and must be
	// documented by the implementing type.
	ArithmeticMultiplicator interface {
		Mul(Value, Value) (Value, error)
	}

	// ArithmeticDivider wraps the arithmetic div operation, usually represented
	// by a '/'. The actual division is defined and must be documented by the
	// implementing type.
	ArithmeticDivider interface {
		Div(Value, Value) (Value, error)
	}

	// ArithmeticModulator wraps the arithmetic mod operation, usually
	// represented by a '%'. The actual modulation is defined and must be
	// documented by the implementing type.
	ArithmeticModulator interface {
		Mod(Value, Value) (Value, error)
	}

	// ArithmeticExponentiator wraps the arithmetic pow operation, usually
	// represented by a '**'. The actual exponentiation is defined and must be
	// documented by the implementing type.
	ArithmeticExponentiator interface {
		Pow(Value, Value) (Value, error)
	}
)
