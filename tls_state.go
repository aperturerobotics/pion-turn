// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build !tinygo

package turn

import (
	"context"
	"crypto/tls"
	"net"
	"time"
)

func connectionTLSState(conn net.Conn) (*tls.ConnectionState, error) {
	tlsConn, ok := conn.(*tls.Conn)
	if !ok {
		return nil, nil
	}

	// Force TLS handshake to complete before extracting connection state.
	// Per crypto/tls docs: handshakes are lazy and complete on first I/O.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := tlsConn.HandshakeContext(ctx); err != nil {
		return nil, err
	}
	cs := tlsConn.ConnectionState()

	return &cs, nil
}
