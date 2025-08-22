package main

import (
	"fmt"
	"sync"
)

type Philosopher struct {
	id             int
	leftChopstick  *sync.Mutex
	rightChopstick *sync.Mutex
}

func (p *Philosopher) Think() {
	fmt.Printf("Philosopher %d is thinking. \n", p.id)
	//time.Sleep(200 * time.Millisecond)
}

// Eat simulates the eating activity, if start is 'left' the philosopher will lock the left chopstick first and then the right chopstick, 
// otherwise, if start is right then the reverse order happen. We keep this parameter to control deadlocks for learning purposes.
// This method is also responsible for putting down the chopsticks (i.e unlocking the mutexes)
func (p *Philosopher) Eat(start string) {
	if start == "left" {
		p.leftChopstick.Lock()
		defer p.leftChopstick.Unlock()
		p.rightChopstick.Lock()
		defer p.rightChopstick.Unlock()
	}else{
		p.rightChopstick.Lock()
		defer p.rightChopstick.Unlock()	
	}
	fmt.Printf("Philosopher %d is eating. \n", p.id)
	//time.Sleep(200 * time.Millisecond)
}

// Live is just a wrapper method for the eat and think activities for k times
// startChopstick is a string passed to the Eat method, read Eat comments for more.
// wg is a pointer to a Waitgroup for waiting all philo life cycles to end.
func (p *Philosopher) Live(k int, startChopstick string, wg *sync.WaitGroup){
	// Decrement the wait group
	defer wg.Done()
	for i := 0; i < k; i++{ //  A philosopher life is an endless loop of eating and thinking
		p.Eat(startChopstick)
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
		// This part is commented, it solves the deadlock issue by breaking the symmetry
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
	for id,ph := range philos{
		wg.Add(1)
		// The life of each philosopher is a goroutine that runs concurrently
		// If we use left as the starting chopstick, we might face a deadlock if all philosopher start eating at same time and lock their left chopstick
		// We can fix this deadlock situation by breaking the cycle, we only need one philosopher to start with right.
		if id == 0{
			go ph.Live(numCycles, "right", &wg)
		}else{
			go ph.Live(numCycles, "left", &wg)
		}
		
	}
	wg.Wait()
	
}
