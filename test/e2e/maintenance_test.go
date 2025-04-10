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

func TestMaintenanceWindows(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	g.generateProject("maintenance")

	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)

	g.Run("update", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			maintenanceEntity,
			"update",
			"--dayOfWeek",
			"1",
			"--hourOfDay",
			"1",
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		expected := "Maintenance window updated.\n"
		assert.Equal(t, expected, string(resp))
	})

	g.Run("describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			maintenanceEntity,
			"describe",
			"-o",
			"json",
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var maintenanceWindow atlasv2.GroupMaintenanceWindow
		require.NoError(t, json.Unmarshal(resp, &maintenanceWindow))
		a := assert.New(t)
		a.Equal(1, maintenanceWindow.GetDayOfWeek())
		a.Equal(1, maintenanceWindow.GetHourOfDay())
	})

	g.Run("clear", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			maintenanceEntity,
			"clear",
			"--force",
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		expected := "Maintenance window removed.\n"
		assert.Equal(t, expected, string(resp))
	})
}
