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
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/apache/casbin-gateway/auth"
	"github.com/apache/casbin-gateway/conf"
	"github.com/apache/casbin-gateway/object"
	beegoContext "github.com/beego/beego/context"
	"github.com/beego/beego/session"
)

func newTestControllerWithSession(user *auth.User) (*ApiController, *httptest.ResponseRecorder) {
	recorder := httptest.NewRecorder()
	ctx := beegoContext.NewContext()
	ctx.Reset(recorder, httptest.NewRequest(http.MethodGet, "/api/get-account", nil))

	manager, _ := session.NewManager("memory", &session.ManagerConfig{
		CookieName:      "test_session",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
	})
	store, _ := manager.SessionStart(recorder, ctx.Request)

	c := &ApiController{}
	c.Ctx = ctx
	c.CruSession = store
	c.Data = make(map[interface{}]interface{})

	if user != nil {
		claims := auth.Claims{User: *user}
		store.Set("user", claims)
	}
	return c, recorder
}

func initTestDb(t *testing.T) {
	t.Helper()
	dbPath := filepath.Join(t.TempDir(), "test.db")
	conf.SetSettingOverrides(map[string]string{
		"driverName":      "sqlite",
		"dataSourceName":  dbPath,
		"dbName":          "",
		"casdoorEndpoint": "",
	})
	object.InitAdapter()
	object.CreateTables()
}

func TestRequireAdminBlocksDefaultPassword(t *testing.T) {
	initTestDb(t)
	object.InitUsers()

	user, err := object.GetUser("admin")
	if err != nil || user == nil {
		t.Fatalf("failed to retrieve seeded admin: %v", err)
	}

	if !object.IsAdminUsingDefaultPassword() {
		t.Fatal("seeded admin must initially have default password")
	}

	casdoorUser := user.ToCasdoorUser()
	c, _ := newTestControllerWithSession(&casdoorUser)

	// Cycle 1 Red test: RequireAdmin & RequireSignedIn must block when admin uses default password
	if !c.RequireSignedIn() {
		t.Fatal("RequireSignedIn() should block admin who is using default password")
	}

	cAdmin, recorderAdmin := newTestControllerWithSession(&casdoorUser)
	blocked := cAdmin.RequireAdmin()
	if !blocked {
		t.Fatal("RequireAdmin() should block admin who is using default password")
	}

	var resp Response
	if err := json.Unmarshal(recorderAdmin.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if resp.Status != "error" || resp.Msg != "please change default password first" {
		t.Fatalf("expected error 'please change default password first', got status=%q, msg=%q", resp.Status, resp.Msg)
	}

	// Update password to non-default
	if err := object.UpdateUserPassword(user, "NewStrongPassword123!"); err != nil {
		t.Fatalf("failed to update password: %v", err)
	}

	// Now RequireAdmin and RequireSignedIn should NOT block
	c2, _ := newTestControllerWithSession(&casdoorUser)
	if c2.RequireSignedIn() {
		t.Fatal("RequireSignedIn() should allow admin after password is changed")
	}
	if c2.RequireAdmin() {
		t.Fatal("RequireAdmin() should allow admin after password is changed")
	}
}
