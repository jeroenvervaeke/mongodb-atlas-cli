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
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/jsonwriter"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type getAWSCustomDNSOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string
}

func (opts *getAWSCustomDNSOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getAWSCustomDNSOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetAWSCustomDNSApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.AWSClustersDNSApi.GetAWSCustomDNSWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getAWSCustomDNSBuilder() *cobra.Command {
	opts := getAWSCustomDNSOpts{}
	cmd := &cobra.Command{
		Use:   "getAWSCustomDNS",
		Short: "Return One Custom DNS Configuration for Atlas Clusters on AWS",
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
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type toggleAWSCustomDNSOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *toggleAWSCustomDNSOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *toggleAWSCustomDNSOpts) readData() (*admin.AWSCustomDNSEnabled, error) {
	var out *admin.AWSCustomDNSEnabled

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

func (opts *toggleAWSCustomDNSOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.ToggleAWSCustomDNSApiParams{
		GroupId: opts.groupId,

		AWSCustomDNSEnabled: data,
	}
	resp, _, err := opts.client.AWSClustersDNSApi.ToggleAWSCustomDNSWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func toggleAWSCustomDNSBuilder() *cobra.Command {
	opts := toggleAWSCustomDNSOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "toggleAWSCustomDNS",
		Short: "Toggle State of One Custom DNS Configuration for Atlas Clusters on AWS",
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
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

func aWSClustersDNSBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "aWSClustersDNS",
		Short: `Returns and edits custom DNS configurations for MongoDB Cloud database deployments on AWS. The resource requires your Project ID. If you use the VPC peering on AWS and you use your own DNS servers instead of Amazon Route 53, enable custom DNS. Before 31 March 2020, applications deployed within AWS using custom DNS services and VPC-peered with MongoDB Cloud couldn&#39;t connect over private IP addresses. Custom DNS resolved to public IP addresses. AWS internal DNS resolved to private IP addresses. Applications deployed with custom DNS services in AWS should use Private IP for Peering connection strings.`,
	}
	cmd.AddCommand(
		getAWSCustomDNSBuilder(),
		toggleAWSCustomDNSBuilder(),
	)
	return cmd
}
