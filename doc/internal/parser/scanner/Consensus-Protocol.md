# The consensus

Before talking about consensus, we need to discuss some logistics based on how the systems can co-exist.

* Communication: Distributed systems need a method to communicate between each other. Remote Procedure Calls is the mechanism using which a standalone server can talk to another. The standard Go package [RPC](https://golang.org/pkg/net/rpc/) serves us the purpose. 
* Security: Access control mechanisms need to be in place to decide on access to functions in the servers based on their state (leader, follower, candidate)

Maintaining consensus is one of the major parts of a distributed system. To know to have achieved a stable system, we need the following two parts of implementation.

## The Raft protocol

A raft server may be in any of the 3 states; leader, follower or candidate. All requests are serviced through the leader and it then decides how and if the logs must be replicated in the follower machines. The raft protocol has 3 almost independent modules:
1. Leader Election
2. Log Replication
3. Safety

## A strict testing mechanism

The testing mechanism to be implemented will enable us in figuring out the problems existing in the implementation leading to a more resilient system.
We have to test for the following basic failures:
1. Network partitioning.
2. Un-responsive peers.
3. Overloading peer.
4. Corruption of data in transit.

## Graceful handling of failures

Accepting failures exist and handling them gracefully enables creation of more resilient and stable systems. Having _circuit breakers_, _backoff mechanisms in clients_ and _validation and coordination mechanisms_ are some of the pointers to be followed. 
