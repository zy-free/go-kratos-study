package limit

import (
	"errors"
)

// ErrLimitReturn indicates that the more than borrowed elements were returned.
var ErrLimitReturn = errors.New("discarding limited token, resource pool is full, someone returned multiple times")
// Limit controls the concurrent requests.
type Limit struct {
	pool chan struct{}
}

// NewLimit creates a Limit that can borrow n elements from it concurrently.
func NewLimit(n int) Limit {
	return Limit{
		pool: make(chan struct{}, n),
	}
}

// Borrow borrows an element from Limit in blocking mode.
func (l Limit) Borrow() {
	l.pool <- struct{}{}
}

// Return returns the borrowed resource, returns error only if returned more than borrowed.
func (l Limit) Return() error {
	select {
	case <-l.pool:
		return nil
	default:
		return ErrLimitReturn
	}
}

// TryBorrow tries to borrow an element from Limit, in non-blocking mode.
// If success, true returned, false for otherwise.
func (l Limit) TryBorrow() bool {
	select {
	case l.pool <- struct{}{}:
		return true
	default:
		return false
	}
}
