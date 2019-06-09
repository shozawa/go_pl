package bank

type withdraw struct {
	amount int
	ok     chan bool
}

var balances = make(chan int)
var deposits = make(chan int)
var withdraws = make(chan withdraw)

// Deposit amount
func Deposit(amount int) { deposits <- amount }

// Balance returns balance
func Balance() int { return <-balances }

// Withdraw amount from balance
func Withdraw(amount int) bool {
	ok := make(chan bool)
	withdraws <- withdraw{amount, ok}
	return <-ok
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case w := <-withdraws:
			if w.amount <= balance {
				balance -= w.amount
				w.ok <- true
			} else {
				w.ok <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
