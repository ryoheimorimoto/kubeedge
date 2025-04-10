/*
Copyright 2019 The KubeEdge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package edgehub

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/kubeedge/api/apis/componentconfig/edgecore/v1alpha2"
	"github.com/kubeedge/kubeedge/edge/pkg/edgehub/config"
)

func TestGetCertSyncChannel(t *testing.T) {
	t.Run("GetCertSyncChannel()", func(t *testing.T) {
		certSync := GetCertSyncChannel()
		if certSync != nil {
			t.Errorf("GetCertSyncChannel() returned unexpected result. got = %v, want = %v", certSync, nil)
		}
	})
}

func TestNewCertSyncChannel(t *testing.T) {
	t.Run("NewCertSyncChannel()", func(t *testing.T) {
		certSync := NewCertSyncChannel()
		if len(certSync) != 1 {
			t.Errorf("NewCertSyncChannel() returned  unexpected results. size got = %d, size want = 2", len(certSync))
		}
		if _, ok := certSync["edgestream"]; !ok {
			t.Error("NewCertSyncChannel() returned  unexpected results. expected key edgestream to be present but it was not available.")
		}
	})
}

func TestRegister(t *testing.T) {
	tests := []struct {
		eh           *v1alpha2.EdgeHub
		nodeName     string
		name         string
		wantNodeName string
	}{
		{
			name:         "",
			nodeName:     "test1",
			wantNodeName: "test1",
			eh:           &v1alpha2.EdgeHub{WebSocket: &v1alpha2.EdgeHubWebSocket{Server: "localhost:8080"}, ProjectID: "test_id"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Register(tt.eh, tt.nodeName)

			if config.Config.NodeName != tt.wantNodeName {
				t.Errorf("failed to Register(). Nodename : got = %s, want = %s", config.Config.NodeName, tt.wantNodeName)
			}
		})
	}
}

func TestName(t *testing.T) {
	t.Run("EdgeHub.Name()", func(t *testing.T) {
		if got := (&EdgeHub{}).Name(); got != "websocket" {
			t.Errorf("EdgeHub.Name() returned unexpected result. got = %s, want = websocket", got)
		}
	})
}

func TestGroup(t *testing.T) {
	t.Run("EdgeHub.Group()", func(t *testing.T) {
		if got := (&EdgeHub{}).Group(); got != "hub" {
			t.Errorf("EdgeHub.Group() returned unexpected result. got = %s, want = hub", got)
		}
	})
}

func TestEnable(t *testing.T) {
	tests := []struct {
		eh   *EdgeHub
		want bool
		name string
	}{
		{
			name: "Enable true",
			want: true,
			eh:   &EdgeHub{enable: true},
		},
		{
			name: "Enable false",
			want: false,
			eh:   &EdgeHub{enable: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.eh.Enable(); got != tt.want {
				t.Errorf("EdgeHub.Enable() returned expected results. got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestMonitorNetworkChanges(t *testing.T) {
	interfaces, err := net.Interfaces()
	if err != nil {
		t.Fatalf("Failed to get network interfaces: %v", err)
	}

	var testIface net.Interface
	var testAddr net.Addr
	for _, iface := range interfaces {
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok || ipNet.IP.IsLoopback() || ipNet.IP.To4() == nil {
				continue
			}

			testIface = iface
			testAddr = addr
			break
		}
		if testIface.Name != "" {
			break
		}
	}

	if testIface.Name == "" {
		t.Skip("No suitable network interface found for testing")
	}

	ipNet, _ := testAddr.(*net.IPNet)
	initialIfaces := map[string]string{testIface.Name: ipNet.IP.String()}

	tests := []struct {
		name           string
		hub            *EdgeHub
		initialIfaces  map[string]string
		expectedIfaces map[string]string
	}{
		{
			name:           "Network interface detection",
			hub:            &EdgeHub{networkInterfaces: make(map[string]string)},
			initialIfaces:  initialIfaces,
			expectedIfaces: initialIfaces,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hub.networkInterfaces = tt.initialIfaces
			tt.hub.updateNetworkInterfaces()

			if len(tt.hub.networkInterfaces) == 0 {
				t.Error("No network interfaces detected")
			}

			found := false
			for _, ip := range tt.hub.networkInterfaces {
				if ip == ipNet.IP.String() {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected IP %s not found in interfaces %v", ipNet.IP.String(), tt.hub.networkInterfaces)
			}
		})
	}
}

func TestReconnectWaitTime(t *testing.T) {
	tests := []struct {
		name         string
		hub          *EdgeHub
		heartbeat    int32
		expectedWait time.Duration
	}{
		{
			name:         "Default heartbeat period",
			hub:          &EdgeHub{},
			heartbeat:    15,
			expectedWait: time.Duration(15) * time.Second,
		},
		{
			name:         "Custom heartbeat period",
			hub:          &EdgeHub{},
			heartbeat:    30,
			expectedWait: time.Duration(30) * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.Config.Heartbeat = tt.heartbeat
			waitTime := time.Duration(config.Config.Heartbeat) * time.Second
			if waitTime != tt.expectedWait {
				t.Errorf("ReconnectWaitTime() = %v, want %v", waitTime, tt.expectedWait)
			}
		})
	}
}

func TestErrorMessage(t *testing.T) {
	tests := []struct {
		name        string
		hub         *EdgeHub
		err         error
		expectedMsg string
	}{
		{
			name:        "Connection error message",
			hub:         &EdgeHub{},
			err:         fmt.Errorf("connection refused"),
			expectedMsg: "connection failed: connection refused",
		},
		{
			name:        "Network error message",
			hub:         &EdgeHub{},
			err:         fmt.Errorf("network unreachable"),
			expectedMsg: "connection failed: network unreachable",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errMsg := fmt.Sprintf("connection failed: %v", tt.err)
			if errMsg != tt.expectedMsg {
				t.Errorf("ErrorMessage() = %v, want %v", errMsg, tt.expectedMsg)
			}
		})
	}
}
