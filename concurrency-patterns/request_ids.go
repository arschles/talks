package main

import "errors"

var ErrNotReserved = errors.New("not reserved")

// RequestState is an in-memory table to track the state of all requests.
// States can be any possible uint64 - it's up to the caller to decide on the meaning of each.
type RequestState interface {
	// GetAndReserve waits for the state to be available, then returns it.
	// On success, you have ownership of this state. Use the returned string as an
	// ownership token to indicate your ownership of id in later calls.
	// Returns an error if communication with underlying storage fails.
	// On failure, you don't have ownership.
	GetAndReserve(id string) (uint64, string, error)

	// SetReserved sets the state of the given request ID, given the ownership ID.
	// Doesn't set the state and returns ErrNotReserved if the state wasn't reserved
	// by the given ownership ID. After this call returns successfully, you no
	// longer have ownership of id and ownershipID is invalidated. Do not reuse it.
	SetReserved(id, ownershipID string, state uint64) error
}
