# The consensus

Maintaining consensus is one of the major parts of a distributed system. To know to have achieved a stable system, we need the following two parts of implementation.

## The Raft protocol


## A strict testing mechanism

The testing mechanism to be implemented will enable us in figuring out the problems existing in the implementation leading to a more resilient system.
We have to test for the following basic failures:
1. Network partitioning.
2. Un-responsive peers.
3. Overloading peer.
4. Corruption of data in transit.

## Graceful handling of failures

Accepting failures exist and handling them gracefully enables creation of more resilient and stabler systems. Having _circuit breakers_, _backoff mechanisms in clients_ and _validation and coordination mechanisms_ are some of the pointers to be followed. 
