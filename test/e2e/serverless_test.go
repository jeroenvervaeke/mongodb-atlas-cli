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
//go:build e2e || (atlas && serverless && instance)

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

func TestServerless(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	g.generateProject("serverless")

	cliPath, err := AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterName := g.memory("clusterName", must(RandClusterName())).(string)

	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			serverlessEntity,
			"create",
			clusterName,
			"--region=US_EAST_1",
			"--provider=AWS",
			"--tag", "env=test",
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster atlasv2.ServerlessInstanceDescription
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		a := assert.New(t)
		a.Equal(clusterName, *cluster.Name)
	})

	g.Run("Watch", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			serverlessEntity,
			"watch",
			"--projectId", g.projectID,
			clusterName,
		)
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		a := assert.New(t)
		a.Contains(string(resp), "Instance available")
	})

	g.Run("Update", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			serverlessEntity,
			"update",
			clusterName,
			"--disableTerminationProtection",
			"--tag", "env=e2e",
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster atlasv2.ServerlessInstanceDescription
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		a := assert.New(t)
		a.Equal(clusterName, *cluster.Name)
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			serverlessEntity,
			"ls",
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var clusters atlasv2.PaginatedServerlessInstanceDescription
		err = json.Unmarshal(resp, &clusters)
		req.NoError(err)

		a := assert.New(t)
		a.NotEmpty(clusters.Results)
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			serverlessEntity,
			"describe",
			clusterName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster atlasv2.ServerlessInstanceDescription
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		a := assert.New(t)
		a.Equal(clusterName, *cluster.Name)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(
			cliPath,
			serverlessEntity,
			"delete",
			clusterName,
			"--projectId", g.projectID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Serverless instance '%s' deleted\n", clusterName)
		a := assert.New(t)
		a.Equal(expected, string(resp))
	})

	if skipCleanup() {
		return
	}

	g.Run("Watch deletion", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			serverlessEntity,
			"watch",
			clusterName,
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		// this command will fail with 404 once the cluster is deleted
		// we just need to wait for this to close the project
		resp, _ := RunAndGetStdOut(cmd)
		t.Log(string(resp))
	})
}
