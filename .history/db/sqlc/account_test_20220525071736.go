package db

import (
	"context"
)

func TestCreateAccount(t *testing) {
	arg := CreateAccountParams{
		Owner:    "tom",
		Balance:  100,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	if err != nil {
		return nil, err
	}
}
