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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/jsonwriter"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type getAuditingConfigurationOpts struct {
	cli.GlobalOpts
	client  *admin.APIClient
	groupId string
}

func (opts *getAuditingConfigurationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getAuditingConfigurationOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetAuditingConfigurationApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.AuditingApi.GetAuditingConfigurationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getAuditingConfigurationBuilder() *cobra.Command {
	opts := getAuditingConfigurationOpts{}
	cmd := &cobra.Command{
		Use:   "getAuditingConfiguration",
		Short: "Return the Auditing Configuration for One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type updateAuditingConfigurationOpts struct {
	cli.GlobalOpts
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *updateAuditingConfigurationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updateAuditingConfigurationOpts) readData() (*admin.AuditLog, error) {
	var out *admin.AuditLog

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

func (opts *updateAuditingConfigurationOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.UpdateAuditingConfigurationApiParams{
		GroupId: opts.groupId,

		AuditLog: data,
	}
	resp, _, err := opts.client.AuditingApi.UpdateAuditingConfigurationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func updateAuditingConfigurationBuilder() *cobra.Command {
	opts := updateAuditingConfigurationOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateAuditingConfiguration",
		Short: "Update Auditing Configuration for One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

func auditingBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auditing",
		Short: `Returns and edits database auditing settings for MongoDB Cloud projects.`,
	}
	cmd.AddCommand(
		getAuditingConfigurationBuilder(),
		updateAuditingConfigurationBuilder(),
	)
	return cmd
}
