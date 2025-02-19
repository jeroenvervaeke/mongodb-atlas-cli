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

package events

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"go.mongodb.org/atlas-sdk/v20241113005/admin"
)

func Test_orgListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOrganizationEventLister(ctrl)

	expected := &admin.OrgPaginatedEvent{}
	listOpts := &orgListOpts{
		store: mockStore,
	}
	listOpts.OrgID = "1"
	anyMock := gomock.Any()
	mockStore.
		EXPECT().OrganizationEvents(anyMock).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func Test_orgListOpts_Run_WithDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOrganizationEventLister(ctrl)

	expected := &admin.OrgPaginatedEvent{}
	listOpts := &orgListOpts{
		store: mockStore,
		EventListOpts: EventListOpts{
			MaxDate: "2024-03-18T15:00:03-0000",
			MinDate: "2024-03-18T14:40:03-0000",
		},
	}
	listOpts.OrgID = "1"
	anyMock := gomock.Any()
	mockStore.
		EXPECT().OrganizationEvents(anyMock).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func Test_orgListOpts_Run_WithInvalidDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOrganizationEventLister(ctrl)

	listOpts := &orgListOpts{
		store: mockStore,
		EventListOpts: EventListOpts{
			MaxDate: "2024-03-18T15:00:03+00:00Z",
			MinDate: "2024-03-18T13:00:03+00:00Z",
		},
	}
	listOpts.OrgID = "1"

	if err := listOpts.Run(); err == nil {
		t.Fatalf("Expected inavlid date error from Run() got none")
	}
}
