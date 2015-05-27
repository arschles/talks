package main

import "errors"

var ErrNotReserved = errors.New("not reserved")

// Reserver is a map[string]interface{} where each key/value pair
// can be reserved by callers for mutually exclusive read-modify-write
// operations.
type Reserver interface {
	// GetAndReserve waits for the key to be unreserved, then reserves it for mutual exclusion.
	// On success, returns the current state and reservation ID. Use the latter in
	// future calls that require a reservation ID. If a non-nil error is returned, no
	// reservation is established and the returned value and reservation ID are invalid.
	GetAndReserve(key string) (val interface{}, reservationID string, err error)

	// SetReserved sets the value of the given key if reservationID is valid
	// and points to the current reservation. After this call returns successfully,
	// the caller doesn't have ownership of key and reservationID is invalid.
	// Returns ErrNotReserved if the reservation ID is invalid or not the current reservation.
	// On any non-nil error, neither the value nor the current reservation are changed.
	SetReserved(key, reservationID string, value interface{}) error
}
