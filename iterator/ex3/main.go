package main

import (
	"fmt"
	"github.com/gkuga/go-examples/iterator/ex3/iter"
	"github.com/gkuga/go-examples/iterator/ex3/myiter"
)

func runMyiter() {
	user1 := &myiter.User{
		Name: "a",
		Age:  30,
	}
	user2 := &myiter.User{
		Name: "b",
		Age:  20,
	}

	userCollection := &myiter.UserCollection{
		Users: []*myiter.User{user1, user2},
	}

	iterator := userCollection.CreateIterator()

	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}

func runIter() {
	user1 := &iter.User{
		Name: "a",
		Age:  30,
	}
	user2 := &iter.User{
		Name: "b",
		Age:  20,
	}

	users := iter.CreateIterator([]*iter.User{user1, user2})

	for user := range users {
		fmt.Printf("User is %+v\n", user)
	}
}

func main() {
	runMyiter()
	runIter()
}
