package db

import (
	"context"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing) {
	arg := CreateAccountParams{
		Owner:    "tom",
		Balance:  100,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
}
