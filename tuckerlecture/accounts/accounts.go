package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw you are poor")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// Account내 변수를 변경하고자 한다면 *(포인터)를 사용해야함!
func (recv_this *Account) Deposit(amount int) {
	recv_this.balance += amount
}

// GetBalance is
// 단순히 참조만 하기위한 함수라면 *(포인터) 사용하지 않아도됨!(단 값의 복사가 일어남)
func (recv_this Account) GetBalance() int {
	return recv_this.balance
}

// Withdraw is
func (recv_this *Account) Withdraw(amount int) error {
	if recv_this.balance < amount {
		return errNoMoney
	}

	recv_this.balance -= amount
	return nil
}

// ChangeOwner of the account
func (recv_this *Account) ChangeOwner(newOwner string) {
	recv_this.owner = newOwner
}

// GetOwner is
func (recv_this Account) GetOwner() string {
	return recv_this.owner
}

// String is
func (recv_this Account) String() string {
	return fmt.Sprint(recv_this.GetOwner(), "'s account.\nHas: ", recv_this.GetBalance())
}
