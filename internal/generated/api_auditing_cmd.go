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

type GetAuditingConfigurationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *GetAuditingConfigurationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetAuditingConfigurationOpts) Run(ctx context.Context) error {
	params := &admin.GetAuditingConfigurationApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.AuditingApi.GetAuditingConfigurationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetAuditingConfigurationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetAuditingConfigurationOpts{}
	cmd := &cobra.Command{
		Use:     "getAuditingConfiguration",
		// Aliases: []string{"?"},
		Short:   "Return the Auditing Configuration for One Project",
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}
type UpdateAuditingConfigurationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	
}

func (opts *UpdateAuditingConfigurationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateAuditingConfigurationOpts) Run(ctx context.Context) error {
	params := &admin.UpdateAuditingConfigurationApiParams{
		GroupId: opts.groupId,
		
	}
	resp, _, err := opts.client.AuditingApi.UpdateAuditingConfigurationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdateAuditingConfigurationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdateAuditingConfigurationOpts{}
	cmd := &cobra.Command{
		Use:     "updateAuditingConfiguration",
		// Aliases: []string{"?"},
		Short:   "Update Auditing Configuration for One Project",
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

func AuditingBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "auditing",
		Short:   "Returns and edits database auditing settings for MongoDB Cloud projects.",
	}
	cmd.AddCommand(
		GetAuditingConfigurationBuilder(),
		UpdateAuditingConfigurationBuilder(),
	)
	return cmd
}

