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
	"go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func TestAtlasOrgs(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)

	var orgID string
	// This test must run first to grab the ID of the org to later describe
	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			orgEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err2 := RunAndGetStdOut(cmd)
		require.NoError(t, err2, string(resp))
		var orgs admin.PaginatedOrganization
		err = json.Unmarshal(resp, &orgs)
		require.NoError(t, err, string(resp))
		assert.NotEmpty(t, orgs.GetResults())
		orgID = orgs.GetResults()[0].GetId()
	})
	require.NotEmpty(t, orgID)

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			orgEntity,
			"describe",
			orgID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err2 := RunAndGetStdOut(cmd)
		require.NoError(t, err2, string(resp))
	})

	var userID string
	g.Run("List Org Users", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			orgEntity,
			usersEntity,
			"ls",
			"--orgId",
			orgID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err2 := RunAndGetStdOut(cmd)
		require.NoError(t, err2, string(resp))
		var users admin.PaginatedOrgUser
		require.NoError(t, json.Unmarshal(resp, &users), string(resp))
		assert.NotEmpty(t, users.GetResults())
		userID = users.GetResults()[0].GetId()
	})
	require.NotEmpty(t, userID)

	n := g.memoryRand("rand", 255)
	orgName := fmt.Sprintf("e2e-org-%v", n)
	var (
		publicAPIKey  string
		privateAPIKey string
	)
	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		t.Skip("Skipping create org e2e test, exceeded max number of linked orgs. Will reenable post cleanup")
		cmd := exec.Command(cliPath,
			orgEntity,
			"create",
			orgName,
			"--ownerId",
			userID,
			"--apiKeyRole",
			"ORG_OWNER",
			"--apiKeyDescription",
			"test",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var org admin.CreateOrganizationResponse
		require.NoError(t, json.Unmarshal(resp, &org), string(resp))
		orgID = org.Organization.GetId()
		publicAPIKey = org.ApiKey.GetPublicKey()
		privateAPIKey = org.ApiKey.GetPrivateKey()

		require.NotEmpty(t, publicAPIKey)
		require.NotEmpty(t, privateAPIKey)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		t.Skip("Skipping delete org e2e test, exceeded max number of linked orgs. Will re-enable post cleanup")
		if os.Getenv("MONGODB_ATLAS_SERVICE") == cloudgov {
			t.Skip("not available for gov")
		}
		t.Setenv("MONGODB_ATLAS_PUBLIC_API_KEY", publicAPIKey)
		t.Setenv("MONGODB_ATLAS_PRIVATE_API_KEY", privateAPIKey)
		cmd := exec.Command(cliPath,
			orgEntity,
			"delete",
			orgID,
			"--force",
		)
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})
}
