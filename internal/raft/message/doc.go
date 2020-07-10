// Package message implements messages used in the raft module for
// communication. Create a message by calling constructor functions. An example
// follows
//
//  // need to create a request-vote-message and send it
//  msg := message.NewRequestVoteRequest(term, candidate.ID(), lastLogIndex, lastLogTerm) // create the message
//  data, err := message.Marshal(msg) // marshal it
//  // handle err
//  conn.Send(data) // sent it through the network
//
// When receiving data however, follow this example. In here, we will receive
// bytes, unmarshal them as a message, and then process the message.
//
//	data := conn.Receive()
//	msg, err := message.Unmarshal(data)
//	switch msg.Kind() {
//		case message.KindRequestVoteResponse:
//			// process a request-vote-response
//		default:
//			panic("cannot handle the message")
//	}
package message
