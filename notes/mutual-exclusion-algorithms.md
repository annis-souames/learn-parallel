# Algorithms for Mutual Exclusion

Mutual exclusion problem is a fundemental problem in concurrency where two or more concurrent processes try to share resources. For example a variable that is incremented (non-atomic operation) at same time by two concurrent threads will result in a wrong value.

Below are some important/famous algorithms to solve the mutual exclusion problem. The goal of these algorithms is to avoid having 2 or more threads access the same CS

- Simple locking/unlocking: Easy to implement but can easily fall in deadlocks, livelocks and/or starvation (see [[Deadlock, Livelocks and Starvation]]).
- Peterson Algorithm:
	- Used for 2 threads
	- Is deadlock free 
	- Is starvation free
	- Uses flags and turn to avoid starvation.
- Deckers Algorithm:
	- For two threads
- Lamport's Bakery Algorithm:
	- For n threads
	- Uses same concept as ticketing in a bakery, only the thread that has the lowest ticket can proceed.
	- Is deadlock-free and starvation-free

There are other algorithms for the mutual exclusion problem (also known as Critical Section Problem).