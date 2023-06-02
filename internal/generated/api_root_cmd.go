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

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package generated

import (
	"context"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/admin"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
)

type GetSystemStatusOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
}

func (opts *GetSystemStatusOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetSystemStatusOpts) Run(ctx context.Context) error {
	params := &admin.GetSystemStatusApiParams{
	}
	resp, _, err := opts.client.RootApi.GetSystemStatusWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetSystemStatusBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetSystemStatusOpts{}
	cmd := &cobra.Command{
		Use:     "getSystemStatus",
		// Aliases: []string{"?"},
		Short:   "Return the status of this MongoDB application",
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	return cmd
}

func RootBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "root",
		Short:   "Returns details that describe the MongoDB Cloud build and the access token that requests this resource. This starts the MongoDB Cloud API.",
	}
	cmd.AddCommand(
		GetSystemStatusBuilder(),
	)
	return cmd
}

