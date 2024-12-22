package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	ctx := context.Background()
	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		fmt.Println("Task 1 started")
		time.Sleep(2 * time.Second)
		fmt.Println("Task 1 finished")
		return nil
	})

	group.Go(func() error {
		fmt.Println("Task 2 started")
		time.Sleep(1 * time.Second)
		fmt.Println("Task 2 encountered an error")
		return fmt.Errorf("error in task 2")
	})

	group.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("Task 3 canceled due to context cancellation")
			return ctx.Err()
		case <-time.After(3 * time.Second):
			fmt.Println("Task 3 finished")
			return nil
		}
	})

	if err := group.Wait(); err != nil {
		fmt.Printf("Group finished with error: %v\n", err)
	} else {
		fmt.Println("Group finished successfully")
	}
}
