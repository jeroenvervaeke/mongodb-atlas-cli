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
//go:build e2e || (atlas && performanceAdvisor)

package e2e_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPerformanceAdvisor(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot(), withSnapshotNameFunc(snapshotHashedName))
	g.generateProjectAndCluster("performanceAdvisor")

	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)

	hostname, err := g.getHostnameAndPort()
	require.NoError(t, err)

	g.Run("List namespaces", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			performanceAdvisorEntity,
			namespacesEntity,
			"list",
			"--processName", hostname,
			"--projectId", g.projectID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("List slow query logs", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			performanceAdvisorEntity,
			slowQueryLogsEntity,
			"list",
			"--processName", hostname,
			"--projectId", g.projectID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("List suggested indexes", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			performanceAdvisorEntity,
			suggestedIndexesEntity,
			"list",
			"--processName", hostname,
			"--projectId", g.projectID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("Enable Managed Slow Operation Threshold", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			performanceAdvisorEntity,
			slowOperationThresholdEntity,
			"enable",
			"--projectId", g.projectID,
		)

		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("Disable Managed Slow Operation Threshold", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			performanceAdvisorEntity,
			slowOperationThresholdEntity,
			"disable",
			"--projectId", g.projectID,
		)

		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})
}
