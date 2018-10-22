package test

import (
	"io/ioutil"
	"testing"

	"github.com/mysza/paymentsapi/domain"
)

func loadBytes(t *testing.T, path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}

func PaymentFromFile(t *testing.T, path string) *domain.Payment {
	bytes := loadBytes(t, path)
	payment, err := domain.PaymentFromByteSlice(bytes)
	if err != nil {
		t.Fatal(err)
	}
	return payment
}
