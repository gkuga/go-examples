package foo

type bar struct {
	Val int
}

func Newbar(n int) *bar {
	return &bar{n}
}
