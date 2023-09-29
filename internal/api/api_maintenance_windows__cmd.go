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

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type deferMaintenanceWindowOpts struct {
	client  *admin.APIClient
	groupId string
}

func (opts *deferMaintenanceWindowOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *deferMaintenanceWindowOpts) run(ctx context.Context, _ io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.DeferMaintenanceWindowApiParams{
		GroupId: opts.groupId,
	}

	_, err := opts.client.MaintenanceWindowsApi.DeferMaintenanceWindowWithParams(ctx, params).Execute()
	return err
}

func deferMaintenanceWindowBuilder() *cobra.Command {
	opts := deferMaintenanceWindowOpts{}
	cmd := &cobra.Command{
		Use:   "deferMaintenanceWindow",
		Short: "Defer One Maintenance Window for One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type getMaintenanceWindowOpts struct {
	client  *admin.APIClient
	groupId string
}

func (opts *getMaintenanceWindowOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getMaintenanceWindowOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.GetMaintenanceWindowApiParams{
		GroupId: opts.groupId,
	}

	resp, _, err := opts.client.MaintenanceWindowsApi.GetMaintenanceWindowWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func getMaintenanceWindowBuilder() *cobra.Command {
	opts := getMaintenanceWindowOpts{}
	cmd := &cobra.Command{
		Use:   "getMaintenanceWindow",
		Short: "Return One Maintenance Window for One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type resetMaintenanceWindowOpts struct {
	client  *admin.APIClient
	groupId string
}

func (opts *resetMaintenanceWindowOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *resetMaintenanceWindowOpts) run(ctx context.Context, _ io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.ResetMaintenanceWindowApiParams{
		GroupId: opts.groupId,
	}

	_, err := opts.client.MaintenanceWindowsApi.ResetMaintenanceWindowWithParams(ctx, params).Execute()
	return err
}

func resetMaintenanceWindowBuilder() *cobra.Command {
	opts := resetMaintenanceWindowOpts{}
	cmd := &cobra.Command{
		Use:   "resetMaintenanceWindow",
		Short: "Reset One Maintenance Window for One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type toggleMaintenanceAutoDeferOpts struct {
	client  *admin.APIClient
	groupId string
}

func (opts *toggleMaintenanceAutoDeferOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *toggleMaintenanceAutoDeferOpts) run(ctx context.Context, _ io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.ToggleMaintenanceAutoDeferApiParams{
		GroupId: opts.groupId,
	}

	_, err := opts.client.MaintenanceWindowsApi.ToggleMaintenanceAutoDeferWithParams(ctx, params).Execute()
	return err
}

func toggleMaintenanceAutoDeferBuilder() *cobra.Command {
	opts := toggleMaintenanceAutoDeferOpts{}
	cmd := &cobra.Command{
		Use:   "toggleMaintenanceAutoDefer",
		Short: "Toggle Automatic Deferral of Maintenance for One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type updateMaintenanceWindowOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *updateMaintenanceWindowOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *updateMaintenanceWindowOpts) readData() (*admin.GroupMaintenanceWindow, error) {
	var out *admin.GroupMaintenanceWindow

	var buf []byte
	var err error
	if opts.filename == "" {
		buf, err = io.ReadAll(os.Stdin)
	} else {
		if exists, errExists := afero.Exists(opts.fs, opts.filename); !exists || errExists != nil {
			return nil, fmt.Errorf("file not found: %s", opts.filename)
		}
		buf, err = afero.ReadFile(opts.fs, opts.filename)
	}
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(buf, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (opts *updateMaintenanceWindowOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.UpdateMaintenanceWindowApiParams{
		GroupId: opts.groupId,

		GroupMaintenanceWindow: data,
	}

	resp, _, err := opts.client.MaintenanceWindowsApi.UpdateMaintenanceWindowWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func updateMaintenanceWindowBuilder() *cobra.Command {
	opts := updateMaintenanceWindowOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateMaintenanceWindow",
		Short: "Update Maintenance Window for One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

func maintenanceWindowsBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "maintenanceWindows",
		Short: `Returns, edits, and removes maintenance windows. The maintenance procedure that MongoDB Cloud performs requires at least one replica set election during the maintenance window per replica set. You can defer a scheduled maintenance event for a project up to two times. Deferred maintenance events occur during your preferred maintenance window exactly one week after the previously scheduled date and time.`,
	}
	cmd.AddCommand(
		deferMaintenanceWindowBuilder(),
		getMaintenanceWindowBuilder(),
		resetMaintenanceWindowBuilder(),
		toggleMaintenanceAutoDeferBuilder(),
		updateMaintenanceWindowBuilder(),
	)
	return cmd
}
