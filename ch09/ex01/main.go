package main

import (
	"fmt"

	"github.com/shozawa/go_pl/ch09/ex01/bank"
)

func main() {
	bank.Deposit(100)
	fmt.Println(bank.Withdraw(100))
	fmt.Println(bank.Withdraw(100))
	fmt.Println(bank.Balance())
	bank.Deposit(100)
	fmt.Println(bank.Withdraw(50))
	fmt.Println(bank.Balance())
}
