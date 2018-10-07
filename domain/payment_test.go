package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaymentMarshalling(t *testing.T) {
	payment := CreateTestPayment()
	assert := assert.New(t)
	var (
		data []byte
		err  error
	)
	t.Run("Should marshal without error", func(t *testing.T) {
		data, err = PaymentToByteSlice(payment)
		assert.Nil(err)
		assert.NotEmpty(data)
	})
	t.Run("Should unmarshal without error", func(t *testing.T) {
		retPayment, err := PaymentFromByteSlice(data)
		assert.Nil(err)
		assert.NotEmpty(retPayment)
	})
}
