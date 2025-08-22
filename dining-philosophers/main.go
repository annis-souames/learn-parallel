package main

import (
	"fmt"
	"sync"
	"time"
)


type Chopstick struct {
	id     int
	state bool
	mux    sync.Mutex
}

// Hold will set the chopstick state to true and 
func (cp *Chopstick) Hold(philId int) {
	cp.mux.Lock()
	cp.state = true
	// Logging to make sure that a chopstick is not being used twice by different philosophers
	fmt.Printf("Philo of ID %d is holding chopstick of ID %d", philId, cp.id)
}

// PutDown will set the state to false and unlock the chopstick
func (cp *Chopstick) PutDown(philId int) {
	cp.state = false
	// Logging to make sure that a chopstick is not put down twice by different philosophers
	fmt.Printf("Philo of ID %d is putting down chopstick of ID %d", philId, cp.id)
	// Unlock the mutex to allow other philosophers to use the chopstick
	cp.mux.Unlock()
}
type Philosopher struct {
	id             int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

func (p *Philosopher) Think() {
	p.leftChopstick.PutDown(p.id)
	p.rightChopstick.PutDown(p.id)
	fmt.Printf("Philosopher %d is thinking and has put down his chopsticks", p.id)
	time.Sleep(200 * time.Millisecond)
}

func (p *Philosopher) Eat() {
	p.leftChopstick.Hold(p.id)
	p.rightChopstick.Hold(p.id)
	fmt.Printf("Philosopher %d is eating", p.id)
	time.Sleep(200 * time.Millisecond)
}

// Live is just a wrapper method for the eat and think activities.
// It is important to start with eat activity to lock the mutexes, starting with think activity will cause a runtime error because we are trying to unlock a non-locked mutex.
func (p *Philosopher) Live(){
	p.Eat()
	p.Think()
}



func main() {
	// Create 5 philosophers
	// Create 5 chopsticks
	// Loop around the 5 infinitely in a circular way and run `go Live()`

}
