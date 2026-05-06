// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build tinygo

package turn

import (
	"crypto/tls"
	"net"
)

func connectionTLSState(net.Conn) (*tls.ConnectionState, error) {
	return nil, nil
}
