/*
Copyright 2026 The KubeEdge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0
*/

package conn

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"

	"github.com/kubeedge/beehive/pkg/core/model"
	"github.com/kubeedge/kubeedge/pkg/viaduct/pkg/api"
	"github.com/kubeedge/kubeedge/pkg/viaduct/pkg/fifo"
	"github.com/kubeedge/kubeedge/pkg/viaduct/pkg/keeper"
)

// newTestWSConn spins up a server-side websocket that does nothing and
// connects a client to it. The returned WSConnection is configured with the
// supplied read deadline interval; the caller must close srv to release the
// goroutine.
func newTestWSConn(t *testing.T, readDeadlineInterval time.Duration) (*WSConnection, *httptest.Server) {
	t.Helper()

	upgrader := websocket.Upgrader{
		CheckOrigin: func(*http.Request) bool { return true },
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Errorf("upgrade: %v", err)
			return
		}
		defer c.Close()
		// Hold the connection without sending anything so the client side
		// blocks on Read until either the deadline fires or we close.
		for {
			if _, _, err := c.ReadMessage(); err != nil {
				return
			}
		}
	}))

	wsURL := "ws" + strings.TrimPrefix(srv.URL, "http")
	wsConn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		srv.Close()
		t.Fatalf("dial: %v", err)
	}

	conn := &WSConnection{
		wsConn:               wsConn,
		state:                &ConnectionState{State: api.StatConnected, Headers: http.Header{}},
		connUse:              api.UseTypeMessage,
		messageFifo:          fifo.NewMessageFifo(),
		syncKeeper:           keeper.NewSyncKeeper(),
		readDeadlineInterval: readDeadlineInterval,
	}
	return conn, srv
}

// TestHandleMessageReadDeadlineFiresWithinInterval verifies the S1 fix:
// when readDeadlineInterval is set, handleMessage exits within roughly that
// interval even though the peer never sends anything. Without the fix the
// goroutine would block until kernel TCP retransmission timeout (~15min).
func TestHandleMessageReadDeadlineFiresWithinInterval(t *testing.T) {
	conn, srv := newTestWSConn(t, 100*time.Millisecond)
	defer srv.Close()

	done := make(chan struct{})
	go func() {
		conn.handleMessage()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("handleMessage did not return within 2s; deadline not applied")
	}

	// messageFifo must be closed so callers blocked on Get() observe the
	// error immediately rather than waiting for the next keepalive failure.
	msg := &model.Message{}
	if err := conn.ReadMessage(msg); err == nil {
		t.Fatal("expected ReadMessage to return error after handleMessage timeout")
	}
}

// TestHandleMessageZeroReadDeadlineKeepsLegacyBehavior verifies that the
// existing zero-value behavior (no deadline = block forever) is preserved.
func TestHandleMessageZeroReadDeadlineKeepsLegacyBehavior(t *testing.T) {
	conn, srv := newTestWSConn(t, 0)
	defer srv.Close()

	done := make(chan struct{})
	go func() {
		conn.handleMessage()
		close(done)
	}()

	select {
	case <-done:
		t.Fatal("handleMessage returned unexpectedly with no read deadline")
	case <-time.After(300 * time.Millisecond):
		// expected: blocks indefinitely
	}

	// Cleanup: close the underlying conn to unblock the goroutine.
	_ = conn.wsConn.Close()
	<-done
}

// TestSetReadDeadlinePropagatesToWSConn verifies the SetReadDeadline bug
// fix: the call must reach gorilla/websocket's underlying conn. We trigger
// this by setting a past deadline and observing that the next read errors
// out immediately.
func TestSetReadDeadlinePropagatesToWSConn(t *testing.T) {
	conn, srv := newTestWSConn(t, 0)
	defer srv.Close()
	// Close the client conn explicitly so the server-side handler goroutine
	// (looping on ReadMessage) exits even if srv.Close alone does not
	// terminate the hijacked websocket connection.
	defer conn.wsConn.Close()

	if err := conn.SetReadDeadline(time.Now().Add(-time.Second)); err != nil {
		t.Fatalf("SetReadDeadline: %v", err)
	}

	if _, _, err := conn.wsConn.ReadMessage(); err == nil {
		t.Fatal("expected immediate read error after past deadline; SetReadDeadline did not propagate")
	}
}
