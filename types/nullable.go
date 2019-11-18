package types

type Numeric struct {
	Valid bool
	Value int
}

func NewNumeric(value int) Numeric {
	return Numeric{
		Valid: true,
		Value: value,
	}
}

type AlphaNumeric struct {
	Valid bool
	Value string
}

func NewAlphaNumeric(value string) AlphaNumeric {
	return AlphaNumeric{
		Valid: true,
		Value: value,
	}
}

type Bool struct {
	Valid bool
	Value bool
}

func NewBool(value bool) Bool {
	return Bool{
		Valid: true,
		Value: value,
	}
}
