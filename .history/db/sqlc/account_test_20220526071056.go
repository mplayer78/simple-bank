package db

import (
	"context"
	"database/sql"
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

func TestUpdateAccount(t *testing.T) {
	testAccount := createRandomAccount(t)

	newBalance := util.RandomMoney()

	arg := UpdateAccountParams{
		ID:      testAccount.ID,
		Balance: newBalance,
	}

	testAccount2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, testAccount2)

	require.Equal(t, testAccount2.Owner, testAccount.Owner)
	require.Equal(t, testAccount2.Balance, newBalance)
	require.Equal(t, testAccount2.Currency, testAccount.Currency)
	require.Equal(t, testAccount2.CreatedAt, testAccount.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	testAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), testAccount.ID)
	require.NoError(t, err)

	testAccount2, err := testQueries.GetAccount(context.Background(), testAccount.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, testAccount2)
}

func TestGetAllAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := GetAllAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.GetAllAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
