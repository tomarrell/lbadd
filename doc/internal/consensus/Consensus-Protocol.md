# The consensus

Before talking about consensus, we need to discuss some logistics based on how the systems can co-exist.

* Communication: Distributed systems need a method to communicate between each other. Remote Procedure Calls is the mechanism using which a standalone server can talk to another. The existing `network` layer of the database will handle all the communication between servers. 
* Security: Access control mechanisms need to be in place to decide on access to functions in the servers based on their state (leader, follower, candidate)
* Routing to leader: One of the issues with a varying leader is for the clients to know which IP address to contact for the service. We can solve this problem by advertising any/all IPs of the cluster and simply forward this request to the current leader; OR have a proxy that can forward the request to the current leader wheneve the requests come in. (Section client interaction of post has another approach which works too)
* The servers will be implemented in the `interal/node` folders which will import the raft API and perform their functions.

Maintaining consensus is one of the major parts of a distributed system. To know to have achieved a stable system, we need the following two parts of implementation.

## The Raft protocol

The raft protocol will be implemented in `internal/raft` and will implement APIs that each node can call.

Raft is an algorithm to handle replicated log, and we maintain the "log" of the SQL stmts applied on a DB and have a completely replicated cluster.

#### General Implementation rules:
* All RPC calls are done parallely to obtain the best performance.
* Request retries are done in case of network failures.
* Raft does not assume network preserves ordering of the packets.

A raft server may be in any of the 3 states; leader, follower or candidate. All requests are serviced through the leader and it then decides how and if the logs must be replicated in the follower machines. The raft protocol has 3 almost independent modules:
1. Leader Election
2. Log Replication
3. Safety

A detailed description of all the modules follow:

### Leader Election

#### Spec
* Startup: All servers start in the follower state and begin by requesting votes to be elected as a leader.
* Pre-election: The server increments its `currentTerm` by one, changes to `candidate` state and sends out `RequestVotes` RPC parallely to all the peers. 
* Vote condition: FCFS basis. If there was no request to the server, it votes for itself (read 3.6 and clear out when to vote for itself)
* Election timeout: A preset time for which the server waits to see if a peer requested a vote. It is randomly chosen between 150-300ms.
* Election is repeated after an election timeout until:
  1. The server wins the election
  2. A peer establishes itself as leader.
  3. Election timer times out or a split vote occurs (leading to no leader) and the process will be repeated.
 * Election win: Majority votes in the term. (More details in safety) The state of the winner is now `Leader` and the others are `Followers`.
 * The term problem: Current terms are exchanged when-ever servers communicate; if one server’s current term is smaller than the other’s, then it updatesits current term to the larger value. If a candidate or leader discovers that its term is out of date,it immediately reverts to follower state. If a server receives a request with a stale term number, itrejects the request.
 * Maintaining leaders reign: The leader sends `heartbeats` to all servers to establish its reign. This also checks whether other servers are alive based on the response and informs other servers that the leader still is alive too. If the servers do not get timely heartbeat messages, they transform from the `follower` state to `candidate` state.
 * Transition from working state to Election happens when a leader fails.
 * Maintaining sanity: While waiting for votes, if a `AppendEntriesRPC` is received by the server, and the term of the leader is greater than of equal to the "waiter"'s term, the leader is considered to be legitimate and the waiter becomes a follower of the leader. If the term of the leader is lesser, it is rejected.
 * The split vote problem: Though not that common, split votes can occur. To make sure this doesnt continue indefinitely, election timeouts are randomised, making the split votes less probable.
 

#### Implementation

* The raft module will provide a `StartElection` function that enables a node to begin election. This function just begins the election and doesnt return any result of the election. The decision of the election will be handled by the votes and each server independently. 
* The Leader node is the only one to know its the leader in the beginning. It realises it has obtained the majority votes, and starts behaving like the leader node. During this period, other nodes wait for a possible leader and begin to proceed in the candidate state by advancing to the next term unless the leader contacts them.

### Log Replication

#### Spec

* Pre-log replication: Once a leader is elected, it starts servicing the client. The leader appends a new request to its `New Entry` log then issues `AppendEntriesRPC` in parallel to all its peers. 
* Successful log: When all logs have been applied successfully to all follower machines, the leader applies the entry to its state machine and returns the result to the client.
* Repeating `AppendEntries`: `AppendEntriesRPC` are repeated indefinitely until all followers eventually store all log entries.
* Log entry storage: Log entries are a queue of state machine commands which are applied to that particular state machine. Log entries are associated with a term number to indicate the term of application of that log along with an integer index to identify a particular logs position.
* A committed entry: When a leader decides that the log entry is safe to apply to other state machines, that entry is called committed. All committed entries are durable and _will eventually be executed_ by all state machines.
* An entry -> Committed entry: A log entry is called committed once its replicated on the majority of the servers in the cluster. Once an entry is committed, it commits all the previous entries in the leaders log, including the entries created by the previous leaders.
* The  leader keeps track of the highest known index that it knows is committed and it is included in all the future `AppendEntriesRPC` (including heartbeats) to inform other servers.
* Theres some issue about log committing - "A log entry is committed once the leader that createdthe entry has replicated it on a majority of the servers" and " Once a follower learns that a log entry is committed, it applies theentry to its local state machine (in log order)." are not clear whether replicating and applying to state machine are the same. If they are its kind of a contradiction, else "aplication" can mean executing the STMT in the DB in our case.
* Log matching property: 
  * If two entries in different logs have the same index and term, then they store the same com-mand.
  * If two entries in different logs have the same index and term, then the logs are identical in allpreceding entries.
* When sending an AppendEntriesRPC, the leader includes the index and term of the entry in its log that immediately precedes thenew entries. If the follower does not find an entry in its log with the same index and term, then itrefuses the new entries. This helps in log matching. Which implies, a successful `AppendEntries` RPC means a synced log.
* Leader crashes inducing inconsistencies in logs: In Raft, the leader handles inconsistencies by forcing the followers’ logs to duplicate its own (the leader's). To do so, the leader must find the latest log entry where  the  two  logs  agree,  delete  any  entries  in  the  follower’s  log  after  that  point,  and  send  the follower all of the leader’s entries after that point. All of these actions happen in response to the consistency check performed by AppendEntries RPCs. Meaning, the leader checks for the consistency by maintaining a `nextIndex` value and dropping it down and sending `AppendEntriesRPC` (which does the consistency check and fails unless they're same) until a success is returned. These "check `AppendEntries` can be empty to save BandWidth(BW). The follower can also help here by sending the smallest agreeing index in the first RPC instead of the leader probing until it reaches the index. 
* Leader append-only property: A leader never overwrites or deletes entries in its own log.

#### Implementation

### Safety

This module ensures that the above protocol runs as expected, eliminating the corner cases.

#### Spec

* Election restriction: Raft uses the voting process to prevent a candidate from winning an election unless its log contains all committed entries. A candidate must contact a majority of the cluster in order to be elected, which means that every committed entry must be present in at least one of those servers. If the candidate’s log is at least as up-to-date as any other log in that majority, then it will hold all the committed entries. The `RequestVote` RPC implements thisrestriction: the RPC includes information about the candidate’s log, and the voter denies its vote if its own log is more up-to-date than that of the candidate.
* The decision of _which is the more updated_ log is based on the index and the term of the last log. Term gets priority, if terms are same, index is checked for precedence.
* Committing from previous term: Raft never commits log entries from previous terms by counting replicas. Only log entries from the leader’s current term are committed bycounting replicas; once an entry from the current term has been committed in this way, then all prior entries are committed indirectly because of the Log Matching Property.
* Follower or Candidate crashes: `RequestVotes` or `AppendEntries` RPC are tried indefinetely. Raft RPC are _idempotent_ making the retries harmless. (Possible judgement call: Should I wait for the results of the votes to base it for my decision in the election?)
* `Current term`, `last applied index` and `vote` that was casted along with the logs must be persisted in case of server crashes.

#### Implementation

#### Client interaction:
* Idempotency is maintained by having a unique client ID and the request ID. The same request by the same client cannot be requested twice, we assume here that the client didn't receive the responses due to network errors etc. Each server maintains a _session_ for each client. The session tracks the latest serial number processed for the client, along with the associated response. If a server receives a command whose serial number has already been executed, it responds immediately without re-executing the request.
*  With each request, the client includes the lowest sequencenumber for which it has not yet received a response, and the state machine then discards all responsesfor lower sequence numbers. Quite similar to TCP.
* The session for a client are _open on all nodes_. Session expiry happens on all nodes at once and in a deterministic way. LRU or an upper bound for sessions can be used. Some sort of timely probing is done to remove stale sessions.
* Live clients issue keep-alive requests during periods of inactivity, which are also augmented withthe leader’s timestamp and committed to the Raft log, in order to maintain their sessions.
* Reads can bypass the Raft log only if: 
  * If the leader has not yet marked an entry from its current term committed, it waits until ithas done so. The Leader Completeness Property guarantees that a leader has all committed entries, but at the start of its term, it may not know which those are. To find out, it needs to commit an entry from its term. Raft handles this by having each leader commit a blank `no-op` entry into the log at the start of its term. As soon as this no-op entry is committed, the leader’scommit index will be at least as large as any other servers’ during its term.
  * The leader saves its current commit index in a local variable `readIndex`. This will be used as a lower bound for the version of the state that the query operates against.
  * The leader needs to make sure it hasn’t been superseded by a newer leader of which it is unaware. It issues a new round of heartbeats and waits for their acknowledgments from amajority of the cluster. Once these acknowledgments are received, the leader knows that there could not have existed a leader for a greater term at the moment it sent the heartbeats. Thus, the readIndex was, at the time, the largest commit index ever seen by any server in the cluster.
  * The leader waits for its state machine to advance at least as far as the readIndex; this is current enough to satisfy linearizability
  * Finally, the leader issues the query against its state machine and replies to the client with the results.


## A strict testing mechanism

The testing mechanism to be implemented will enable us in figuring out the problems existing in the implementation leading to a more resilient system.
We have to test for the following basic failures:
1. Network partitioning.
2. Un-responsive peers.
3. Overloading peer.
4. Corruption of data in transit.

## Graceful handling of failures

Accepting failures exist and handling them gracefully enables creation of more resilient and stable systems. Having _circuit breakers_, _backoff mechanisms in clients_ and _validation and coordination mechanisms_ are some of the pointers to be followed. 

## Running Lbadd on Raft

* Background: Raft is just a consensus protocol that helps keep different database servers in sync. We need methods to issue a command and enable the sync between servers.
* Logistics: The `AppendEntriesRPC` will have the command to be executed by the client. This command goes through the leader, is applied by all the followers and then committed by the leader. Thus ensuring an in-sync distributed database.
* Each server will independently run the SQL statement once the statement is committed. This ensures that the state of the database is in sync. The data that moves around in the raft cluster can be a compiled SQL statement that each node will run independently.


## Appendix

* The difference between _commit_, _replicate_ and _apply_ with respect raft: What I have understood till now is, applying means letting the log run through the node's state machine. This is the end process, happens after a commit. A commit is once replication happens on a majority of the nodes. While replication is simply appending of a log on _one_ node. 
* Some gotchas I thought about: 
  * Client connects to the leader and leader crashes -> reject the connection. Let the client connect when the new leader is established.
  * Some sort of idempotency must be maintained w.r.t. the client-cluster communication. Multiple requests submitted by the client should not cause problems due to network errors.
