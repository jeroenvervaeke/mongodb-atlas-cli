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

//go:build unit

package teams

import (
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115001/admin"
)

func TestUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockTeamRolesUpdater(ctrl)

	expected := atlasv2.PaginatedTeamRole{}

	updateOpts := &UpdateOpts{
		roles: []string{"test"},
		store: mockStore,
	}

	mockStore.
		EXPECT().
		UpdateProjectTeamRoles(updateOpts.ProjectID, updateOpts.teamID, updateOpts.newTeamUpdateRoles()).
		Return(&expected, nil).
		Times(1)

	if err := updateOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
