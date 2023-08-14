package main

import (
	"fmt"
	"sync"
)

type Interface1 interface {
	I1coba() string
}

type Interface2 interface {
	I2Bisa() string
}

type Coba struct {
	coba1 string
}

func (d Coba) I1coba() string {
	return "[coba1 coba2 coba3]"
}

type Bisa struct {
	Bisa2 string
}

func (c Bisa) I2Bisa() string {
	return "[bisa1 bisa2 bisa3]"
}

func PrintCoba(a Interface1, wg *sync.WaitGroup, cobaCh, vehicleCh chan struct{}) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		<-cobaCh // Menunggu isyarat dari vehicleCh
		fmt.Printf("%s %d\n" ,a.I1coba(),i)
		vehicleCh <- struct{}{} // Memberi isyarat kepada goroutine vehicle
	}
}

func PrintBisa(v Interface2, wg *sync.WaitGroup, cobaCh, vehicleCh chan struct{}) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		<-vehicleCh // Menunggu isyarat dari cobaCh
		fmt.Printf("%s %d\n" ,v.I2Bisa(),i)
		cobaCh <- struct{}{} // Memberi isyarat kepada goroutine animal
	}
}

func main() {
	var wg sync.WaitGroup
	cobaCh := make(chan struct{}, 1) // Channel untuk goroutine animal
	vehicleCh := make(chan struct{}, 1) // Channel untuk goroutine vehicle

	interface1 := Coba{coba1: "Buddy"}
	interface2 := Bisa{Bisa2: "Toyota"}

	wg.Add(2)
	// Memulai dengan goroutine animal
	cobaCh <- struct{}{}
	go PrintCoba(interface1, &wg, cobaCh, vehicleCh)
	go PrintBisa(interface2, &wg, cobaCh, vehicleCh)

	wg.Wait()
	close(cobaCh)
	close(vehicleCh)
}