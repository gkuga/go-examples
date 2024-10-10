package myiter

type Iterator interface {
	HasNext() bool
	GetNext() *User
}
