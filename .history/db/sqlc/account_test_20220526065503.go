package db

import (
	"context"
	"testing"
	"time"

	"github.com/mplayer78/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount((t))
}

func TestGetAccount(t *testing.T) {
	testAccount := createRandomAccount(t)

	account, err := testQueries.GetAccount(context.Background(), testAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, testAccount.Owner, account.Owner)
	require.Equal(t, testAccount.Balance, account.Balance)
	require.Equal(t, testAccount.Currency, account.Currency)
	require.WithinDuration(t, testAccount.CreatedAt, account.CreatedAt, time.Second)
}
