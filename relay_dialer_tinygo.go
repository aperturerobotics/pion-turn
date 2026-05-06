// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build tinygo

package turn

import (
	"net"

	"github.com/pion/transport/v4"
)

func relayDialer(n transport.Net, local net.Addr) transport.Dialer {
	return n.CreateDialer(&net.Dialer{
		LocalAddr: local,
	})
}

func relayListenConfig(n transport.Net) transport.ListenConfig {
	return n.CreateListenConfig(&net.ListenConfig{})
}
