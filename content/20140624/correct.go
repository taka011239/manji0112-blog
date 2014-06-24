var a, b int
var ch chan int

func f() {
	a = 1
	b = 2
	<-ch
}

func g() {
	print(b)
	print(a)
}

func main() {
	ch = make(chan int)
	go f()
	ch <- 0
	g()
}
