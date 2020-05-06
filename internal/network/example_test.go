package network_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/network"
)

func ExampleServer() {
	// When using, please don't ignore all the errors as we do here.

	ctx := context.Background()

	srv := network.NewTCPServer(zerolog.Nop()) // or whatever server is available
	srv.OnConnect(func(conn network.Conn) {
		_ = conn.Send(ctx, []byte("Hello, World!"))
	})
	go func() {
		if err := srv.Open(":59513"); err != nil {
			log.Fatal(err)
		}
	}()

	time.Sleep(10 * time.Millisecond)
	client, _ := network.DialTCP(ctx, ":59513")
	defer func() {
		_ = client.Close()
	}()
	received, _ := client.Receive(ctx)
	fmt.Println(string(received))
	// Output: Hello, World!
}
