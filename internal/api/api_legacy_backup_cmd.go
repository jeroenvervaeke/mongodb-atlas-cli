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

type deleteLegacySnapshotOpts struct {
	client      *admin.APIClient
	groupId     string
	clusterName string
	snapshotId  string
}

func (opts *deleteLegacySnapshotOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *deleteLegacySnapshotOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.DeleteLegacySnapshotApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
		SnapshotId:  opts.snapshotId,
	}

	resp, _, err := opts.client.LegacyBackupApi.DeleteLegacySnapshotWithParams(ctx, params).Execute()
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

func deleteLegacySnapshotBuilder() *cobra.Command {
	opts := deleteLegacySnapshotOpts{}
	cmd := &cobra.Command{
		Use:   "deleteLegacySnapshot",
		Short: "Remove One Legacy Backup Snapshot",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)
	cmd.Flags().StringVar(&opts.snapshotId, "snapshotId", "", `Unique 24-hexadecimal digit string that identifies the desired snapshot.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("snapshotId")
	return cmd
}

type getLegacyBackupCheckpointOpts struct {
	client       *admin.APIClient
	groupId      string
	checkpointId string
	clusterName  string
}

func (opts *getLegacyBackupCheckpointOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getLegacyBackupCheckpointOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.GetLegacyBackupCheckpointApiParams{
		GroupId:      opts.groupId,
		CheckpointId: opts.checkpointId,
		ClusterName:  opts.clusterName,
	}

	resp, _, err := opts.client.LegacyBackupApi.GetLegacyBackupCheckpointWithParams(ctx, params).Execute()
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

func getLegacyBackupCheckpointBuilder() *cobra.Command {
	opts := getLegacyBackupCheckpointOpts{}
	cmd := &cobra.Command{
		Use:   "getLegacyBackupCheckpoint",
		Short: "Return One Legacy Backup Checkpoint",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.checkpointId, "checkpointId", "", `Unique 24-hexadecimal digit string that identifies the checkpoint.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster that contains the checkpoints that you want to return.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("checkpointId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type getLegacyBackupRestoreJobOpts struct {
	client      *admin.APIClient
	groupId     string
	clusterName string
	jobId       string
}

func (opts *getLegacyBackupRestoreJobOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getLegacyBackupRestoreJobOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.GetLegacyBackupRestoreJobApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
		JobId:       opts.jobId,
	}

	resp, _, err := opts.client.LegacyBackupApi.GetLegacyBackupRestoreJobWithParams(ctx, params).Execute()
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

func getLegacyBackupRestoreJobBuilder() *cobra.Command {
	opts := getLegacyBackupRestoreJobOpts{}
	cmd := &cobra.Command{
		Use:   "getLegacyBackupRestoreJob",
		Short: "Return One Legacy Backup Restore Job",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster with the snapshot you want to return.`)
	cmd.Flags().StringVar(&opts.jobId, "jobId", "", `Unique 24-hexadecimal digit string that identifies the restore job.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("jobId")
	return cmd
}

type getLegacySnapshotOpts struct {
	client      *admin.APIClient
	groupId     string
	clusterName string
	snapshotId  string
}

func (opts *getLegacySnapshotOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getLegacySnapshotOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.GetLegacySnapshotApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
		SnapshotId:  opts.snapshotId,
	}

	resp, _, err := opts.client.LegacyBackupApi.GetLegacySnapshotWithParams(ctx, params).Execute()
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

func getLegacySnapshotBuilder() *cobra.Command {
	opts := getLegacySnapshotOpts{}
	cmd := &cobra.Command{
		Use:   "getLegacySnapshot",
		Short: "Return One Legacy Backup Snapshot",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)
	cmd.Flags().StringVar(&opts.snapshotId, "snapshotId", "", `Unique 24-hexadecimal digit string that identifies the desired snapshot.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("snapshotId")
	return cmd
}

type getLegacySnapshotScheduleOpts struct {
	client      *admin.APIClient
	groupId     string
	clusterName string
}

func (opts *getLegacySnapshotScheduleOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getLegacySnapshotScheduleOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.GetLegacySnapshotScheduleApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
	}

	resp, _, err := opts.client.LegacyBackupApi.GetLegacySnapshotScheduleWithParams(ctx, params).Execute()
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

func getLegacySnapshotScheduleBuilder() *cobra.Command {
	opts := getLegacySnapshotScheduleOpts{}
	cmd := &cobra.Command{
		Use:   "getLegacySnapshotSchedule",
		Short: "Return One Snapshot Schedule",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster with the snapshot you want to return.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type listLegacyBackupCheckpointsOpts struct {
	client       *admin.APIClient
	groupId      string
	clusterName  string
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listLegacyBackupCheckpointsOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listLegacyBackupCheckpointsOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.ListLegacyBackupCheckpointsApiParams{
		GroupId:      opts.groupId,
		ClusterName:  opts.clusterName,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}

	resp, _, err := opts.client.LegacyBackupApi.ListLegacyBackupCheckpointsWithParams(ctx, params).Execute()
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

func listLegacyBackupCheckpointsBuilder() *cobra.Command {
	opts := listLegacyBackupCheckpointsOpts{}
	cmd := &cobra.Command{
		Use:   "listLegacyBackupCheckpoints",
		Short: "Return All Legacy Backup Checkpoints",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster that contains the checkpoints that you want to return.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type listLegacyBackupRestoreJobsOpts struct {
	client       *admin.APIClient
	groupId      string
	clusterName  string
	includeCount bool
	itemsPerPage int
	pageNum      int
	batchId      string
}

func (opts *listLegacyBackupRestoreJobsOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listLegacyBackupRestoreJobsOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.ListLegacyBackupRestoreJobsApiParams{
		GroupId:      opts.groupId,
		ClusterName:  opts.clusterName,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
		BatchId:      &opts.batchId,
	}

	resp, _, err := opts.client.LegacyBackupApi.ListLegacyBackupRestoreJobsWithParams(ctx, params).Execute()
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

func listLegacyBackupRestoreJobsBuilder() *cobra.Command {
	opts := listLegacyBackupRestoreJobsOpts{}
	cmd := &cobra.Command{
		Use:   "listLegacyBackupRestoreJobs",
		Short: "Return All Legacy Backup Restore Jobs",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster with the snapshot you want to return.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)
	cmd.Flags().StringVar(&opts.batchId, "batchId", "", `Unique 24-hexadecimal digit string that identifies the batch of restore jobs to return. Timestamp in ISO 8601 date and time format in UTC when creating a restore job for a sharded cluster, Application creates a separate job for each shard, plus another for the config host. Each of these jobs comprise one batch. A restore job for a replica set can&#39;t be part of a batch.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type listLegacySnapshotsOpts struct {
	client       *admin.APIClient
	groupId      string
	clusterName  string
	includeCount bool
	itemsPerPage int
	pageNum      int
	completed    string
}

func (opts *listLegacySnapshotsOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listLegacySnapshotsOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.ListLegacySnapshotsApiParams{
		GroupId:      opts.groupId,
		ClusterName:  opts.clusterName,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
		Completed:    &opts.completed,
	}

	resp, _, err := opts.client.LegacyBackupApi.ListLegacySnapshotsWithParams(ctx, params).Execute()
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

func listLegacySnapshotsBuilder() *cobra.Command {
	opts := listLegacySnapshotsOpts{}
	cmd := &cobra.Command{
		Use:   "listLegacySnapshots",
		Short: "Return All Legacy Backup Snapshots",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)
	cmd.Flags().StringVar(&opts.completed, "completed", "&quot;true&quot;", `Human-readable label that specifies whether to return only completed, incomplete, or all snapshots. By default, MongoDB Cloud only returns completed snapshots.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type updateLegacySnapshotRetentionOpts struct {
	client      *admin.APIClient
	groupId     string
	clusterName string
	snapshotId  string

	filename string
	fs       afero.Fs
}

func (opts *updateLegacySnapshotRetentionOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *updateLegacySnapshotRetentionOpts) readData() (*admin.BackupSnapshot, error) {
	var out *admin.BackupSnapshot

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

func (opts *updateLegacySnapshotRetentionOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.UpdateLegacySnapshotRetentionApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
		SnapshotId:  opts.snapshotId,

		BackupSnapshot: data,
	}

	resp, _, err := opts.client.LegacyBackupApi.UpdateLegacySnapshotRetentionWithParams(ctx, params).Execute()
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

func updateLegacySnapshotRetentionBuilder() *cobra.Command {
	opts := updateLegacySnapshotRetentionOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateLegacySnapshotRetention",
		Short: "Change One Legacy Backup Snapshot Expiration",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)
	cmd.Flags().StringVar(&opts.snapshotId, "snapshotId", "", `Unique 24-hexadecimal digit string that identifies the desired snapshot.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("snapshotId")
	return cmd
}

type updateLegacySnapshotScheduleOpts struct {
	client      *admin.APIClient
	groupId     string
	clusterName string

	filename string
	fs       afero.Fs
}

func (opts *updateLegacySnapshotScheduleOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *updateLegacySnapshotScheduleOpts) readData() (*admin.ApiAtlasSnapshotSchedule, error) {
	var out *admin.ApiAtlasSnapshotSchedule

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

func (opts *updateLegacySnapshotScheduleOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.UpdateLegacySnapshotScheduleApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,

		ApiAtlasSnapshotSchedule: data,
	}

	resp, _, err := opts.client.LegacyBackupApi.UpdateLegacySnapshotScheduleWithParams(ctx, params).Execute()
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

func updateLegacySnapshotScheduleBuilder() *cobra.Command {
	opts := updateLegacySnapshotScheduleOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateLegacySnapshotSchedule",
		Short: "Update Snapshot Schedule for One Cluster",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster with the snapshot you want to return.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

func legacyBackupBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "legacyBackup",
		Short: `Manages Legacy Backup snapshots, restore jobs, schedules and checkpoints.`,
	}
	cmd.AddCommand(
		deleteLegacySnapshotBuilder(),
		getLegacyBackupCheckpointBuilder(),
		getLegacyBackupRestoreJobBuilder(),
		getLegacySnapshotBuilder(),
		getLegacySnapshotScheduleBuilder(),
		listLegacyBackupCheckpointsBuilder(),
		listLegacyBackupRestoreJobsBuilder(),
		listLegacySnapshotsBuilder(),
		updateLegacySnapshotRetentionBuilder(),
		updateLegacySnapshotScheduleBuilder(),
	)
	return cmd
}
