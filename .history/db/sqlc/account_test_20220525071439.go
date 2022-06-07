package db

func TestCreateAccount(t *testing) {
	arg := CreateAccountParams{
		Owner:    "tom",
		Balance:  100,
		Currency: "USD",
	}
}
