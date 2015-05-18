package main

import "errors"

var ErrNotReserved = errors.New("not reserved")

// RequestState is an in-memory table to track the state of all requests.
// States can be any positive number.
type RequestState interface {
	// GetAndReserve waits for the state to be available, then returns it.
	// After this call returns, you have ownership of this state. Use
	// the returned string as an ownership token to indicate your ownership of id
	// in later calls. Returns an error if communication with underlying storage fails.
	// If this func returns error, you do not have ownership.
	GetAndReserve(id string) (int, string, error)

	// SetReserved sets the state of the given request ID, given the ownership ID.
	// doesn't set the state and returned ErrNotReserved if the state wasn't reserved
	// by the given ownership ID. After this call returns successfully, you no
	// longer have ownership of id and ownershipID is invalidated.
	SetReserved(id, ownershipID string, state int) error
}
