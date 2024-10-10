package iter

import "iter"

type UserCollection iter.Seq[*User]

func CreateIterator(users []*User) UserCollection {
	return func(yield func(*User) bool) {
		for _, user := range users {
			yield(user)
		}
	}
}
