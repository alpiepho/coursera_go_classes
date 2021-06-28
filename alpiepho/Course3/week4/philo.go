package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
}

func (p Philo) eat() {
	for {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("eating")

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}

	// will deadlock
	// dykstra: each picks up lowest numbered chopstick first
	// last may starve
}

func test() {

}

func main() {
	Csticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		Csticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{Csticks[i], Csticks[(i+1)%5]}
	}

	// start eating
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}

	// block - TBD, waitgroup?
}
