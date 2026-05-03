package money

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Currency string

const (
	USD Currency = "USD"
	GBP Currency = "GBP"
	EUR Currency = "EUR"
)

type Money struct {
	amount   decimal.Decimal
	currency Currency
}

func New(amount decimal.Decimal, currency Currency) (Money, error) {
	if currency == "" {
		return Money{}, fmt.Errorf("%w: %s", ErrInvalidCurrency, currency)
	}

	return Money{
		amount: amount,
		currency: currency,
	}, nil
}

func (m Money) Amount() decimal.Decimal {
	return m.amount
}

func (m Money) Currency() Currency {
	return m.currency
}

func (m Money) Add(other Money) (Money, error) {
	if other.currency != m.currency {
		return Money{}, fmt.Errorf("%w: %s vs %s", ErrCurrencyMismatch, m.currency, other.currency)
	}

	return Money{
		amount: m.amount.Add(other.amount),
		currency: m.currency,
	}, nil
}
