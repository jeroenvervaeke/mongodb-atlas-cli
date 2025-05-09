// Copyright 2023 MongoDB Inc
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

//go:build e2e || (atlas && generic)

package e2e_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func TestAuditing(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	g.generateProject("auditing")
	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			auditingEntity,
			"describe",
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var setting *atlasv2.AuditLog
		require.NoError(t, json.Unmarshal(resp, &setting), string(resp))
	})

	g.Run("Update", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			auditingEntity,
			"update",
			"--projectId", g.projectID,
			"--enabled",
			"--auditAuthorizationSuccess",
			"--auditFilter", "{\"atype\": \"authenticate\"}",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var setting *atlasv2.AuditLog
		require.NoError(t, json.Unmarshal(resp, &setting), string(resp))
		assert.True(t, *setting.Enabled)
		assert.True(t, *setting.AuditAuthorizationSuccess)
		assert.JSONEq(t, "{\"atype\": \"authenticate\"}", *setting.AuditFilter)
	})

	g.Run("Update via file", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			auditingEntity,
			"update",
			"--projectId", g.projectID,
			"--enabled",
			"--auditAuthorizationSuccess",
			"-f", "data/update_auditing.json",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var setting *atlasv2.AuditLog
		require.NoError(t, json.Unmarshal(resp, &setting), string(resp))
	})
}
