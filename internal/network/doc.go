// Package network implements a communication layer with a server and a client.
// Clients can connect to servers and receive and send messages. Depending on
// the implementation, communication may differ.
//
//  srv := network.NewTCPServer(zerolog.Nop()) // or any other available server
//  srv.OnConnect(handleConnection)
//  if err := srv.Open(":3900"); err != nil {
//      panic(err)
//  }
//
// In the above example, handleConnection is a func that accepts a network.Conn
// as parameter. Do with this connection what you'd like. The connections have
// an ID.
//
//  func handleConnection(conn network.Conn) {
//      loginMsg, err := conn.Receive() // receive a message
//      // handle loginMsg and err
//      // create loginResponse
//      err = conn.Send(loginResponse) // send a message
//      // handle err
//      connectionPool.Add(conn) // remember the connection for further use
//  }
//
// To connect to the above server, do as follows.
//
//  conn, err := network.DialTCP(":3900") // or any other available dial method
//  // handle err
//  defer conn.Close()
//  err = conn.Send(loginMsg) // send a message
//  // handle err
//  loginResponse, err := conn.Receive() // receive a message
//
// Please note, that the dial functions will only work with the respective
// server, e.g. DialTCP will only work on TCPServers.
package network
