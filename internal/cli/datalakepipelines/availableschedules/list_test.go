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

// This code was autogenerated at 2023-04-27T17:56:12+01:00. Note: Manual updates are allowed, but may be overwritten.

package availableschedules

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113005/admin"
)

func TestListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockPipelineAvailableSchedulesLister(ctrl)

	expected := []atlasv2.DiskBackupApiPolicyItem{
		{
			Id:                pointer.Get("5e4e593f70dfbf1010295836"),
			FrequencyInterval: 1,
			FrequencyType:     "daily",
			RetentionUnit:     "months",
			RetentionValue:    1,
		},
	}

	buf := new(bytes.Buffer)
	listOpts := &ListOpts{
		store:        mockStore,
		pipelineName: "Pipeline1",
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		PipelineAvailableSchedules(listOpts.ProjectID, listOpts.pipelineName).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `ID                         FREQUENCY INTERVAL   FREQUENCY TYPE   RETENTION UNIT   RETENTION VALUE
5e4e593f70dfbf1010295836   1                    daily            months           1`, buf.String())
	t.Log(buf.String())
}
