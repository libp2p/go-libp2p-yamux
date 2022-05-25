// Deprecated: This package has moved into go-libp2p as a sub-package: github.com/libp2p/go-libp2p/p2p/muxer/yamux.
package sm_yamux

import (
	"github.com/libp2p/go-libp2p/p2p/muxer/yamux"
)

// Deprecated: use github.com/libp2p/go-libp2p/p2p/muxer/yamux.DefaultTransport instead.
var DefaultTransport = yamux.DefaultTransport

// Transport implements mux.Multiplexer that constructs
// yamux-backed muxed connections.
// Deprecated: use github.com/libp2p/go-libp2p/p2p/muxer/yamux.Transport instead.
type Transport = yamux.Transport
