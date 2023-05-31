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

type CreatePeeringConnectionOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *CreatePeeringConnectionOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreatePeeringConnectionOpts) Run(ctx context.Context) error {
	params := &admin.CreatePeeringConnectionApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.CreatePeeringConnectionWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreatePeeringConnectionBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreatePeeringConnectionOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Create One New Network Peering Connection",
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
type CreatePeeringContainerOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *CreatePeeringContainerOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreatePeeringContainerOpts) Run(ctx context.Context) error {
	params := &admin.CreatePeeringContainerApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.CreatePeeringContainerWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreatePeeringContainerBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreatePeeringContainerOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Create One New Network Peering Container",
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
type DeletePeeringConnectionOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	peerId string
}

func (opts *DeletePeeringConnectionOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeletePeeringConnectionOpts) Run(ctx context.Context) error {
	params := &admin.DeletePeeringConnectionApiParams{
		GroupId: opts.groupId,
		PeerId: opts.peerId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.DeletePeeringConnectionWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DeletePeeringConnectionBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeletePeeringConnectionOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Remove One Existing Network Peering Connection",
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
	cmd.Flags().StringVar(&opts.peerId, "peerId", , "Unique 24-hexadecimal digit string that identifies the network peering connection that you want to delete.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("peerId")

	return cmd
}
type DeletePeeringContainerOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	containerId string
}

func (opts *DeletePeeringContainerOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeletePeeringContainerOpts) Run(ctx context.Context) error {
	params := &admin.DeletePeeringContainerApiParams{
		GroupId: opts.groupId,
		ContainerId: opts.containerId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.DeletePeeringContainerWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DeletePeeringContainerBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeletePeeringContainerOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Remove One Network Peering Container",
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
	cmd.Flags().StringVar(&opts.containerId, "containerId", , "Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("containerId")

	return cmd
}
type DisablePeeringOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *DisablePeeringOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DisablePeeringOpts) Run(ctx context.Context) error {
	params := &admin.DisablePeeringApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.DisablePeeringWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DisablePeeringBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DisablePeeringOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Disable Connect via Peering Only Mode for One Project",
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
type GetPeeringConnectionOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	peerId string
}

func (opts *GetPeeringConnectionOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetPeeringConnectionOpts) Run(ctx context.Context) error {
	params := &admin.GetPeeringConnectionApiParams{
		GroupId: opts.groupId,
		PeerId: opts.peerId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.GetPeeringConnectionWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetPeeringConnectionBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetPeeringConnectionOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return One Network Peering Connection in One Project",
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
	cmd.Flags().StringVar(&opts.peerId, "peerId", , "Unique 24-hexadecimal digit string that identifies the network peering connection that you want to retrieve.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("peerId")

	return cmd
}
type GetPeeringContainerOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	containerId string
}

func (opts *GetPeeringContainerOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetPeeringContainerOpts) Run(ctx context.Context) error {
	params := &admin.GetPeeringContainerApiParams{
		GroupId: opts.groupId,
		ContainerId: opts.containerId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.GetPeeringContainerWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetPeeringContainerBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetPeeringContainerOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return One Network Peering Container",
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
	cmd.Flags().StringVar(&opts.containerId, "containerId", , "Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("containerId")

	return cmd
}
type ListPeeringConnectionsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	includeCount bool
	itemsPerPage int32
	pageNum int32
	providerName string
}

func (opts *ListPeeringConnectionsOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListPeeringConnectionsOpts) Run(ctx context.Context) error {
	params := &admin.ListPeeringConnectionsApiParams{
		GroupId: opts.groupId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
		ProviderName: opts.providerName,
	}
	resp, _, err := opts.client.NetworkPeeringApi.ListPeeringConnectionsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListPeeringConnectionsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListPeeringConnectionsOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return All Network Peering Connections in One Project",
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
	cmd.Flags().StringVar(&opts.providerName, "providerName", &quot;AWS&quot;, "Cloud service provider to use for this VPC peering connection.")

	_ = cmd.MarkFlagRequired("groupId")

	return cmd
}
type ListPeeringContainerByCloudProviderOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	includeCount bool
	itemsPerPage int32
	pageNum int32
	providerName string
}

func (opts *ListPeeringContainerByCloudProviderOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListPeeringContainerByCloudProviderOpts) Run(ctx context.Context) error {
	params := &admin.ListPeeringContainerByCloudProviderApiParams{
		GroupId: opts.groupId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
		ProviderName: opts.providerName,
	}
	resp, _, err := opts.client.NetworkPeeringApi.ListPeeringContainerByCloudProviderWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListPeeringContainerByCloudProviderBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListPeeringContainerByCloudProviderOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return All Network Peering Containers in One Project for One Cloud Provider",
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
	cmd.Flags().StringVar(&opts.providerName, "providerName", &quot;AWS&quot;, "Cloud service provider that serves the desired network peering containers.")

	_ = cmd.MarkFlagRequired("groupId")

	_ = cmd.MarkFlagRequired("providerName")
	return cmd
}
type ListPeeringContainersOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	includeCount bool
	itemsPerPage int32
	pageNum int32
}

func (opts *ListPeeringContainersOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListPeeringContainersOpts) Run(ctx context.Context) error {
	params := &admin.ListPeeringContainersApiParams{
		GroupId: opts.groupId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
	}
	resp, _, err := opts.client.NetworkPeeringApi.ListPeeringContainersWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListPeeringContainersBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListPeeringContainersOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Return All Network Peering Containers in One Project",
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
type UpdatePeeringConnectionOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	peerId string
}

func (opts *UpdatePeeringConnectionOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdatePeeringConnectionOpts) Run(ctx context.Context) error {
	params := &admin.UpdatePeeringConnectionApiParams{
		GroupId: opts.groupId,
		PeerId: opts.peerId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.UpdatePeeringConnectionWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdatePeeringConnectionBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdatePeeringConnectionOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Update One New Network Peering Connection",
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
	cmd.Flags().StringVar(&opts.peerId, "peerId", , "Unique 24-hexadecimal digit string that identifies the network peering connection that you want to update.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("peerId")

	return cmd
}
type UpdatePeeringContainerOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	containerId string
}

func (opts *UpdatePeeringContainerOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdatePeeringContainerOpts) Run(ctx context.Context) error {
	params := &admin.UpdatePeeringContainerApiParams{
		GroupId: opts.groupId,
		ContainerId: opts.containerId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.UpdatePeeringContainerWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdatePeeringContainerBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdatePeeringContainerOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Update One Network Peering Container",
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
	cmd.Flags().StringVar(&opts.containerId, "containerId", , "Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("containerId")

	return cmd
}
type VerifyConnectViaPeeringOnlyModeForOneProjectOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *VerifyConnectViaPeeringOnlyModeForOneProjectOpts) initClient(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *VerifyConnectViaPeeringOnlyModeForOneProjectOpts) Run(ctx context.Context) error {
	params := &admin.VerifyConnectViaPeeringOnlyModeForOneProjectApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.NetworkPeeringApi.VerifyConnectViaPeeringOnlyModeForOneProjectWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func VerifyConnectViaPeeringOnlyModeForOneProjectBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := VerifyConnectViaPeeringOnlyModeForOneProjectOpts{}
	cmd := &cobra.Command{
		Use:     "<<use>>",
		// Aliases: []string{"?"},
		Short:   "Verify Connect via Peering Only Mode for One Project",
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
