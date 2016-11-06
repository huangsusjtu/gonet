package network

import (
	"testing"
)

func TestConnection1(t *testing.T) {
	conn := newConnection(nil)
	conn.id = 123
}
