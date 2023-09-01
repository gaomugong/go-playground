package wallets

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		fmt.Println("wallet.balance address in test", &wallet.balance)

		wallet.Deposit(10)
		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			// Bitcoin implement Stringer interface
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(10)
		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

}

// wallet.balance address in test 0x140001161d0
// wallet.balance address in method 0x140001161d8
