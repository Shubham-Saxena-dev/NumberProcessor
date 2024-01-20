package models

type NumberType string

const (
	Primes NumberType = "primes"
	Fibo   NumberType = "fibo"
	Odd    NumberType = "odd"
	Rand   NumberType = "rand"
	Err    NumberType = "err"
)

func (c NumberType) ValidateTypes() bool {
	switch c {
	case Primes, Fibo, Odd, Rand:
		return true
	}
	return false
}
