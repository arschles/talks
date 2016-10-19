package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func fact(ctx context.Context, num int64) (int64, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}
	if num < 0 {
		return 0, fmt.Errorf("invalid number %d", num)
	} else if num <= 1 {
		return 1, nil
	}
	f, err := fact(ctx, num-1)
	if err != nil {
		return 0, err
	}
	return num * f, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	computed, err := fact(ctx, 20)
	if err != nil {
		log.Fatalf("error (%s)", err)
		return
	}
	log.Println(computed)
}
