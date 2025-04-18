// Copyright 2020 MongoDB Inc
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
//go:build e2e || (atlas && processes)

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

func TestProcesses(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	g.generateProjectAndCluster("processes")

	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)

	var processes *atlasv2.PaginatedHostViewAtlas

	g.Run("list", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			processesEntity,
			"list",
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		require.NoError(t, json.Unmarshal(resp, &processes))
		require.NotEmpty(t, processes.Results)
	})

	g.Run("list compact", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			processesEntity,
			"list",
			"-c",
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var hostViewsCompact []atlasv2.ApiHostViewAtlas
		require.NoError(t, json.Unmarshal(resp, &hostViewsCompact))
		require.NotEmpty(t, hostViewsCompact)
	})

	g.Run("describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			processesEntity,
			"describe",
			processes.GetResults()[0].GetId(),
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var p *atlasv2.ApiHostViewAtlas
		require.NoError(t, json.Unmarshal(resp, &p))
		assert.Equal(t, *p.Id, *processes.GetResults()[0].Id)
	})
}
