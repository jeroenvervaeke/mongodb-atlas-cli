// Copyright 2021 MongoDB Inc
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
//go:build e2e || (atlas && networking)

package e2e_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

var regionsAWS = []string{
	"us-east-1",
	"us-east-2",
	"us-west-1",
	"ca-central-1",
	"sa-east-1",
	"eu-west-1",
	"eu-central-1",
}

func TestPrivateEndpointsAWS(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	g.generateProject("privateEndpointsAWS")

	n := g.memoryRand("rand", int64(len(regionsAWS)))

	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)

	region := regionsAWS[n.Int64()]
	var id string

	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			awsEntity,
			"create",
			"--region",
			region,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var r atlasv2.EndpointService
		require.NoError(t, json.Unmarshal(resp, &r))
		id = r.GetId()
	})
	require.NotEmpty(t, id)

	g.Run("Watch", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			awsEntity,
			"watch",
			id,
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()

		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			awsEntity,
			"describe",
			id,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var r atlasv2.EndpointService
		require.NoError(t, json.Unmarshal(resp, &r))
		assert.Equal(t, id, r.GetId())
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			awsEntity,
			"ls",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var r []atlasv2.EndpointService
		require.NoError(t, json.Unmarshal(resp, &r))
		assert.NotEmpty(t, r)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			awsEntity,
			"delete",
			id,
			"--projectId",
			g.projectID,
			"--force")
		cmd.Env = os.Environ()

		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("Private endpoint '%s' deleted\n", id)
		assert.Equal(t, expected, string(resp))
	})

	if skipCleanup() {
		return
	}

	g.Run("Watch", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			awsEntity,
			"watch",
			id,
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		// We expect a 404 error once the private endpoint has been completely deleted
		require.Error(t, err, string(resp))
		assert.Contains(t, string(resp), "404")
	})
}

var regionsAzure = []string{
	"US_EAST_2",
	"EUROPE_NORTH",
	"US_WEST_2",
	"ASIA_SOUTH_EAST",
}

func TestPrivateEndpointsAzure(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	g.generateProject("privateEndpointsAzure")

	n := g.memoryRand("rand", int64(len(regionsAzure)))

	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)

	region := regionsAzure[n.Int64()]
	var id string

	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			azureEntity,
			"create",
			"--region",
			region,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()

		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var r atlasv2.EndpointService
		require.NoError(t, json.Unmarshal(resp, &r))
		id = r.GetId()
	})
	if id == "" {
		assert.FailNow(t, "Failed to create private endpoint")
	}

	g.Run("Watch", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			azureEntity,
			"watch",
			id,
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()
		_, err := RunAndGetStdOut(cmd)
		require.NoError(t, err)
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			azureEntity,
			"describe",
			id,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var r atlasv2.EndpointService
		require.NoError(t, json.Unmarshal(resp, &r))
		assert.Equal(t, id, r.GetId())
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			azureEntity,
			"ls",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var r []atlasv2.EndpointService
		require.NoError(t, json.Unmarshal(resp, &r))
		assert.NotEmpty(t, r)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			azureEntity,
			"delete",
			id,
			"--force",
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()

		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("Private endpoint '%s' deleted\n", id)
		assert.Equal(t, expected, string(resp))
	})

	if skipCleanup() {
		return
	}

	g.Run("Watch", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			azureEntity,
			"watch",
			id,
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		// We expect a 404 error once the private endpoint has been completely deleted
		require.Error(t, err)
		assert.Contains(t, string(resp), "404")
	})
}

var regionsGCP = []string{
	"CENTRAL_US",
	"US_EAST_4",
	"NORTH_AMERICA_NORTHEAST_1",
	"SOUTH_AMERICA_EAST_1",
	"WESTERN_US",
	"US_WEST_2",
	"AUSTRALIA_SOUTHEAST_2",
	"ASIA_SOUTHEAST_2",
	"WESTERN_EUROPE",
	"EUROPE_NORTH_1",
	"EUROPE_WEST_2",
	"EUROPE_CENTRAL_2",
}

func TestPrivateEndpointsGCP(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	g.generateProject("privateEndpointsGPC")

	n := g.memoryRand("rand", int64(len(regionsGCP)))

	region := regionsGCP[n.Int64()]

	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)
	var id string

	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			gcpEntity,
			"create",
			"--region="+region,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var r atlasv2.EndpointService
		require.NoError(t, json.Unmarshal(resp, &r))
		id = r.GetId()
		assert.NotEmpty(t, id)
	})

	g.Run("Watch", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			gcpEntity,
			"watch",
			id,
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()

		_, err := RunAndGetStdOut(cmd)
		require.NoError(t, err)
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			gcpEntity,
			"describe",
			id,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var r atlasv2.EndpointService
		require.NoError(t, json.Unmarshal(resp, &r))
		assert.Equal(t, id, r.GetId())
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			gcpEntity,
			"ls",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var r []atlasv2.EndpointService
		require.NoError(t, json.Unmarshal(resp, &r))
		assert.NotEmpty(t, r)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			gcpEntity,
			"delete",
			id,
			"--force",
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()

		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("Private endpoint '%s' deleted\n", id)
		assert.Equal(t, expected, string(resp))
	})

	if skipCleanup() {
		return
	}

	g.Run("Watch", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			gcpEntity,
			"watch",
			id,
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		// We expect a 404 error once the private endpoint has been completely deleted
		require.Error(t, err)
		assert.Contains(t, string(resp), "404")
	})
}

func TestRegionalizedPrivateEndpointsSettings(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	g.generateProject("regionalizedPrivateEndpointsSettings")

	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)

	g.Run("Enable regionalized private endpoint setting", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			regionalModeEntity,
			"enable",
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()

		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		assert.Equal(t, "Regionalized private endpoint setting enabled.\n", string(resp))
	})

	g.Run("Disable regionalized private endpoint setting", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			regionalModeEntity,
			"disable",
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()

		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		assert.Equal(t, "Regionalized private endpoint setting disabled.\n", string(resp))
	})

	g.Run("Get regionalized private endpoint setting", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			privateEndpointsEntity,
			regionalModeEntity,
			"get",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()

		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var r atlasv2.ProjectSettingItem
		require.NoError(t, json.Unmarshal(resp, &r))
		assert.False(t, r.GetEnabled())
	})
}
