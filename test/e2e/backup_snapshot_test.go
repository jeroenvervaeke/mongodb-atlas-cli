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
//go:build e2e || (atlas && backup && snapshot)

package e2e_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func TestSnapshots(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	cliPath, err := AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	clusterName := g.memory("clusterName", must(RandClusterName())).(string)

	mdbVersion, err := MongoDBMajorVersion()
	r.NoError(err)

	var snapshotID string

	g.Run("Create cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"create",
			clusterName,
			"--backup",
			"--tier", tierM10,
			"--region=US_EAST_1",
			"--provider", e2eClusterProvider,
			"--mdbVersion", mdbVersion,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var cluster *atlasClustersPinned.AdvancedClusterDescription
		require.NoError(t, json.Unmarshal(resp, &cluster))
		ensureCluster(t, cluster, clusterName, mdbVersion, 10, false)
	})
	t.Cleanup(func() {
		require.NoError(t, deleteClusterForProject("", clusterName))
	})
	require.NoError(t, watchCluster("", clusterName))

	g.Run("Create snapshot", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"create",
			clusterName,
			"--desc",
			"test-snapshot",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))

		var snapshot atlasv2.DiskBackupSnapshot
		require.NoError(t, json.Unmarshal(resp, &snapshot))
		assert.Equal(t, "test-snapshot", snapshot.GetDescription())
		snapshotID = snapshot.GetId()
	})

	g.Run("Watch creation", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"watch",
			snapshotID,
			"--clusterName",
			clusterName)
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"list",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
		var backups atlasv2.PaginatedCloudBackupReplicaSet
		require.NoError(t, json.Unmarshal(resp, &backups))
		assert.NotEmpty(t, backups)
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"describe",
			snapshotID,
			"--clusterName",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))

		var result atlasv2.DiskBackupReplicaSet
		require.NoError(t, json.Unmarshal(resp, &result))
		assert.Equal(t, snapshotID, result.GetId())
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"delete",
			snapshotID,
			"--clusterName",
			clusterName,
			"--force")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	if skipCleanup() {
		return
	}

	g.Run("Watch deletion", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"watch",
			snapshotID,
			"--clusterName",
			clusterName)
		cmd.Env = os.Environ()
		resp, _ := RunAndGetStdOut(cmd)
		t.Log(string(resp))
	})
}
