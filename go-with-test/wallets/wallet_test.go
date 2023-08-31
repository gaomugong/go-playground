package wallets

import (
	"fmt"
	"testing"
)

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	fmt.Println("wallet.balance address in method", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	fmt.Println("wallet.balance address in test", &wallet.balance)

	wallet.Deposit(10)
	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

// wallet.balance address in test 0x140001161d0
// wallet.balance address in method 0x140001161d8
