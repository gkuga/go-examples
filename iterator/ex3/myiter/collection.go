package myiter

type Collection interface {
	CreateIterator() Iterator
}
