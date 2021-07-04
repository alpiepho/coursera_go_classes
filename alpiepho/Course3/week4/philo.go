package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Implement the dining philosopher’s problem with the following constraints/modifications.
//
// There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
// Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
// The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
// In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
// The host allows no more than 2 philosophers to eat concurrently.
// Each philosopher is numbered, 1 through 5.
// When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
// When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

// Submission: Upload your source code for the program.
// The goals of this activity are to explore the design of concurrent algorithms and resulting synchronization issues.

// Review criteria
// less
// The maximum score of this assignment is 10 points.

// Students will receive points based on whether all the design criteria are met: 10 points if all criteria are met and 6 points (maximum) if more than one of the criteria are NOT met.

// How it works:
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
		// get locks on chopsticks randomly
		p.randLock()

		// ask permissions from main app
		ask <- p.id
		// get answer from main app
		start := <-p.start
		if start {
			// granted permission
			fmt.Printf("starting to eat %d\n", id)
			count++
			fmt.Printf("finishing eating %d\n", id)
			finished <- true
		}

		// release locks on chopsticks randomly
		p.randUnock()
	}
	done <- true
}

func main() {
	// build chopsticks
	csticks := make([]*ChopS, TOTAL)
	for i := 0; i < TOTAL; i++ {
		csticks[i] = new(ChopS)
	}
	// build philosophers and assign chopsticks, also create start[i] channel
	philos := make([]*Philo, TOTAL)
	for i := 0; i < TOTAL; i++ {
		philos[i] = &Philo{
			id:      i,
			leftCS:  csticks[i],
			rightCS: csticks[(i+1)%TOTAL],
			start:   make(chan bool, 1),
		}
	}

	// creat global channels for philosphers to ask, report finished eating and done eating (3 times)
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
		// watch for ask requests
		case id := <-ask:
			if activeCount < ALLOWED {
				activeCount++
				philos[id].start <- true
			} else {
				philos[id].start <- false
			}
		// watch for finished eating
		case <-finished:
			if activeCount > 0 {
				activeCount--
			}
		// watch for done eating (3 times)
		case <-done:
			doneCount++
		}
	}
}
