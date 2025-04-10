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

//go:build e2e || (iam && atlas)

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

func TestAtlasOrgAPIKeyAccessList(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	cliPath, er := AtlasCLIBin()
	require.NoError(t, er)

	apiKeyID, e := createOrgAPIKey()
	require.NoError(t, e)

	t.Cleanup(func() {
		require.NoError(t, deleteOrgAPIKey(apiKeyID))
	})

	n := g.memoryRand("rand", 255)
	entry := fmt.Sprintf("192.168.0.%d", n)

	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			apiKeyAccessListEntity,
			"create",
			"--apiKey",
			apiKeyID,
			"--ip",
			entry,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var key atlasv2.PaginatedApiUserAccessListResponse
		require.NoError(t, json.Unmarshal(resp, &key))
		assert.NotEmpty(t, key.Results)
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			apiKeyAccessListEntity,
			"list",
			apiKeyID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var key atlasv2.PaginatedApiUserAccessListResponse
		require.NoError(t, json.Unmarshal(resp, &key))
		assert.NotEmpty(t, key.Results)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		deleteAtlasAccessListEntry(t, cliPath, entry, apiKeyID)
	})

	g.Run("Create Current IP", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			apiKeyAccessListEntity,
			"create",
			"--apiKey",
			apiKeyID,
			"--currentIp",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var key atlasv2.PaginatedApiUserAccessListResponse
		require.NoError(t, json.Unmarshal(resp, &key))
		assert.NotEmpty(t, key.Results)
		entry = *key.GetResults()[0].IpAddress
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		deleteAtlasAccessListEntry(t, cliPath, entry, apiKeyID)
	})
}

func deleteAtlasAccessListEntry(t *testing.T, cliPath, entry, apiKeyID string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		orgEntity,
		apiKeysEntity,
		apiKeyAccessListEntity,
		"rm",
		entry,
		"--apiKey",
		apiKeyID,
		"--force")
	cmd.Env = os.Environ()
	resp, err := RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
	expected := fmt.Sprintf("Access list entry '%s' deleted\n", entry)
	assert.Equal(t, expected, string(resp))
}
