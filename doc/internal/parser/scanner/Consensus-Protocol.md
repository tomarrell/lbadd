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

A detailed description of all the modules follow:

### Leader Election
* Startup: All servers start in the follower state and begin by requesting votes to be elected as a leader.
* Pre-election: The server increments its `currentTerm` by one, changes to `candidate` state and sends out `RequestVotes` RPC parallely to all the peers. 
* Vote condition: FCFS basis. If there was no request to the server, it votes for itself (read 3.6 and clear out when to vote for itself)
* Election timeout: A preset time for which the server waits to see if a peer requested a vote. It is randomly chosen between 150-300ms.
* Election is repeated after an election timeout until:
  1. The server wins the election
  2. A peer establishes itself as leader.
  3. Election timer times out or a split vote occurs (leading to no leader) and the process will be repeated.
 * Election win: Majority votes in the term. (More details in safety) The state of the winner is now `Leader` and the others are `Followers`.
 * Maintaining leaders reign: The leader sends `heartbeats` to all servers to establish its reign. This also checks whether other servers are alive based on the response and informs other servers that the leader still is alive too. If the servers do not get timely heartbeat messages, they transform from the `follower` state to `candidate` state.
 * Transition from working state to Election happens when a leader fails.
 * Maintaining sanity: While waiting for votes, if a `AppendEntriesRPC` is received by the server, and the term of the leader is greater than of equal to the "waiter"'s term, the leader is considered to be legitimate and the waiter becomes a follower of the leader. If the term of the leader is lesser, it is rejected.
 * The split vote problem: Though not that common, split votes can occur. To make sure this doesnt continue indefinitely, election timeouts are randomised,
 
### Log Replication

### Safety

## A strict testing mechanism

The testing mechanism to be implemented will enable us in figuring out the problems existing in the implementation leading to a more resilient system.
We have to test for the following basic failures:
1. Network partitioning.
2. Un-responsive peers.
3. Overloading peer.
4. Corruption of data in transit.

## Graceful handling of failures

Accepting failures exist and handling them gracefully enables creation of more resilient and stable systems. Having _circuit breakers_, _backoff mechanisms in clients_ and _validation and coordination mechanisms_ are some of the pointers to be followed. 
