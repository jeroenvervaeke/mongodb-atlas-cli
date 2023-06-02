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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/admin"
)

type createOrganizationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
}

func (opts *createOrganizationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createOrganizationOpts) Run(ctx context.Context) error {
	params := &admin.CreateOrganizationApiParams{}
	resp, _, err := opts.client.OrganizationsApi.CreateOrganizationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func createOrganizationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := createOrganizationOpts{}
	cmd := &cobra.Command{
		Use:   "createOrganization",
		Short: "Create One Organization",
		Annotations: map[string]string{
			"output": template,
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

type createOrganizationInvitationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId  string
}

func (opts *createOrganizationInvitationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createOrganizationInvitationOpts) Run(ctx context.Context) error {
	params := &admin.CreateOrganizationInvitationApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.CreateOrganizationInvitationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func createOrganizationInvitationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := createOrganizationInvitationOpts{}
	cmd := &cobra.Command{
		Use:   "createOrganizationInvitation",
		Short: "Invite One MongoDB Cloud User to Join One Atlas Organization",
		Annotations: map[string]string{
			"output": template,
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type deleteOrganizationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId  string
}

func (opts *deleteOrganizationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deleteOrganizationOpts) Run(ctx context.Context) error {
	params := &admin.DeleteOrganizationApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.DeleteOrganizationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func deleteOrganizationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := deleteOrganizationOpts{}
	cmd := &cobra.Command{
		Use:   "deleteOrganization",
		Short: "Remove One Organization",
		Annotations: map[string]string{
			"output": template,
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type deleteOrganizationInvitationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	orgId        string
	invitationId string
}

func (opts *deleteOrganizationInvitationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deleteOrganizationInvitationOpts) Run(ctx context.Context) error {
	params := &admin.DeleteOrganizationInvitationApiParams{
		OrgId:        opts.orgId,
		InvitationId: opts.invitationId,
	}
	resp, _, err := opts.client.OrganizationsApi.DeleteOrganizationInvitationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func deleteOrganizationInvitationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := deleteOrganizationInvitationOpts{}
	cmd := &cobra.Command{
		Use:   "deleteOrganizationInvitation",
		Short: "Cancel One Organization Invitation",
		Annotations: map[string]string{
			"output": template,
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.invitationId, "invitationId", "", `Unique 24-hexadecimal digit string that identifies the invitation.`)

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("invitationId")
	return cmd
}

type getOrganizationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId  string
}

func (opts *getOrganizationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getOrganizationOpts) Run(ctx context.Context) error {
	params := &admin.GetOrganizationApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.GetOrganizationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func getOrganizationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := getOrganizationOpts{}
	cmd := &cobra.Command{
		Use:   "getOrganization",
		Short: "Return One Organization",
		Annotations: map[string]string{
			"output": template,
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type getOrganizationInvitationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	orgId        string
	invitationId string
}

func (opts *getOrganizationInvitationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getOrganizationInvitationOpts) Run(ctx context.Context) error {
	params := &admin.GetOrganizationInvitationApiParams{
		OrgId:        opts.orgId,
		InvitationId: opts.invitationId,
	}
	resp, _, err := opts.client.OrganizationsApi.GetOrganizationInvitationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func getOrganizationInvitationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := getOrganizationInvitationOpts{}
	cmd := &cobra.Command{
		Use:   "getOrganizationInvitation",
		Short: "Return One Organization Invitation",
		Annotations: map[string]string{
			"output": template,
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.invitationId, "invitationId", "", `Unique 24-hexadecimal digit string that identifies the invitation.`)

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("invitationId")
	return cmd
}

type getOrganizationSettingsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId  string
}

func (opts *getOrganizationSettingsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getOrganizationSettingsOpts) Run(ctx context.Context) error {
	params := &admin.GetOrganizationSettingsApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.GetOrganizationSettingsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func getOrganizationSettingsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := getOrganizationSettingsOpts{}
	cmd := &cobra.Command{
		Use:   "getOrganizationSettings",
		Short: "Return Settings for One Organization",
		Annotations: map[string]string{
			"output": template,
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type listOrganizationInvitationsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client   *admin.APIClient
	orgId    string
	username string
}

func (opts *listOrganizationInvitationsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listOrganizationInvitationsOpts) Run(ctx context.Context) error {
	params := &admin.ListOrganizationInvitationsApiParams{
		OrgId:    opts.orgId,
		Username: &opts.username,
	}
	resp, _, err := opts.client.OrganizationsApi.ListOrganizationInvitationsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func listOrganizationInvitationsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := listOrganizationInvitationsOpts{}
	cmd := &cobra.Command{
		Use:   "listOrganizationInvitations",
		Short: "Return All Organization Invitations",
		Annotations: map[string]string{
			"output": template,
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.username, "username", "", `Email address of the user account invited to this organization. If you exclude this parameter, this resource returns all pending invitations.`)

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type listOrganizationProjectsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	orgId        string
	includeCount bool
	itemsPerPage int
	pageNum      int
	name         string
}

func (opts *listOrganizationProjectsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listOrganizationProjectsOpts) Run(ctx context.Context) error {
	params := &admin.ListOrganizationProjectsApiParams{
		OrgId:        opts.orgId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
		Name:         &opts.name,
	}
	resp, _, err := opts.client.OrganizationsApi.ListOrganizationProjectsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func listOrganizationProjectsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := listOrganizationProjectsOpts{}
	cmd := &cobra.Command{
		Use:   "listOrganizationProjects",
		Short: "Return One or More Projects in One Organization",
		Annotations: map[string]string{
			"output": template,
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)
	cmd.Flags().StringVar(&opts.name, "name", "", `Human-readable label of the project to use to filter the returned list. Performs a case-insensitive search for a project within the organization which is prefixed by the specified name.`)

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type listOrganizationUsersOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	orgId        string
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listOrganizationUsersOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listOrganizationUsersOpts) Run(ctx context.Context) error {
	params := &admin.ListOrganizationUsersApiParams{
		OrgId:        opts.orgId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}
	resp, _, err := opts.client.OrganizationsApi.ListOrganizationUsersWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func listOrganizationUsersBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := listOrganizationUsersOpts{}
	cmd := &cobra.Command{
		Use:   "listOrganizationUsers",
		Short: "Return All MongoDB Cloud Users in One Organization",
		Annotations: map[string]string{
			"output": template,
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type listOrganizationsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	includeCount bool
	itemsPerPage int
	pageNum      int
	name         string
}

func (opts *listOrganizationsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listOrganizationsOpts) Run(ctx context.Context) error {
	params := &admin.ListOrganizationsApiParams{
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
		Name:         &opts.name,
	}
	resp, _, err := opts.client.OrganizationsApi.ListOrganizationsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func listOrganizationsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := listOrganizationsOpts{}
	cmd := &cobra.Command{
		Use:   "listOrganizations",
		Short: "Return All Organizations",
		Annotations: map[string]string{
			"output": template,
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
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)
	cmd.Flags().StringVar(&opts.name, "name", "", `Human-readable label of the organization to use to filter the returned list. Performs a case-insensitive search for an organization that starts with the specified name.`)

	return cmd
}

type renameOrganizationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId  string
}

func (opts *renameOrganizationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *renameOrganizationOpts) Run(ctx context.Context) error {
	params := &admin.RenameOrganizationApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.RenameOrganizationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func renameOrganizationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := renameOrganizationOpts{}
	cmd := &cobra.Command{
		Use:   "renameOrganization",
		Short: "Rename One Organization",
		Annotations: map[string]string{
			"output": template,
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type updateOrganizationInvitationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId  string
}

func (opts *updateOrganizationInvitationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updateOrganizationInvitationOpts) Run(ctx context.Context) error {
	params := &admin.UpdateOrganizationInvitationApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.UpdateOrganizationInvitationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func updateOrganizationInvitationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := updateOrganizationInvitationOpts{}
	cmd := &cobra.Command{
		Use:   "updateOrganizationInvitation",
		Short: "Update One Organization Invitation",
		Annotations: map[string]string{
			"output": template,
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type updateOrganizationInvitationByIdOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	orgId        string
	invitationId string
}

func (opts *updateOrganizationInvitationByIdOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updateOrganizationInvitationByIdOpts) Run(ctx context.Context) error {
	params := &admin.UpdateOrganizationInvitationByIdApiParams{
		OrgId:        opts.orgId,
		InvitationId: opts.invitationId,
	}
	resp, _, err := opts.client.OrganizationsApi.UpdateOrganizationInvitationByIdWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func updateOrganizationInvitationByIdBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := updateOrganizationInvitationByIdOpts{}
	cmd := &cobra.Command{
		Use:   "updateOrganizationInvitationById",
		Short: "Update One Organization Invitation by Invitation ID",
		Annotations: map[string]string{
			"output": template,
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)
	cmd.Flags().StringVar(&opts.invitationId, "invitationId", "", `Unique 24-hexadecimal digit string that identifies the invitation.`)

	_ = cmd.MarkFlagRequired("orgId")
	_ = cmd.MarkFlagRequired("invitationId")
	return cmd
}

type updateOrganizationSettingsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	orgId  string
}

func (opts *updateOrganizationSettingsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updateOrganizationSettingsOpts) Run(ctx context.Context) error {
	params := &admin.UpdateOrganizationSettingsApiParams{
		OrgId: opts.orgId,
	}
	resp, _, err := opts.client.OrganizationsApi.UpdateOrganizationSettingsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func updateOrganizationSettingsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := updateOrganizationSettingsOpts{}
	cmd := &cobra.Command{
		Use:   "updateOrganizationSettings",
		Short: "Update Settings for One Organization",
		Annotations: map[string]string{
			"output": template,
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.`)

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

func organizationsBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "organizations",
		Short: `Returns, adds, and edits organizational units in MongoDB Cloud.`,
	}
	cmd.AddCommand(
		createOrganizationBuilder(),
		createOrganizationInvitationBuilder(),
		deleteOrganizationBuilder(),
		deleteOrganizationInvitationBuilder(),
		getOrganizationBuilder(),
		getOrganizationInvitationBuilder(),
		getOrganizationSettingsBuilder(),
		listOrganizationInvitationsBuilder(),
		listOrganizationProjectsBuilder(),
		listOrganizationUsersBuilder(),
		listOrganizationsBuilder(),
		renameOrganizationBuilder(),
		updateOrganizationInvitationBuilder(),
		updateOrganizationInvitationByIdBuilder(),
		updateOrganizationSettingsBuilder(),
	)
	return cmd
}
