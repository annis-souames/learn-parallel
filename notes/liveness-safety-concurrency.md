# Liveness and Safety in Concurrent Systems

These are two properties of a concurrent system:

- Safety: A bad thing will never happen, for example traffic lights should not go all green at the same time in an intersection, in programming this means that a shared resources should not be modified by several processes at the same time.
- Liveness: A good thing (in the system logical sense) will eventually happen: this means we don't have any deadlocks (infinite wait for a lock) and we don't have any starvation.

Usually real world systems will need have a certain trade off between the two: some systems might prefer safety over liveness (rather do nothing - no liveness then being not safe) such as critical software, while others will prefer liveness over safety.

Nice read: https://www.thecoder.cafe/p/safety-liveness
