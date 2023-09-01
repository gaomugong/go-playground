package wallets

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			// Bitcoin implement Stringer interface
			t.Errorf("got %s, want %s", got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		fmt.Println("wallet.balance address in test", &wallet.balance)
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	assertError := func(t *testing.T, err error, want string) {
		if err == nil {
			t.Errorf("didn't get an error but wanted one")
		}

		if got := err.Error(); got != want {
			t.Errorf("get %s, want %s", got, want)
		}
	}

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{balance: startBalance}
		err := wallet.Withdraw(100)

		assertBalance(t, wallet, startBalance)
		assertError(t, err, "cannot withdraw, insufficient funds")
	})

}

// wallet.balance address in test 0x140001161d0
// wallet.balance address in method 0x140001161d8
