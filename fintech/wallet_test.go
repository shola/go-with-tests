package fintech

import (
	"testing"
)

func TestWallet (t *testing.T) {
	assertBalance := func (t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	}
	t.Run("balance", func(t *testing.T){
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T){
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(5))
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		assertBalance(t, wallet, Bitcoin(5))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T){
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(20))
		assertBalance(t, wallet, Bitcoin(10))
		assertError(t, err)	
	})
}