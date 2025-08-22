package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	id             int
	leftChopstick  *sync.Mutex
	rightChopstick *sync.Mutex
}

func (p *Philosopher) Think() {
	fmt.Printf("Philosopher %d is thinking for 200ms", p.id)
	time.Sleep(200 * time.Millisecond)
}

func (p *Philosopher) Eat() {
	p.leftChopstick.Lock()
	defer p.leftChopstick.Unlock()
	p.rightChopstick.Lock()
	defer p.rightChopstick.Unlock()
	fmt.Printf("Philosopher %d is eating for 200ms", p.id)
	time.Sleep(200 * time.Millisecond)
}

// Live is just a wrapper method for the eat and think activities for k times
func (p *Philosopher) Live(k int, wg *sync.WaitGroup){
	// Decrement the wait group
	defer wg.Done()
	for i := 0; i < k; i++{ //  A philosopher life is an endless loop of eating and thinking
		p.Eat()
		p.Think()
	}
}

func main() {
	// Create 5 chopsticks and 5 philosophers
	numPhilos := 5
	chopSticks := make([]sync.Mutex, numPhilos)
	for i := 0; i < numPhilos; i++{
		chopSticks[i] = sync.Mutex{}
	}

	// Initiate the N philosophers and attribute their left and right chopstick/mutex (pointer to a mutex)
	philos := make([]Philosopher, numPhilos)
	for i := 0; i < numPhilos; i++{
		philos[i] = Philosopher{
			id: i,
			leftChopstick: &chopSticks[i],
			rightChopstick: &chopSticks[(i+1) % numPhilos], // We use mod for circular behaviour, they are sitting in a ring like setup, so philosopher 4 will have a right chopstick with ID 0
		}
	}

	// numCycle represents how many times the philosopher run the Eat and Think activities 
	// You could also change the code to have the philosopher eat and think an indefinitely
	numCycles := 10
	var wg sync.WaitGroup
	// Now run Live method for each philosopher as a go routine
	for _,ph := range philos{
		wg.Add(1)
		go ph.Live(numCycles, &wg)
	}
	wg.Wait()
	// This should result in a deadlock, to fix, one philosopher should start with their right chopstick
}
