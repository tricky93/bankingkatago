// Bank Kata
// Develop an application that can deposit funds, withdraw funds and print a statement of a customer’s
// bank account activity.
// Here's the specification for an acceptance test that expresses the desired behaviour for this:
// Given a client makes a deposit of 1000 on 10/01/2021
// And a deposit of 2000 on 13/01/2021
// And a withdrawal of 500 on 14/01/2021
// When they print their bank statement
// Then they would see:
// Date || Amount || Balance
// 14/01/2012 || -500 || 2500
// 13/01/2012 || 2000 || 3000
// 10/01/2012 || 1000 || 1000
// Further Guidance
// • Writing your tests first may help steer you in your design, you can use a testing framework
// such as Mocha.
// • Don't worry about spacing and indentation in the statement output. (You could instruct your
// acceptance test to ignore whitespace if you wanted to.)
// • Use the acceptance test to guide your progress towards the solution.
// • When in doubt, go for the simplest solution!
// • You can Google syntax.
package main

import (
	"errors"
	"fmt"
	"time"
)


type account struct {
	name string
	balance float64
	transactions []transaction
}

type transaction struct {
	date time.Time
	amount float64
	balance float64
}

func main() {

	return
}

func newAccount(name string) *account {
	a := account{name: name}
	return &a
}

func (a *account) getBalance() float64 {
	return a.balance
}

func (a *account) deposit(amount float64) {
	t := transaction{
		date: time.Now(),
		amount: amount,
		balance: a.balance + amount,
	}
	a.balance += amount
	a.transactions = append(a.transactions, t)
}

func (a *account) withdraw(amount float64) error {

	if (a.balance - amount) < 0 {
		return errors.New("insufficient funds")
	} 

	t := transaction {
		date: time.Now(),
		amount: -amount,
		balance: a.balance - amount,
	}
	a.balance -= amount
	a.transactions = append(a.transactions, t)

	return nil
}

func (a *account) getStatement() {
	fmt.Printf("Date || Amount || Balance\n")

	for _, v := range a.transactions {
		day := v.date.Day()
		month := v.date.Month()
		year := v.date.Year()
		fmt.Printf("%v/%v/%v || %v || %v\n", day, month, year, v.amount, v.balance)
	}
}