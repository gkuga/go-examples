package main

// #include <stdio.h>
// void hello() {
//    printf("Hello from C");
//}
import "C"

func main() {
	// let's call it
	C.hello()
}
