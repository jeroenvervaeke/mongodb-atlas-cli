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

type AddAllTeamsToProjectOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *AddAllTeamsToProjectOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *AddAllTeamsToProjectOpts) Run(ctx context.Context) error {
	params := &admin.AddAllTeamsToProjectApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.TeamsApi.AddAllTeamsToProjectWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func AddAllTeamsToProjectBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := AddAllTeamsToProjectOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Add One or More Teams to One Project",
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
type AddTeamUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	teamId string
}

func (opts *AddTeamUserOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *AddTeamUserOpts) Run(ctx context.Context) error {
	params := &admin.AddTeamUserApiParams{
		OrgId: opts.orgId,
		TeamId: opts.teamId,
	}
	resp, _, err := opts.client.TeamsApi.AddTeamUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func AddTeamUserBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := AddTeamUserOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Assign MongoDB Cloud Users from One Organization to One Team",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")
	cmd.Flags().StringVar(&opts.teamId, "teamId", , "Unique 24-hexadecimal character string that identifies the team to which you want to add MongoDB Cloud users.")

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("teamId")

	return cmd
}
type CreateTeamOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
}

func (opts *CreateTeamOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateTeamOpts) Run(ctx context.Context) error {
	params := &admin.CreateTeamApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.TeamsApi.CreateTeamWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreateTeamBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreateTeamOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Create One Team in One Organization",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")

	_ = cmd.MarkFlagRequired("orgId")

	return cmd
}
type DeleteTeamOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	teamId string
}

func (opts *DeleteTeamOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteTeamOpts) Run(ctx context.Context) error {
	params := &admin.DeleteTeamApiParams{
		OrgId: opts.orgId,
		TeamId: opts.teamId,
	}
	resp, _, err := opts.client.TeamsApi.DeleteTeamWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DeleteTeamBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeleteTeamOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Remove One Team from One Organization",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")
	cmd.Flags().StringVar(&opts.teamId, "teamId", , "Unique 24-hexadecimal digit string that identifies the team that you want to delete.")

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("teamId")

	return cmd
}
type GetTeamByIdOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	teamId string
}

func (opts *GetTeamByIdOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetTeamByIdOpts) Run(ctx context.Context) error {
	params := &admin.GetTeamByIdApiParams{
		OrgId: opts.orgId,
		TeamId: opts.teamId,
	}
	resp, _, err := opts.client.TeamsApi.GetTeamByIdWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetTeamByIdBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetTeamByIdOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return One Team using its ID",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")
	cmd.Flags().StringVar(&opts.teamId, "teamId", , "Unique 24-hexadecimal digit string that identifies the team whose information you want to return.")

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("teamId")

	return cmd
}
type GetTeamByNameOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	teamName string
}

func (opts *GetTeamByNameOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetTeamByNameOpts) Run(ctx context.Context) error {
	params := &admin.GetTeamByNameApiParams{
		OrgId: opts.orgId,
		TeamName: opts.teamName,
	}
	resp, _, err := opts.client.TeamsApi.GetTeamByNameWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetTeamByNameBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetTeamByNameOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return One Team using its Name",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")
	cmd.Flags().StringVar(&opts.teamName, "teamName", , "Name of the team whose information you want to return.")

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("teamName")

	return cmd
}
type ListOrganizationTeamsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	itemsPerPage int32
	includeCount bool
	pageNum int32
}

func (opts *ListOrganizationTeamsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListOrganizationTeamsOpts) Run(ctx context.Context) error {
	params := &admin.ListOrganizationTeamsApiParams{
		OrgId: opts.orgId,
		ItemsPerPage: opts.itemsPerPage,
		IncludeCount: opts.includeCount,
		PageNum: opts.pageNum,
	}
	resp, _, err := opts.client.TeamsApi.ListOrganizationTeamsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListOrganizationTeamsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListOrganizationTeamsOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return All Teams in One Organization",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")
	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")
	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")

	_ = cmd.MarkFlagRequired("orgId")

	return cmd
}
type ListProjectTeamsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	includeCount bool
	itemsPerPage int32
	pageNum int32
}

func (opts *ListProjectTeamsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListProjectTeamsOpts) Run(ctx context.Context) error {
	params := &admin.ListProjectTeamsApiParams{
		GroupId: opts.groupId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
	}
	resp, _, err := opts.client.TeamsApi.ListProjectTeamsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListProjectTeamsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListProjectTeamsOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return All Teams in One Project",
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
	cmd.Flags().StringVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")
	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")
	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")

	_ = cmd.MarkFlagRequired("groupId")

	return cmd
}
type ListTeamUsersOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	teamId string
	itemsPerPage int32
	pageNum int32
}

func (opts *ListTeamUsersOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListTeamUsersOpts) Run(ctx context.Context) error {
	params := &admin.ListTeamUsersApiParams{
		OrgId: opts.orgId,
		TeamId: opts.teamId,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
	}
	resp, _, err := opts.client.TeamsApi.ListTeamUsersWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListTeamUsersBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListTeamUsersOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return All MongoDB Cloud Users Assigned to One Team",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")
	cmd.Flags().StringVar(&opts.teamId, "teamId", , "Unique 24-hexadecimal digit string that identifies the team whose application users you want to return.")
	cmd.Flags().StringVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")
	cmd.Flags().StringVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("teamId")

	return cmd
}
type RemoveProjectTeamOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	teamId string
}

func (opts *RemoveProjectTeamOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *RemoveProjectTeamOpts) Run(ctx context.Context) error {
	params := &admin.RemoveProjectTeamApiParams{
		GroupId: opts.groupId,
		TeamId: opts.teamId,
	}
	resp, _, err := opts.client.TeamsApi.RemoveProjectTeamWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func RemoveProjectTeamBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := RemoveProjectTeamOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Remove One Team from One Project",
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
	cmd.Flags().StringVar(&opts.teamId, "teamId", , "Unique 24-hexadecimal digit string that identifies the team that you want to remove from the specified project.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("teamId")

	return cmd
}
type RemoveTeamUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	teamId string
	userId string
}

func (opts *RemoveTeamUserOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *RemoveTeamUserOpts) Run(ctx context.Context) error {
	params := &admin.RemoveTeamUserApiParams{
		OrgId: opts.orgId,
		TeamId: opts.teamId,
		UserId: opts.userId,
	}
	_, err := opts.client.TeamsApi.RemoveTeamUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(nil)
}

func RemoveTeamUserBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := RemoveTeamUserOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Remove One MongoDB Cloud User from One Team",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")
	cmd.Flags().StringVar(&opts.teamId, "teamId", , "Unique 24-hexadecimal digit string that identifies the team from which you want to remove one database application user.")
	cmd.Flags().StringVar(&opts.userId, "userId", , "Unique 24-hexadecimal digit string that identifies MongoDB Cloud user that you want to remove from the specified team.")

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("teamId")
	_ = cmd.MarkFlagRequired("userId")

	return cmd
}
type RenameTeamOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId string
	teamId string
}

func (opts *RenameTeamOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *RenameTeamOpts) Run(ctx context.Context) error {
	params := &admin.RenameTeamApiParams{
		OrgId: opts.orgId,
		TeamId: opts.teamId,
	}
	resp, _, err := opts.client.TeamsApi.RenameTeamWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func RenameTeamBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := RenameTeamOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Rename One Team",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", , "Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.")
	cmd.Flags().StringVar(&opts.teamId, "teamId", , "Unique 24-hexadecimal digit string that identifies the team that you want to rename.")

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("teamId")

	return cmd
}
type UpdateTeamRolesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	teamId string
}

func (opts *UpdateTeamRolesOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateTeamRolesOpts) Run(ctx context.Context) error {
	params := &admin.UpdateTeamRolesApiParams{
		GroupId: opts.groupId,
		TeamId: opts.teamId,
	}
	resp, _, err := opts.client.TeamsApi.UpdateTeamRolesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdateTeamRolesBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdateTeamRolesOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Update Team Roles in One Project",
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
	cmd.Flags().StringVar(&opts.teamId, "teamId", , "Unique 24-hexadecimal digit string that identifies the team for which you want to update roles.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("teamId")

	return cmd
}

func TeamsApiBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "<<use>>",
		Short:   "Returns, adds, edits, or removes teams.",
	}
	cmd.AddCommand(
		AddAllTeamsToProjectBuilder(),
		AddTeamUserBuilder(),
		CreateTeamBuilder(),
		DeleteTeamBuilder(),
		GetTeamByIdBuilder(),
		GetTeamByNameBuilder(),
		ListOrganizationTeamsBuilder(),
		ListProjectTeamsBuilder(),
		ListTeamUsersBuilder(),
		RemoveProjectTeamBuilder(),
		RemoveTeamUserBuilder(),
		RenameTeamBuilder(),
		UpdateTeamRolesBuilder(),
	)
	return cmd
}
