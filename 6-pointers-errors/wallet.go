package wallet

import (
	"errors"
	"fmt"
)

// Pointers
// - go copies values when you pass them to functions/methods, so if you are writing a function that needs to mutate state you will need it to take a pointer to te thing you want to change
// - the fact that go takes a copy of values is useful a lot of the time, but sometimes you dont want your system to make a copy of something. In which case you need to pass a reference. Examples include referencing very large data structures or things where only one instance is necessary

// nil
// Pointers can be nil
// when a function returns a pointer to somethings, you need to make sure you check if it's nil or you might raise a runtime exception, the compiler won't help you here
// useful for when you want to describe a value that could be missing

// Errors
// Errors are the way to signify failure when calling a function/method

// create new types from existing ones
// useful for adding more domain specific meaning to values
// can let you implement interfaces

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount

}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
