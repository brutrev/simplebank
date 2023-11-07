package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/brutrev/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.GenerateOwner(),
		Balance:  util.GenerateMoney(),
		Currency: util.GenerateCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	return account
}

func TestCreateAccount(t *testing.T) {
	account := createRandomAccount(t)
	expected, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)

	require.Equal(t, expected.ID, account.ID)
	require.Equal(t, expected.Owner, account.Owner)
	require.Equal(t, expected.Balance, account.Balance)
	require.Equal(t, expected.Currency, account.Currency)
	require.Equal(t, expected.ID, account.ID)
	require.WithinDuration(t, expected.CreatedAt, account.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.GenerateMoney(),
	}

	expected, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, expected)

	require.Equal(t, expected.ID, account.ID)
	require.Equal(t, expected.Owner, account.Owner)
	require.Equal(t, expected.Balance, arg.Balance)
	require.Equal(t, expected.Currency, account.Currency)
	require.Equal(t, expected.ID, account.ID)
	require.WithinDuration(t, expected.CreatedAt, account.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	expected, err := testQueries.GetAccount(context.Background(), account.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, expected)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
