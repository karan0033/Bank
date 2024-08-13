package db

// func createRandomAccount(t *testing.T) Account {
// 	user := createRandomUser(t)
// 	arg := CreateAccountParams{
// 		Owner:    user.Username,
// 		Balance:  utils.RandomBalance(),
// 		Currency: utils.RandomCurrency(),
// 	}

// 	account, err := testQuery.CreateAccount(context.Background(), arg)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, account)

// 	require.Equal(t, arg.Balance, account.Balance)
// 	require.Equal(t, arg.Currency, account.Currency)
// 	require.Equal(t, arg.Owner, account.Owner)

// 	require.NotEmpty(t, account.ID)
// 	require.NotEmpty(t, account.CreatedAt)

// 	return account
// }

// func TestCreateAccount(t *testing.T) {
// 	createRandomAccount(t)
// }

// func TestGetAccoutn(t *testing.T) {
// 	account1 := createRandomAccount(t)

// 	account2, err := testQuery.GetAccount(context.Background(), account1.ID)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, account2)
// 	require.Equal(t, account1.ID, account2.ID)
// 	require.Equal(t, account1.Balance, account2.Balance)
// 	require.Equal(t, account1.Owner, account2.Owner)
// 	require.Equal(t, account1.Currency, account2.Currency)
// 	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
// }

// func TestUpdateAccount(t *testing.T) {
// 	account1 := createRandomAccount(t)

// 	arg := UpdateAccountParams{
// 		ID:      account1.ID,
// 		Balance: utils.RandomBalance(),
// 	}

// 	account2, err := testQuery.UpdateAccount(context.Background(), arg)

// 	require.NoError(t, err)
// 	require.NotEmpty(t, account2)

// 	require.Equal(t, account1.ID, account2.ID)
// 	require.Equal(t, account1.Currency, account2.Currency)
// 	require.Equal(t, arg.Balance, account2.Balance)
// 	require.Equal(t, account1.Owner, account2.Owner)
// 	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
// }

// func TestDeleteAccount(t *testing.T) {
// 	account1 := createRandomAccount(t)

// 	err := testQuery.DeleteAccount(context.Background(), account1.ID)
// 	require.NoError(t, err)

// 	account2, err := testQuery.GetAccount(context.Background(), account1.ID)
// 	require.Error(t, err)
// 	require.Empty(t, account2)
// }

// func TestListAccounts(t *testing.T) {
// 	var LastAccount Account
// 	for i := 0; i < 10; i++ {
// 		LastAccount = createRandomAccount(t)
// 	}

// 	arg := ListAccountsParams{
// 		Owner:  LastAccount.Owner,
// 		Limit:  5,
// 		Offset: 0,
// 	}

// 	accounts, err := testQuery.ListAccounts(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, accounts)

// 	for _, acc := range accounts {
// 		require.NotEmpty(t, acc)
// 		require.Equal(t, LastAccount.Owner, acc.Owner)
// 	}
// }
