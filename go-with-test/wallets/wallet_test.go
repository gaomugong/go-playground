package wallets

import (
	"fmt"
	"testing"
)

//指针
//当你传值给函数或方法时，Go 会复制这些值。因此，如果你写的函数需要更改状态，你就需要用指针指向你想要更改的值
//Go 取值的副本在大多数时候是有效的，但是有时候你不希望你的系统只使用副本，在这种情况下你需要传递一个引用

//https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
//nil
//指针可以是 nil
//当函数返回一个的指针，你需要确保检查过它是否为 nil，否则你可能会抛出一个执行异常，编译器在这里不能帮到你
//nil 非常适合描述一个可能丢失的值

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()

	if got != want {
		// Bitcoin implement Stringer interface
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("got an error but didn't want one")
	}
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Errorf("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("get %s, want %s", got, want)
	}
}

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		fmt.Println("wallet.balance address in test", &wallet.balance)
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{balance: startBalance}
		err := wallet.Withdraw(100)

		assertBalance(t, wallet, startBalance)
		assertError(t, err, InsufficientFundsError)
	})

}

// wallet.balance address in test 0x140001161d0
// wallet.balance address in method 0x140001161d8
