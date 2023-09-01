package wallets

import "fmt"

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
