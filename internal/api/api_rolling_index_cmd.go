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

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type createRollingIndexOpts struct {
	client      *admin.APIClient
	groupId     string
	clusterName string

	filename string
	fs       afero.Fs
}

func (opts *createRollingIndexOpts) preRun() (err error) {
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

	return err
}

func (opts *createRollingIndexOpts) readData(r io.Reader) (*admin.DatabaseRollingIndexRequest, error) {
	var out *admin.DatabaseRollingIndexRequest

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

func (opts *createRollingIndexOpts) run(ctx context.Context, r io.Reader, _ io.Writer) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.CreateRollingIndexApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,

		DatabaseRollingIndexRequest: data,
	}

	_, err := opts.client.RollingIndexApi.CreateRollingIndexWithParams(ctx, params).Execute()
	return err
}

func createRollingIndexBuilder() *cobra.Command {
	opts := createRollingIndexOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createRollingIndex",
		Short: "Create One Rolling Index",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster on which MongoDB Cloud creates an index.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

func rollingIndexBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rollingIndex",
		Short: `Creates one index to a database deployment in a rolling manner. You can&#39;t create a rolling index on an &#x60;M0&#x60; free cluster or &#x60;M2/M5&#x60; shared cluster.`,
	}
	cmd.AddCommand(
		createRollingIndexBuilder(),
	)
	return cmd
}
