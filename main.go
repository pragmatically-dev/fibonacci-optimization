package main

import "fmt"

type Fibo struct {
	memory map[int64]int64
}

func (f *Fibo) fn(n int64) int64 {

	if k, v := f.memory[n]; v {
		return k
	}
	if n <= 0 {
		return 1
	}
	fn := f.fn(n-1) + f.fn(n-2)
	f.memory[n] = fn
	return f.memory[n]
}

func main() {

	fibonacci := new(Fibo)
	fibonacci.memory = make(map[int64]int64)
	for {
		var n int64
		fmt.Print("\n\nIngrese un numero: ")
		fmt.Scanf("%d", &n)
		fmt.Println("")
		fmt.Printf("fibonacci de %d  es: %d", n, fibonacci.fn(n))
	}

}
