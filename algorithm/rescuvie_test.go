package algorithm

import (
	"fmt"
	"testing"
	"time"
)

func fibonacciFor(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func TestName1(t *testing.T) {
	c := make(chan int, 10)
	go fibonacciFor(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func TestName2(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciSelect(c, quit)
}

func TestName3(t *testing.T) {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func Rescuvie(n int) int {
	if n == 0 {
		return 1
	}
	return n * Rescuvie(n-1)
}

func RescuvieTail(n, a int) int {
	if n == 1 {
		return a
	}
	return RescuvieTail(n-1, a*n)
}

func TestRescuvie(t *testing.T) {
	fmt.Println(Rescuvie(5))
	fmt.Println(RescuvieTail(5, 1))
}
