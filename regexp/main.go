package main

import (
	"fmt"
	"regexp"
	"sync"
)

var finder = func() func(string) []string {
	re := regexp.MustCompile("hello")
	return func(str string) []string {
		re := re.Copy()
		return re.FindAllString(str, -1)
	}
}()

func main() {
	ss := []string{
		"Hello World!",
		"hello world!",
	}
	var wg sync.WaitGroup
	for _, s := range ss {
		tmp := s
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("%s: %s\n", tmp, finder(tmp))
		}()
	}
	wg.Wait()
}
