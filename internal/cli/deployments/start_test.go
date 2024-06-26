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

package deployments

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/test/fixture"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20231115014/admin"
)

func TestStart_RunLocal_PausedContainers(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	deploymentsTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)
	mockPodman := deploymentsTest.MockPodman

	buf := new(bytes.Buffer)
	startOpts := &StartOpts{
		DeploymentOpts: *deploymentsTest.Opts,
		GlobalOpts: cli.GlobalOpts{
			ProjectID: projectID,
		},
		OutputOpts: cli.OutputOpts{
			Template:  startTemplate,
			OutWriter: buf,
		},
	}

	expected := deploymentsTest.MockContainerWithState("paused")
	deploymentsTest.LocalMockFlowWithMockContainer(ctx, expected)

	mockPodman.
		EXPECT().
		ContainerInspect(ctx, startOpts.LocalMongotHostname()).
		Return([]*podman.InspectContainerData{
			{
				NetworkSettings: &podman.InspectNetworkSettings{
					Networks: map[string]*podman.InspectAdditionalNetwork{
						startOpts.LocalNetworkName(): {
							InspectBasicNetworkConfig: podman.InspectBasicNetworkConfig{
								IPAddress: "1.2.3.4",
							},
						},
					},
				},
			},
		}, nil).
		Times(1)

	mockPodman.
		EXPECT().
		Exec(gomock.Any(), startOpts.LocalMongodHostname(), "/bin/sh", "-c", gomock.Any()).
		Return(nil).
		Times(1)

	mockPodman.
		EXPECT().
		UnpauseContainers(ctx, startOpts.LocalMongodHostname(), startOpts.LocalMongotHostname()).
		Return(nil, nil).
		Times(1)

	require.NoError(t, startOpts.Run(ctx))
	assert.Equal(t, fmt.Sprintf("\nStarting deployment '%s'.\n", deploymentName), buf.String())
	t.Log(buf.String())
}

func TestStart_RunLocal_StoppedContainers(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	deploymentsTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)
	mockPodman := deploymentsTest.MockPodman

	buf := new(bytes.Buffer)
	startOpts := &StartOpts{
		DeploymentOpts: *deploymentsTest.Opts,
		GlobalOpts: cli.GlobalOpts{
			ProjectID: projectID,
		},
		OutputOpts: cli.OutputOpts{
			Template:  startTemplate,
			OutWriter: buf,
		},
	}

	expected := deploymentsTest.MockContainerWithState("exited")
	deploymentsTest.LocalMockFlowWithMockContainer(ctx, expected)

	mockPodman.
		EXPECT().
		ContainerInspect(ctx, startOpts.LocalMongotHostname()).
		Return([]*podman.InspectContainerData{
			{
				NetworkSettings: &podman.InspectNetworkSettings{
					Networks: map[string]*podman.InspectAdditionalNetwork{
						startOpts.LocalNetworkName(): {
							InspectBasicNetworkConfig: podman.InspectBasicNetworkConfig{
								IPAddress: "1.2.3.4",
							},
						},
					},
				},
			},
		}, nil).
		Times(1)

	mockPodman.
		EXPECT().
		Exec(gomock.Any(), startOpts.LocalMongodHostname(), "/bin/sh", "-c", gomock.Any()).
		Return(nil).
		Times(1)

	mockPodman.
		EXPECT().
		StartContainers(ctx, startOpts.LocalMongodHostname(), startOpts.LocalMongotHostname()).
		Return(nil, nil).
		Times(1)

	require.NoError(t, startOpts.Run(ctx))
	assert.Equal(t, fmt.Sprintf("\nStarting deployment '%s'.\n", deploymentName), buf.String())
	t.Log(buf.String())
}

func TestStart_RunAtlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterStarter(ctrl)
	ctx := context.Background()
	const deploymentName = "atlasCluster1"

	deploymentsTest := fixture.NewMockAtlasDeploymentOpts(ctrl, deploymentName)

	buf := new(bytes.Buffer)
	opts := &StartOpts{
		store:          mockStore,
		DeploymentOpts: *deploymentsTest.Opts,
		GlobalOpts: cli.GlobalOpts{
			ProjectID: projectID,
		},
		OutputOpts: cli.OutputOpts{
			Template:  startTemplate,
			OutWriter: buf,
		},
	}

	deploymentsTest.CommonAtlasMocks(projectID)

	mockStore.
		EXPECT().
		StartCluster(projectID, deploymentName).
		Return(
			&admin.AdvancedClusterDescription{
				Name: pointer.Get(deploymentName),
			}, nil).
		Times(1)

	require.NoError(t, opts.Run(ctx))
	assert.Equal(t, fmt.Sprintf("\nStarting deployment '%s'.\n", deploymentName), buf.String())
	t.Log(buf.String())
}

func TestStartOpts_PostRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	deploymentsTest := fixture.NewMockLocalDeploymentOpts(ctrl, "localDeployment")
	buf := new(bytes.Buffer)

	opts := &StartOpts{
		DeploymentOpts: *deploymentsTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
	}

	deploymentsTest.
		MockDeploymentTelemetry.
		EXPECT().
		AppendDeploymentType().
		Times(1)

	require.NoError(t, opts.PostRun())
}

func TestStartBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		StartBuilder(),
		0,
		[]string{flag.ProjectID, flag.TypeFlag},
	)
}
