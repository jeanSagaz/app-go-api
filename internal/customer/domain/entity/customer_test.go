package entity_test

import (
	"testing"

	"github.com/jeanSagaz/go-api/internal/customer/domain/entity"
	"github.com/stretchr/testify/require"
)

func TestValidateIfNameIsEmpty(t *testing.T) {
	customer, errors := entity.NewCustomer("", "fulano@gmail.com")

	require.NotNil(t, errors)
	require.Equal(t, errors[0].Value, "Name is required")
	require.Equal(t, len(errors), 1)
	require.Nil(t, customer)
}

func TestValidateIfEmailIsEmpty(t *testing.T) {
	customer, errors := entity.NewCustomer("fulano", "")

	require.NotNil(t, errors)
	require.Equal(t, errors[0].Value, "E-mail is required")
	require.Equal(t, len(errors), 1)
	require.Nil(t, customer)
}
