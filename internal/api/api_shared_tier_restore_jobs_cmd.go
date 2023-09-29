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

type createSharedClusterBackupRestoreJobOpts struct {
	client      *admin.APIClient
	clusterName string
	groupId     string

	filename string
	fs       afero.Fs
}

func (opts *createSharedClusterBackupRestoreJobOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *createSharedClusterBackupRestoreJobOpts) readData() (*admin.TenantRestore, error) {
	var out *admin.TenantRestore

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

func (opts *createSharedClusterBackupRestoreJobOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.CreateSharedClusterBackupRestoreJobApiParams{
		ClusterName: opts.clusterName,
		GroupId:     opts.groupId,

		TenantRestore: data,
	}

	resp, _, err := opts.client.SharedTierRestoreJobsApi.CreateSharedClusterBackupRestoreJobWithParams(ctx, params).Execute()
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

func createSharedClusterBackupRestoreJobBuilder() *cobra.Command {
	opts := createSharedClusterBackupRestoreJobOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createSharedClusterBackupRestoreJob",
		Short: "Create One Restore Job from One M2 or M5 Cluster",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type getSharedClusterBackupRestoreJobOpts struct {
	client      *admin.APIClient
	clusterName string
	groupId     string
	restoreId   string
}

func (opts *getSharedClusterBackupRestoreJobOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getSharedClusterBackupRestoreJobOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.GetSharedClusterBackupRestoreJobApiParams{
		ClusterName: opts.clusterName,
		GroupId:     opts.groupId,
		RestoreId:   opts.restoreId,
	}

	resp, _, err := opts.client.SharedTierRestoreJobsApi.GetSharedClusterBackupRestoreJobWithParams(ctx, params).Execute()
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

func getSharedClusterBackupRestoreJobBuilder() *cobra.Command {
	opts := getSharedClusterBackupRestoreJobOpts{}
	cmd := &cobra.Command{
		Use:   "getSharedClusterBackupRestoreJob",
		Short: "Return One Restore Job for One M2 or M5 Cluster",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.restoreId, "restoreId", "", `Unique 24-hexadecimal digit string that identifies the restore job to return.`)

	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("restoreId")
	return cmd
}

type listSharedClusterBackupRestoreJobsOpts struct {
	client      *admin.APIClient
	clusterName string
	groupId     string
}

func (opts *listSharedClusterBackupRestoreJobsOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listSharedClusterBackupRestoreJobsOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.ListSharedClusterBackupRestoreJobsApiParams{
		ClusterName: opts.clusterName,
		GroupId:     opts.groupId,
	}

	resp, _, err := opts.client.SharedTierRestoreJobsApi.ListSharedClusterBackupRestoreJobsWithParams(ctx, params).Execute()
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

func listSharedClusterBackupRestoreJobsBuilder() *cobra.Command {
	opts := listSharedClusterBackupRestoreJobsOpts{}
	cmd := &cobra.Command{
		Use:   "listSharedClusterBackupRestoreJobs",
		Short: "Return All Restore Jobs for One M2 or M5 Cluster",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

func sharedTierRestoreJobsBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sharedTierRestoreJobs",
		Short: `Returns and adds restore jobs for shared-tier database deployments.`,
	}
	cmd.AddCommand(
		createSharedClusterBackupRestoreJobBuilder(),
		getSharedClusterBackupRestoreJobBuilder(),
		listSharedClusterBackupRestoreJobsBuilder(),
	)
	return cmd
}
