package main

import (
	"fmt"
	"sync"
)

type Data interface {
	Print()
}

type Coba struct {
	Value int
}

func (i1 Coba) Print() {
	fmt.Println("[coba1 coba2 coba3]", i1.Value)
}

type Bisa struct {
	Value int
}

func (i2 Bisa) Print() {
	fmt.Println("[bisa1 bisa2 bisa3]", i2.Value)
}

func main() {
	interface1Data := Coba{Value: 1}
	interface2Data := Coba{Value: 2}
	interface3Data := Coba{Value: 3}
	interface4Data := Coba{Value: 4}
	interface5Data := Bisa{Value: 1}
	interface6Data := Bisa{Value: 2}
	interface7Data := Bisa{Value: 3}
	interface8Data := Bisa{Value: 4}
	var wg sync.WaitGroup
	for i := 1; i < 9; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			data := []Data{interface1Data, interface2Data, interface3Data, interface4Data, interface5Data, interface6Data, interface7Data, interface8Data}
			selectedData := data[idx%8]

			selectedData.Print()

		}(i)
	}
	wg.Wait()
}
