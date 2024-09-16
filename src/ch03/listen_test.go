package ch03

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listener.Close() }()

	t.Logf("bound to %q", listener.Addr())
}

func TestListenerAccept(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")

	if err != nil {
		t.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
		}

		go func(c net.Conn) {
			defer c.Close()

			t.Logf("conn: %v\n", conn)
		}(conn)
	}
}
