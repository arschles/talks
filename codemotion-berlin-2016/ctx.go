package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// compute the factorial of num. return an error if ctx.Done() receives (i.e. a timeout or
// manual cancellation) or num < 0
func fact(ctx context.Context, num int64) (int64, error) {
	// either context was cancelled or timed out, or we continue.
	// the power of a select!
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}

	// edge cases
	if num < 0 {
		return 0, fmt.Errorf("invalid number %d", num)
	} else if num <= 1 {
		return 1, nil
	}

	// the recursion - note that we pass the same ctx down the stack
	f, err := fact(ctx, num-1)
	if err != nil {
		return 0, err
	}
	return num * f, nil
}

func main() {
	// create the context. a few notes:
	//
	// 1. context.Background() is the "root" of the inheritance
	// 2. context.WithTimeout creates a new context, inherited from context.Background.
	//		it returns the new timeout context and a function that cancels the context manually
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	computed, err := fact(ctx, 20)
	if err != nil {
		log.Fatalf("error (%s)", err)
		return
	}
	log.Println(computed)
}
