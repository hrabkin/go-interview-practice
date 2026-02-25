// Package challenge7 contains the solution for Challenge 7: Bank Account with Error Handling.
package challenge7

import (
	"sync"
	// Add any other necessary imports
	"strings"
	"errors"
    "fmt"
)

// BankAccount represents a bank account with balance management and minimum balance requirements.
type BankAccount struct {
	ID         string
	Owner      string
	Balance    float64
	MinBalance float64
	mu         sync.RWMutex // For thread safety
}

// Constants for account operations
const (
	MaxTransactionAmount = 10000.0 // Example limit for deposits/withdrawals
)

// Custom error types
var (
    ErrMaxTransactionExceeded = errors.New("max transaction amount exceeded")
    ErrInitialBalanceLessThanMinBalance = errors.New("initial balance is less than min balance")
)

// AccountError is a general error type for bank account operations.
type AccountError struct {
	// Implement this error type
	ID *string
}

func (e *AccountError) Error() string {
	// Implement error message
	if e.ID == nil {
	    return "account must have an ID defined"
	}
	return fmt.Sprintf("id=%s, account must have an owner defined", *e.ID)
}

// InsufficientFundsError occurs when a withdrawal or transfer would bring the balance below minimum.
type InsufficientFundsError struct {
	// Implement this error type
	ID string
	Owner string
	Amount float64
	Balance float64
	MinBalance float64
	TransferType string
}

func (e *InsufficientFundsError) Error() string {
	// Implement error message
	return fmt.Sprintf("id=%s, owner=%s: insufficient fund $%.2f for transfer type %s, balance $%.2f cannot be less than min balance $%.2f", e.ID, e.Owner, e.Amount, e.TransferType, e.Balance, e.MinBalance)
}

// NegativeAmountError occurs when an amount for deposit, withdrawal, or transfer is negative.
type NegativeAmountError struct {
	// Implement this error type
	ID string
	Owner string
	Amount float64
	AmountType string
}

func (e *NegativeAmountError) Error() string {
	return fmt.Sprintf("id=%s, owner=%s: amount for `%s` cannot be negative $%.2f", e.ID, e.Owner, e.AmountType, e.Amount)
}

// ExceedsLimitError occurs when a deposit or withdrawal amount exceeds the defined limit.
type ExceedsLimitError struct {
	ID string
	Owner string
	Amount float64
	TransferType string
}

func (e *ExceedsLimitError) Error() string {
	// Implement error message
	return fmt.Sprintf("id=%s, owner=%s: amount $%.2f exceeds limit $%.2f for transfer type `%s`", e.ID, e.Owner, e.Amount, MaxTransactionAmount, e.TransferType)
}

// NewBankAccount creates a new bank account with the given parameters.
// It returns an error if any of the parameters are invalid.
func NewBankAccount(id, owner string, initialBalance, minBalance float64) (*BankAccount, error) {
	// Implement account creation with validation
	if strings.TrimSpace(id) == "" {
	    return nil, &AccountError{} 
	}
	if strings.TrimSpace(owner) == "" {
	    return nil, &AccountError{ ID: &id } 
	}
	if initialBalance < 0 {
	    return nil, &NegativeAmountError{ ID: id, Owner: owner, Amount: initialBalance, AmountType: "Initial Balance"}
	}
	if minBalance < 0 {
	    return nil, &NegativeAmountError{ ID: id, Owner: owner, Amount: minBalance, AmountType: "Min Balance" }
	}
	if initialBalance < minBalance {
	    return nil, &InsufficientFundsError{ ID: id, Owner: owner, Amount: initialBalance, Balance: initialBalance, MinBalance: minBalance, TransferType: "NewAccount" }
	}
	return &BankAccount{ ID: id, Owner: owner, Balance: initialBalance, MinBalance: minBalance }, nil
}

// Deposit adds the specified amount to the account balance.
// It returns an error if the amount is invalid or exceeds the transaction limit.
func (a *BankAccount) Deposit(amount float64) error {
	if amount < 0 {
	    return &NegativeAmountError{ ID: a.ID, Owner: a.Owner, Amount: amount, AmountType: "Deposit" }
	}
	if amount > MaxTransactionAmount {
	    return &ExceedsLimitError{ ID: a.ID, Owner: a.Owner, Amount: amount, TransferType: "Deposit" }
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
	return nil
}

// Withdraw removes the specified amount from the account balance.
// It returns an error if the amount is invalid, exceeds the transaction limit,
// or would bring the balance below the minimum required balance.
func (a *BankAccount) Withdraw(amount float64) error {
	if amount < 0 {
	    return &NegativeAmountError{ ID: a.ID, Owner: a.Owner, Amount: amount, AmountType: "Withdraw" }
	}
	if amount > MaxTransactionAmount {
	    return &ExceedsLimitError{ ID: a.ID, Owner: a.Owner, Amount: amount, TransferType: "Withdraw" }
	}
	if a.Balance-amount < a.MinBalance {
    	a.mu.RLock()
    	defer a.mu.RUnlock()
	    return &InsufficientFundsError{ ID: a.ID, Owner: a.Owner, Amount: amount, Balance: a.Balance, MinBalance: a.MinBalance, TransferType: "Withdraw" }
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance -= amount
	return nil
}

// Transfer moves the specified amount from this account to the target account.
// It returns an error if the amount is invalid, exceeds the transaction limit,
// or would bring the balance below the minimum required balance.
func (a *BankAccount) Transfer(amount float64, target *BankAccount) error {
    if amount < 0 {
	    return &NegativeAmountError{ ID: a.ID, Owner: a.Owner, Amount: amount, AmountType: "Transfer" }
	}
	if amount > MaxTransactionAmount {
	    return &ExceedsLimitError{ ID: a.ID, Owner: a.Owner, Amount: amount, TransferType: "Transfer" }
	}
	if a.Balance-amount < a.MinBalance {
    	a.mu.RLock()
    	defer a.mu.RUnlock()
	    return &InsufficientFundsError{ ID: a.ID, Owner: a.Owner, Amount: amount, Balance: a.Balance, MinBalance: a.MinBalance, TransferType: "Transfer" }
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance -= amount
	target.mu.Lock()
	defer target.mu.Unlock()
	target.Balance += amount
	return nil
}