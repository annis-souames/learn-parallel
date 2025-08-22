# Dining Philosophers Problem

The dining philosopher is an example problem on synchronization developed by Djikstra and changed by Tony Hoare. In this example we use chopsticks instead of forks as suggested by the Art of Multiprocessor Programming book. The naming is different, the same solution and logic still applies.

### Statement:

Five [philosophers](https://en.wikipedia.org/wiki/Philosopher "Philosopher") dine together at the same table. Each philosopher has their own plate at the table. There is a fork between each plate. The dish served is a kind of [spaghetti](https://en.wikipedia.org/wiki/Spaghetti "Spaghetti") which has to be eaten with two forks. Each philosopher can only alternately think and eat. Moreover, a philosopher can only eat their spaghetti when they have both a left and right fork. Thus two forks will only be available when their two nearest neighbors are thinking, not eating. After an individual philosopher finishes eating, they will put down both forks. The problem is how to design a regimen (a [concurrent](https://en.wikipedia.org/wiki/Concurrency_\(computer_science\) "Concurrency (computer science)") algorithm) such that any philosopher will not starve; _i.e._, each can forever continue to alternate between eating and thinking, assuming that no philosopher can know when others may want to eat or think.

### Solution summary

All of the code is in main.go and has been fixed to avoid the problem mentioned in the Deadlock Gotcha section explained below. 

We model Philosophers as a simple struct with left and right chopsticks acting as pointers to `sync.Mutex`. Each chopstick is a Go mutex that is locked if it's held for eating and unlocked if it's put down. A philosopher can only eat if both left and right chopsticks (mutexes) are locked.

### Deadlock Gotcha

A simple solution would be to have a Eat method that lock both the left and right chopstick. The issue is that in the scenario where all philosophers start eating at the same time, each philosopher will lock their left chopstick, making the right chopstick of the next philosopher locked which would cause their thread to wait, this race condition results in a **deadlock**.

To fix this we need to break the symmetry by forcing one philosopher to start eating with the right chopstick and the other can still use the left chopstick.

Other solutions exist using CSP/channels approach.

