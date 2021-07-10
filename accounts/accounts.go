package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

var errMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errMoney
	}
	a.balance -= amount
	return nil

	// error 에는 두가지 value가있다. error객체를 return하는거와 nil이라는 것이 있다.
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// __str__ 이런거랑 비슷하다고 보면 된다.
func (a Account) String() string {
	return "whatever you want"
	// or
	// return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
