package ipnet

import (
	"net"
	"os"
)

// EnvKey defines environment variable name for forcing usage of PNet in dms3-p2p
// When environment variable of this name is set to "1" the ForcePrivateNetwork
// variable will be set to true.
const EnvKey = "LIBP2P_FORCE_PNET"

// ForcePrivateNetwork is boolean variable that forces usage of PNet in dms3-p2p
// Setting this variable to true or setting LIBP2P_FORCE_PNET environment variable
// to true will make dms3-p2p to require private network protector.
// If no network protector is provided and this variable is set to true dms3-p2p will
// refuse to connect.
var ForcePrivateNetwork = false

func init() {
	ForcePrivateNetwork = os.Getenv(EnvKey) == "1"
}

// ErrNotInPrivateNetwork is an error that should be returned by dms3-p2p when it
// tries to dial witt ForcePrivateNetwork set and no PNet Protector
var ErrNotInPrivateNetwork = NewError("private network was not configured but" +
	" is enforced by the environment")

// Protector interface is a way for private network implementation to be transparent in
// dms3-p2p. It is created by implementation and use by dms3-p2p-conn to secure connections
// so they can be only established with selected number of peers.
type Protector interface {
	// Wraps passed connection to protect it
	Protect(net.Conn) (net.Conn, error)

	// Returns key fingerprint that is safe to expose
	Fingerprint() []byte
}
