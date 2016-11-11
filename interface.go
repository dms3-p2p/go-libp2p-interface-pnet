package pnet

import (
	"errors"
	"os"

	iconn "github.com/libp2p/go-libp2p-interface-conn"
)

const envKey = "LIBP2P_FORCE_PNET"

var ForcePrivateNetwork bool = false

func init() {
	ForcePrivateNetwork = os.Getenv(envKey) == "1"
}

var ErrNotInPrivateNetwork = errors.New("private network was not configured but" +
	" is enforced by the environment")

type Protector interface {
	Protect(iconn.Conn) (iconn.Conn, error)
}
