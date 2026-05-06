// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build !tinygo

package turn

import (
	"net"

	"github.com/pion/transport/v4"
	"github.com/pion/transport/v4/reuseport"
)

func relayDialer(n transport.Net, local net.Addr) transport.Dialer {
	return n.CreateDialer(&net.Dialer{
		LocalAddr: local,
		// Enable SO_REUSEADDR and SO_REUSEPORT where needed to let multiple connections
		// bind to the same relay address.
		Control: reuseport.Control,
	})
}

func relayListenConfig(n transport.Net) transport.ListenConfig {
	return n.CreateListenConfig(&net.ListenConfig{
		// Enable SO_REUSEADDR and SO_REUSEPORT where needed to let multiple connections
		// bind to the same relay address.
		Control: reuseport.Control,
	})
}
