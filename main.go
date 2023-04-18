package main

import (
	"fmt"
	"sync"
)

func main() {
	a := []interface{}{"bisa1", "bisa2", "bisa3"}
	b := []interface{}{"coba1", "coba2", "coba3"}

	waitGroup := sync.WaitGroup{}
	mutex := sync.Mutex{}
	fmt.Println("Acak")
	waitGroup.Add(8)
	for i := 1; i <= 4; i++ {
		go random(a, i, &waitGroup)
		go random(b, i, &waitGroup)
	}
	waitGroup.Wait()

	fmt.Println("Rapih")
	waitGroup.Add(8)
	for i := 1; i <= 4; i++ {
		go rapih(a, b, i, &waitGroup, &mutex)
	}

	waitGroup.Wait()
}

func random(bisaCoba interface{}, i int, w *sync.WaitGroup) {
	fmt.Println(bisaCoba, i)
	w.Done()
}

func rapih(a, b interface{}, j int, w *sync.WaitGroup, m *sync.Mutex) {
	defer w.Done()
	m.Lock()
	defer m.Unlock()
	fmt.Println(b, j)
	fmt.Println(a, j)
}
