package cluster

import (
	"context"
	"net"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/network"
)

// TestNetwork encompasses the entire network on which
// the tests will be performed.
type TestNetwork struct {
	Clusters []*TestCluster
}

// TestCluster describes a single cluster and all the
// nodes in it.
type TestCluster struct {
	Cluster Cluster
	Nodes   []*network.TestNode
}

// NewTestCluster creates a cluster and joins all the nodes onto the cluster.
func NewTestCluster(nodes []*network.TestNode, log zerolog.Logger) (*TestCluster, error) {
	cluster := NewTCPCluster(log)
	err := joinCluster(nodes, cluster)
	if err != nil {
		return nil, err
	}
	return &TestCluster{
		Cluster: cluster,
		Nodes:   nodes,
	}, nil
}

// joinCluster allows a set of nodes to join a cluster.
func joinCluster(
	tcpServers []*network.TestNode,
	clstr Cluster,
) error {
	ctx := context.TODO()
	for i := 0; i < len(tcpServers); i++ {
		err := clstr.Join(ctx, tcpServers[i].Node.Addr().String())
		if err != nil {
			return err
		}
	}
	return nil
}

// NewTestNetwork returns a ready to use network.
//
// It creates the necessary amount of nodes, links them
// appropriately and then creates clusters based on the
// nodes.
func NewTestNetwork(number int, log zerolog.Logger) (*TestNetwork, error) {

	// Create the number of nodes needed.
	IP := "127.0.0.1"
	basePort := 12000

	var nodes []*network.TestNode
	for i := 0; i < number; i++ {
		port := basePort + i
		node, err := network.NewTestNode(IP, strconv.Itoa(port), log)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	linkNodes(nodes)

	// Using the nodes, group them into clusters.
	// Each node "sees" the other nodes as a cluster,
	// thus we have number of clusters equal to the nodes.
	var clusters []*TestCluster

	for i := 0; i < number; i++ {

		otherNodes := exclude(nodes, i)

		clstr, err := NewTestCluster(otherNodes, log)
		if err != nil {
			return nil, err
		}

		clusters = append(clusters, clstr)
	}
	return &TestNetwork{
		Clusters: clusters,
	}, nil
}

// exclude excludes the i'th element from the slice.
func exclude(s []*network.TestNode, i int) []*network.TestNode {
	l := len(s)

	rs := make([]*network.TestNode, l)
	copy(rs, s)

	return append(rs[:i], rs[i+1:]...)
}

// linkNodes links the network of nodes together with a
// net.Pipe connection over the network.Conn.
//
// TCP connections are created and each node keeps one end
// of the connection and another end of the pipe is added
// to the other nodes list.
func linkNodes(nodes []*network.TestNode) {
	for i := 0; i < len(nodes)-1; i++ {
		numConns := len(nodes) - i - 1
		tcpConnSelf, tcpConnOther := createTCPConns(numConns)
		nodes[i].Conns = append(nodes[i].Conns, tcpConnSelf...)
		for j := i + 1; j < len(nodes); j++ {
			nodes[j].Conns = append(nodes[j].Conns, tcpConnOther[j-i-1])
		}
	}
}

// createTCPConns creates and returns 2 ends of a pipe of a TCP connection.
func createTCPConns(count int) ([]network.Conn, []network.Conn) {
	var tcpSelfSlice, tcpOtherSlice []network.Conn
	for count > 0 {
		c1, c2 := net.Pipe()
		tcpSelf, tcpOther := network.NewTCPConn(c1), network.NewTCPConn(c2)
		tcpSelfSlice = append(tcpSelfSlice, tcpSelf)
		tcpOtherSlice = append(tcpOtherSlice, tcpOther)
		count--
	}
	return tcpSelfSlice, tcpOtherSlice
}
