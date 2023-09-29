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

type createLinkTokenOpts struct {
	client *admin.APIClient
	orgId  string

	filename string
	fs       afero.Fs
}

func (opts *createLinkTokenOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *createLinkTokenOpts) readData() (*admin.TargetOrgRequest, error) {
	var out *admin.TargetOrgRequest

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

func (opts *createLinkTokenOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.CreateLinkTokenApiParams{
		OrgId: opts.orgId,

		TargetOrgRequest: data,
	}

	resp, _, err := opts.client.CloudMigrationServiceApi.CreateLinkTokenWithParams(ctx, params).Execute()
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

func createLinkTokenBuilder() *cobra.Command {
	opts := createLinkTokenOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createLinkToken",
		Short: "Create One Link-Token",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type createPushMigrationOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *createPushMigrationOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *createPushMigrationOpts) readData() (*admin.LiveMigrationRequest, error) {
	var out *admin.LiveMigrationRequest

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

func (opts *createPushMigrationOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.CreatePushMigrationApiParams{
		GroupId: opts.groupId,

		LiveMigrationRequest: data,
	}

	resp, _, err := opts.client.CloudMigrationServiceApi.CreatePushMigrationWithParams(ctx, params).Execute()
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

func createPushMigrationBuilder() *cobra.Command {
	opts := createPushMigrationOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createPushMigration",
		Short: "Migrate One Local Managed Cluster to MongoDB Atlas",
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

type cutoverMigrationOpts struct {
	client          *admin.APIClient
	groupId         string
	liveMigrationId string
}

func (opts *cutoverMigrationOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *cutoverMigrationOpts) run(ctx context.Context, _ io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.CutoverMigrationApiParams{
		GroupId:         opts.groupId,
		LiveMigrationId: opts.liveMigrationId,
	}

	_, err := opts.client.CloudMigrationServiceApi.CutoverMigrationWithParams(ctx, params).Execute()
	return err
}

func cutoverMigrationBuilder() *cobra.Command {
	opts := cutoverMigrationOpts{}
	cmd := &cobra.Command{
		Use:   "cutoverMigration",
		Short: "Cut Over the Migrated Cluster",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.liveMigrationId, "liveMigrationId", "", `Unique 24-hexadecimal digit string that identifies the migration.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("liveMigrationId")
	return cmd
}

type deleteLinkTokenOpts struct {
	client *admin.APIClient
	orgId  string
}

func (opts *deleteLinkTokenOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *deleteLinkTokenOpts) run(ctx context.Context, w io.Writer) error {
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.DeleteLinkTokenApiParams{
		OrgId: opts.orgId,
	}

	resp, _, err := opts.client.CloudMigrationServiceApi.DeleteLinkTokenWithParams(ctx, params).Execute()
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

func deleteLinkTokenBuilder() *cobra.Command {
	opts := deleteLinkTokenOpts{}
	cmd := &cobra.Command{
		Use:   "deleteLinkToken",
		Short: "Remove One Link-Token",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type getPushMigrationOpts struct {
	client          *admin.APIClient
	groupId         string
	liveMigrationId string
}

func (opts *getPushMigrationOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getPushMigrationOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.GetPushMigrationApiParams{
		GroupId:         opts.groupId,
		LiveMigrationId: opts.liveMigrationId,
	}

	resp, _, err := opts.client.CloudMigrationServiceApi.GetPushMigrationWithParams(ctx, params).Execute()
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

func getPushMigrationBuilder() *cobra.Command {
	opts := getPushMigrationOpts{}
	cmd := &cobra.Command{
		Use:   "getPushMigration",
		Short: "Return One Migration Job",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.liveMigrationId, "liveMigrationId", "", `Unique 24-hexadecimal digit string that identifies the migration.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("liveMigrationId")
	return cmd
}

type getValidationStatusOpts struct {
	client       *admin.APIClient
	groupId      string
	validationId string
}

func (opts *getValidationStatusOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getValidationStatusOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.GetValidationStatusApiParams{
		GroupId:      opts.groupId,
		ValidationId: opts.validationId,
	}

	resp, _, err := opts.client.CloudMigrationServiceApi.GetValidationStatusWithParams(ctx, params).Execute()
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

func getValidationStatusBuilder() *cobra.Command {
	opts := getValidationStatusOpts{}
	cmd := &cobra.Command{
		Use:   "getValidationStatus",
		Short: "Return One Migration Validation Job",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.validationId, "validationId", "", `Unique 24-hexadecimal digit string that identifies the validation job.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("validationId")
	return cmd
}

type listSourceProjectsOpts struct {
	client *admin.APIClient
	orgId  string
}

func (opts *listSourceProjectsOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listSourceProjectsOpts) run(ctx context.Context, w io.Writer) error {
	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}

	params := &admin.ListSourceProjectsApiParams{
		OrgId: opts.orgId,
	}

	resp, _, err := opts.client.CloudMigrationServiceApi.ListSourceProjectsWithParams(ctx, params).Execute()
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

func listSourceProjectsBuilder() *cobra.Command {
	opts := listSourceProjectsOpts{}
	cmd := &cobra.Command{
		Use:   "listSourceProjects",
		Short: "Return All Projects Available for Migration",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)

	_ = cmd.MarkFlagRequired("orgId")
	return cmd
}

type validateMigrationOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *validateMigrationOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *validateMigrationOpts) readData() (*admin.LiveMigrationRequest, error) {
	var out *admin.LiveMigrationRequest

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

func (opts *validateMigrationOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.ValidateMigrationApiParams{
		GroupId: opts.groupId,

		LiveMigrationRequest: data,
	}

	resp, _, err := opts.client.CloudMigrationServiceApi.ValidateMigrationWithParams(ctx, params).Execute()
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

func validateMigrationBuilder() *cobra.Command {
	opts := validateMigrationOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "validateMigration",
		Short: "Validate One Migration Request",
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

func cloudMigrationServiceBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cloudMigrationService",
		Short: `Manages the Cloud Migration Service. Source organizations, projects, and MongoDB clusters reside on Cloud Manager or Ops Manager. Destination organizations, projects, and MongoDB clusters reside on MongoDB Cloud. Source databases can&#39;t use any authentication except SCRAM-SHA.`,
	}
	cmd.AddCommand(
		createLinkTokenBuilder(),
		createPushMigrationBuilder(),
		cutoverMigrationBuilder(),
		deleteLinkTokenBuilder(),
		getPushMigrationBuilder(),
		getValidationStatusBuilder(),
		listSourceProjectsBuilder(),
		validateMigrationBuilder(),
	)
	return cmd
}
