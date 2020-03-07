package sm_yamux

import (
	"testing"

	tmux "github.com/libp2p/go-libp2p-testing/suites/mux"
)

func TestDefaultTransport(t *testing.T) {
	tmux.SubtestAll(t, DefaultTransport)
}
