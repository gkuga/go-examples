package main

import (
	"fmt"
	"github.com/gkuga/go-examples/iterator/ex3/myiter"
)

func main() {
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
