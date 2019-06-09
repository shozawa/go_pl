package bank2

import "sync"

var (
	mu      sync.Mutex
	balance int
)

// Deposit amount
func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

// Balance returns balance
func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}
