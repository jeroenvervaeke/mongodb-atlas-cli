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
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type createLegacyBackupRestoreJobOpts struct {
	client      *admin.APIClient
	groupId     string
	clusterName string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
}

func (opts *createLegacyBackupRestoreJobOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
		return err
	}

	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}
	if opts.groupId == "" {
		return errors.New(`required flag(s) "projectId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.groupId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.groupId)
	}

	if opts.format != "" {
		opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n")
	}

	return err
}

func (opts *createLegacyBackupRestoreJobOpts) readData(r io.Reader) (*admin.BackupRestoreJob, error) {
	var out *admin.BackupRestoreJob

	var buf []byte
	var err error
	if opts.filename == "" {
		buf, err = io.ReadAll(r)
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

func (opts *createLegacyBackupRestoreJobOpts) run(ctx context.Context, r io.Reader, w io.Writer) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.CreateLegacyBackupRestoreJobApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,

		BackupRestoreJob: data,
	}

	resp, _, err := opts.client.LegacyBackupRestoreJobsApi.CreateLegacyBackupRestoreJobWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err = fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err = json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	err = opts.tmpl.Execute(w, parsedJSON)
	return err
}

func createLegacyBackupRestoreJobBuilder() *cobra.Command {
	opts := createLegacyBackupRestoreJobOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createLegacyBackupRestoreJob",
		Short: "Create One Legacy Backup Restore Job",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster with the snapshot you want to return.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("clusterName")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

func legacyBackupRestoreJobsBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "legacyBackupRestoreJobs",
		Short: ``,
	}
	cmd.AddCommand(
		createLegacyBackupRestoreJobBuilder(),
	)
	return cmd
}
