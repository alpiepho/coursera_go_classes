package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const TOTAL = 5
const ALLOWED = 2

type ChopS struct{ sync.Mutex }

type Philo struct {
	id      int
	leftCS  *ChopS
	rightCS *ChopS
	start   chan bool
}

func (p Philo) randLock() {
	if rand.Intn(2) == 0 {
		p.leftCS.Lock()
		p.rightCS.Lock()
	} else {
		p.rightCS.Lock()
		p.leftCS.Lock()
	}
}

func (p Philo) randUnock() {
	if rand.Intn(2) == 0 {
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	} else {
		p.leftCS.Unlock()
		p.rightCS.Unlock()
	}
}

func (p Philo) eat(ask chan int, finished chan bool, done chan bool) {
	id := p.id + 1
	count := 0
	for count < 3 {
		p.randLock()

		// ask permissions
		ask <- p.id
		start := <-p.start
		if start {
			fmt.Printf("starting to eat %d\n", id)
			count++
			fmt.Printf("finishing eating %d\n", id)
			finished <- true
		}

		p.randUnock()
	}
	done <- true
}

func main() {
	csticks := make([]*ChopS, TOTAL)
	for i := 0; i < TOTAL; i++ {
		csticks[i] = new(ChopS)
	}
	philos := make([]*Philo, TOTAL)
	for i := 0; i < TOTAL; i++ {
		philos[i] = &Philo{
			id:      i,
			leftCS:  csticks[i],
			rightCS: csticks[(i+1)%TOTAL],
			start:   make(chan bool, 1),
		}
	}

	// process:
	//   philo gets chopsticks randomly
	//   philo sends int over ask channel with id
	//   philo waits on start[id] channel for true/false
	//   if activeCount < ALLOWED
	//      main increments activeCount
	//      main sends TRUE over start[id] channel
	//      philo eats
	//      philo sends TRUE over finished channel
	//      if philo is done
	//         philo sends TRUE on done
	//         main increments doneCount
	//      main gets finished channel and decrements activeCount
	//   else
	//      main sends FALSE over start[id] channel
	//   philo releases chopsticks randmonly

	ask := make(chan int, TOTAL)
	finished := make(chan bool, ALLOWED)
	done := make(chan bool, TOTAL)
	for i := 0; i < TOTAL; i++ {
		go philos[i].eat(ask, finished, done)
	}

	// only allowing 2 to eat concurrently
	doneCount := 0
	activeCount := 0
	for doneCount < TOTAL {
		select {
		case id := <-ask:
			if activeCount < ALLOWED {
				activeCount++
				philos[id].start <- true
			} else {
				philos[id].start <- false
			}
		case <-finished:
			if activeCount > 0 {
				activeCount--
			}
		case <-done:
			doneCount++
		}
	}
}
