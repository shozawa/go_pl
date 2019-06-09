package bank2

var (
	sema    = make(chan struct{}, 1)
	balance int
)

// Deposit amount
func Deposit(amount int) {
	sema <- struct{}{}
	balance = balance + amount
	<-sema
}

// Balance returns balance
func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
