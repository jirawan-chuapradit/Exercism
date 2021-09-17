package account

import "sync"

// Open(initialDeposit int64) *Account
// (*Account) Close() (payout int64, ok bool)
// (*Account) Balance() (balance int64, ok bool)
// (*Account) Deposit(amount int64) (newBalance int64, ok bool)

type Account struct {
	mu sync.Mutex
	amount int64
	status string
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{ amount: initialDeposit, status: "available"}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.status != "available" {
		return 0, false
	}

	a.status = "unavailable"
	return a.amount, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.status != "available" {
		return 0,false
	}

	return a.amount, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool){
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.status != "available"  {
		return 0,false
	}

	// withdraw
	if amount < 0 && ( a.amount + (amount) < 0) {
		return 0,false
	}



	a.amount += amount
	return a.amount, true

}
