package money

import (
	"errors"
	"testing"

	"github.com/shopspring/decimal"
)

func TestNewMoney(t *testing.T) {
	val, err := New(decimal.NewFromFloat(5.99), USD)

	if err != nil {
		t.Errorf("Expected no errors when created a new Money object. Got: %s", err)
	}

	if !val.amount.Equal(decimal.New(599, -2)) {
		t.Errorf("Expected amount to equal 5.99. Got: %s", val.amount)
	}

	if val.currency != USD {
		t.Errorf("Expected currency to be USD. Got: %s", val.currency)
	}
}

func TestNewMoney_InvalidCurrency(t *testing.T) {
	_, err := New(decimal.NewFromFloat(5.99), "")

	if err == nil {
		t.Errorf("Expected an error to be raised when creating a new Money object with empty currency")
	}

	if !errors.Is(err, ErrInvalidCurrency) {
	  t.Errorf("Expected an ErrInvalidCurrency error when creating a new Money object with empty currency")
	}
}
