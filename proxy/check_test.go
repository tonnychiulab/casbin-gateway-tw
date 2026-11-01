// Copyright 2026 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proxy

import (
	"net"
	"testing"
)

func TestIsRestrictedIP(t *testing.T) {
	testCases := []struct {
		ip         string
		restricted bool
	}{
		{"127.0.0.1", true},
		{"::1", true},
		{"0.0.0.0", true},
		{"::", true},
		{"10.0.0.1", true},
		{"172.16.0.1", true},
		{"192.168.1.1", true},
		{"169.254.169.254", true},
		{"8.8.8.8", false},
		{"1.1.1.1", false},
	}

	for _, tc := range testCases {
		parsed := net.ParseIP(tc.ip)
		if parsed == nil {
			t.Fatalf("failed to parse IP: %s", tc.ip)
		}
		got := isRestrictedIP(parsed)
		if got != tc.restricted {
			t.Errorf("isRestrictedIP(%s) = %v; want %v", tc.ip, got, tc.restricted)
		}
	}
}

func TestIsPrivateOrLoopbackHost(t *testing.T) {
	restricted, _ := isPrivateOrLoopbackHost("127.0.0.1")
	if !restricted {
		t.Errorf("expected 127.0.0.1 to be restricted")
	}

	restricted, _ = isPrivateOrLoopbackHost("0.0.0.0")
	if !restricted {
		t.Errorf("expected 0.0.0.0 to be restricted")
	}

	restricted, _ = isPrivateOrLoopbackHost("169.254.169.254")
	if !restricted {
		t.Errorf("expected 169.254.169.254 to be restricted")
	}
}
