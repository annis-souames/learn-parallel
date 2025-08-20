package main

import "fmt"

func Think(philoId int) {
	fmt.Printf("Philosopher %d is thinking", philoId)

}

func Eat(philoId int) {
	fmt.Printf("Philosopher %d is eating", philoId)
}

func main() {

	philosophers := [5]int{1, 2, 3, 4, 5}
	// The forks array acts like a flag system, with 0 meaning the fork is down and 1 meaning the fork is up
	var forks [5]int // This initiates the array to all elements to be zero - eq. to {0,0,0,0,0}

	for philo = range philosophers {
		go Think() // The philosopher will think
		go Eat()   // The philosopher will eat
	}

}
