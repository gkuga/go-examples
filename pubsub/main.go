package main

func main() {
	pub := Broadcaster[int]{subs: make(map[chan int]struct{})}

	sub1 := pub.Subscribe()
	sub2 := pub.Subscribe()

	pub.Publish(42)

	println(<-sub1)
	println(<-sub2)
	println(<-sub2)
}
