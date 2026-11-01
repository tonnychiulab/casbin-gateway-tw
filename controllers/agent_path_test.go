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

package controllers

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsRestrictedSystemPath(t *testing.T) {
	testCases := []struct {
		path       string
		restricted bool
	}{
		{"/", true},
		{"/etc", true},
		{"/etc/passwd", true},
		{`"/etc"`, true},
		{`"/etc/shadow"`, true},
		{"/sys", true},
		{"/proc", true},
		{"/dev", true},
		{"/root", true},
		{"/boot", true},
		{"/var/run", true},
		{"/run", true},
		{"", false},
	}

	for _, tc := range testCases {
		got := isRestrictedSystemPath(tc.path)
		if got != tc.restricted {
			t.Errorf("isRestrictedSystemPath(%q) = %v; want %v", tc.path, got, tc.restricted)
		}
	}

	// Test traversal relative paths that resolve to /etc
	wd, err := os.Getwd()
	if err == nil {
		relToEtc, err := filepath.Rel(wd, "/etc")
		if err == nil {
			if !isRestrictedSystemPath(relToEtc) {
				t.Errorf("isRestrictedSystemPath(%q) should be restricted", relToEtc)
			}
			quotedRel := `"` + relToEtc + `"`
			if !isRestrictedSystemPath(quotedRel) {
				t.Errorf("isRestrictedSystemPath(%q) should be restricted", quotedRel)
			}
		}
	}
}
