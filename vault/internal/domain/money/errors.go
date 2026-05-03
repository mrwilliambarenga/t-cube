package money

import "errors"

var (
	ErrCurrencyMismatch = errors.New("currency mismatch")
	ErrInvalidCurrency = errors.New("invalid currency")
)

