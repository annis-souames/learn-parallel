# Deadlocks, Livelocks and Starvation


### Deadlocks

Deadlocks arise when two or more concurrent threads are waiting on each other because of a locked shared resource. Imagine P1 which access A then B by locking each, and process P2 that access B then A by locking each as well. There might be a situation where Process P1 will lock A and B, then process P2 will lock B, process P1 will then try to access B which was already locked by P2 and P2 cannot access A which was already locked by P1. Both threads are waiting/asleep --> Deadlock. 

For a race condition to be classified as a deadlock it has to meet 4 conditions known as Coffman Conditions. If any of these conditions is not true, there is no deadlock. 

**Deadlocks can sometimes be avoided by breaking symmetry - see the example in Dining Philosophers.**

### Livelocks

Livelocks occur when two processes are blocked but they are still working, a real life example is when two people meet at a corridor and they both try to go in the same direction for several times. Livelocks are hard to identify because the system is indeed running.

### Starvation
Starvation is a case when a certain process doesn't get access to a shared a resource at all because other processes/threads keep locking those resources forever. A good example is also found in the dining philosopher problem.

Imagine a philosopher B sitting between philosopher A and philosopher C who are both very greedy and fast eaters, the philo A will pickup his left fork making philo B not able to eat, as soon as philo A finishes and puts it down and in the nanoseconds before the philo B start eating, his next neighbour, philo C, picks up his left fork, which is the right fork of B making B unable to eat. Now imagine this goes indefinitely, indeed A and C are running fine but one thread, B is starving. 

Starvation usually has less frequent compared to deadlocks but is still a very possible problem. 



**To avoid deadlocks and starvation in shared resources, see algorithms mentioned in [Mutual Exclusion Problem and Algorithms](./mutual-exclusion-algorithms.md)**