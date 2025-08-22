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

// Live is just a wrapper method for the eat and think activities.
func (p *Philosopher) Live(){
	for{ //  A philosopher life is an endless loop of eating and thinking
		p.Eat()
		p.Think()
	}
}

func main() {
	// Create 5 philosophers
	

}
