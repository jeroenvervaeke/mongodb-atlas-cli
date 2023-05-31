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

type DeleteLDAPConfigurationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *DeleteLDAPConfigurationOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteLDAPConfigurationOpts) Run(ctx context.Context) error {
	params := &admin.DeleteLDAPConfigurationApiParams{
		GroupId: opts.groupId,
	}
	_, err := opts.client.LDAPConfigurationApi.DeleteLDAPConfigurationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(nil)
}

func DeleteLDAPConfigurationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeleteLDAPConfigurationOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Remove the Current LDAP User to DN Mapping",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")

	_ = cmd.MarkFlagRequired("groupId")

	return cmd
}
type GetLDAPConfigurationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *GetLDAPConfigurationOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetLDAPConfigurationOpts) Run(ctx context.Context) error {
	params := &admin.GetLDAPConfigurationApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.LDAPConfigurationApi.GetLDAPConfigurationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetLDAPConfigurationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetLDAPConfigurationOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return the Current LDAP or X.509 Configuration",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")

	_ = cmd.MarkFlagRequired("groupId")

	return cmd
}
type GetLDAPConfigurationStatusOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	requestId string
}

func (opts *GetLDAPConfigurationStatusOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetLDAPConfigurationStatusOpts) Run(ctx context.Context) error {
	params := &admin.GetLDAPConfigurationStatusApiParams{
		GroupId: opts.groupId,
		RequestId: opts.requestId,
	}
	resp, _, err := opts.client.LDAPConfigurationApi.GetLDAPConfigurationStatusWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetLDAPConfigurationStatusBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetLDAPConfigurationStatusOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return the Status of One Verify LDAP Configuration Request",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.requestId, "requestId", , "Unique string that identifies the request to verify an &lt;abbr title&#x3D;\&quot;Lightweight Directory Access Protocol\&quot;&gt;LDAP&lt;/abbr&gt; configuration.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("requestId")

	return cmd
}
type SaveLDAPConfigurationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *SaveLDAPConfigurationOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *SaveLDAPConfigurationOpts) Run(ctx context.Context) error {
	params := &admin.SaveLDAPConfigurationApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.LDAPConfigurationApi.SaveLDAPConfigurationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func SaveLDAPConfigurationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := SaveLDAPConfigurationOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Edit the LDAP or X.509 Configuration",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")

	_ = cmd.MarkFlagRequired("groupId")

	return cmd
}
type VerifyLDAPConfigurationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *VerifyLDAPConfigurationOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *VerifyLDAPConfigurationOpts) Run(ctx context.Context) error {
	params := &admin.VerifyLDAPConfigurationApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.LDAPConfigurationApi.VerifyLDAPConfigurationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func VerifyLDAPConfigurationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := VerifyLDAPConfigurationOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Verify the LDAP Configuration in One Project",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"), // how to tell?
		Args:    require.NoArgs,
		Annotations: map[string]string{
			"output":      template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				//opts.ValidateProjectID,
				opts.initClient(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", , "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")

	_ = cmd.MarkFlagRequired("groupId")

	return cmd
}

func LDAPConfigurationApiBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "<<use>>",
		Short:   "Returns, edits, verifies, and removes LDAP configurations. An LDAP configuration defines settings for MongoDB Cloud to connect to your LDAP server over TLS for user authentication and authorization. Your LDAP server must be visible to the internet or connected to your MongoDB Cloud cluster with VPC Peering. Also, your LDAP server must use TLS. You must have the MongoDB Cloud admin user privilege to use these endpoints. Also, to configure user authentication and authorization with LDAPS, your cluster must run MongoDB 3.6 or higher. Groups for which you have configured LDAPS can&#39;t create a cluster using a version of MongoDB 3.6 or lower.",
	}
	cmd.AddCommand(
		DeleteLDAPConfigurationBuilder(),
		GetLDAPConfigurationBuilder(),
		GetLDAPConfigurationStatusBuilder(),
		SaveLDAPConfigurationBuilder(),
		VerifyLDAPConfigurationBuilder(),
	)
	return cmd
}
