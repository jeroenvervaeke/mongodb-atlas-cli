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
	"os"
	"time"

	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/admin"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
)

type DeferMaintenanceWindowOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *DeferMaintenanceWindowOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeferMaintenanceWindowOpts) Run(ctx context.Context) error {
	params := &admin.DeferMaintenanceWindowApiParams{
		GroupId: opts.groupId,
	}
	_, err := opts.client.MaintenanceWindowsApi.DeferMaintenanceWindowWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(nil)
}

func DeferMaintenanceWindowBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeferMaintenanceWindowOpts{}
	cmd := &cobra.Command{
		Use: "deferMaintenanceWindow",
		Short: "Defer One Maintenance Window for One Project",
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
type GetMaintenanceWindowOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *GetMaintenanceWindowOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetMaintenanceWindowOpts) Run(ctx context.Context) error {
	params := &admin.GetMaintenanceWindowApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.MaintenanceWindowsApi.GetMaintenanceWindowWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetMaintenanceWindowBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetMaintenanceWindowOpts{}
	cmd := &cobra.Command{
		Use: "getMaintenanceWindow",
		Short: "Return One Maintenance Window for One Project",
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
type ResetMaintenanceWindowOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *ResetMaintenanceWindowOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ResetMaintenanceWindowOpts) Run(ctx context.Context) error {
	params := &admin.ResetMaintenanceWindowApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.MaintenanceWindowsApi.ResetMaintenanceWindowWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ResetMaintenanceWindowBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ResetMaintenanceWindowOpts{}
	cmd := &cobra.Command{
		Use: "resetMaintenanceWindow",
		Short: "Reset One Maintenance Window for One Project",
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
type ToggleMaintenanceAutoDeferOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
}

func (opts *ToggleMaintenanceAutoDeferOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ToggleMaintenanceAutoDeferOpts) Run(ctx context.Context) error {
	params := &admin.ToggleMaintenanceAutoDeferApiParams{
		GroupId: opts.groupId,
	}
	_, err := opts.client.MaintenanceWindowsApi.ToggleMaintenanceAutoDeferWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(nil)
}

func ToggleMaintenanceAutoDeferBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ToggleMaintenanceAutoDeferOpts{}
	cmd := &cobra.Command{
		Use: "toggleMaintenanceAutoDefer",
		Short: "Toggle Automatic Deferral of Maintenance for One Project",
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
type UpdateMaintenanceWindowOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	
}

func (opts *UpdateMaintenanceWindowOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateMaintenanceWindowOpts) Run(ctx context.Context) error {
	params := &admin.UpdateMaintenanceWindowApiParams{
		GroupId: opts.groupId,
		
	}
	resp, _, err := opts.client.MaintenanceWindowsApi.UpdateMaintenanceWindowWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdateMaintenanceWindowBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdateMaintenanceWindowOpts{}
	cmd := &cobra.Command{
		Use: "updateMaintenanceWindow",
		Short: "Update Maintenance Window for One Project",
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

func MaintenanceWindowsBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "maintenanceWindows",
		Short:   "Returns, edits, and removes maintenance windows. The maintenance procedure that MongoDB Cloud performs requires at least one replica set election during the maintenance window per replica set. You can defer a scheduled maintenance event for a project up to two times. Deferred maintenance events occur during your preferred maintenance window exactly one week after the previously scheduled date and time.",
	}
	cmd.AddCommand(
		DeferMaintenanceWindowBuilder(),
		GetMaintenanceWindowBuilder(),
		ResetMaintenanceWindowBuilder(),
		ToggleMaintenanceAutoDeferBuilder(),
		UpdateMaintenanceWindowBuilder(),
	)
	return cmd
}

