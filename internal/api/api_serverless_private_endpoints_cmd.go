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

type createServerlessPrivateEndpointOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	instanceName string

	filename string
	fs       afero.Fs
}

func (opts *createServerlessPrivateEndpointOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createServerlessPrivateEndpointOpts) readData() (*admin.ServerlessTenantCreateRequest, error) {
	var out *admin.ServerlessTenantCreateRequest

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

func (opts *createServerlessPrivateEndpointOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.CreateServerlessPrivateEndpointApiParams{
		GroupId:      opts.groupId,
		InstanceName: opts.instanceName,

		ServerlessTenantCreateRequest: data,
	}
	resp, _, err := opts.client.ServerlessPrivateEndpointsApi.CreateServerlessPrivateEndpointWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func createServerlessPrivateEndpointBuilder() *cobra.Command {
	opts := createServerlessPrivateEndpointOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createServerlessPrivateEndpoint",
		Short: "Create One Private Endpoint for One Serverless Instance",
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
	cmd.Flags().StringVar(&opts.instanceName, "instanceName", "", `Human-readable label that identifies the serverless instance for which the tenant endpoint will be created.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("instanceName")
	return cmd
}

type deleteServerlessPrivateEndpointOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	instanceName string
	endpointId   string
}

func (opts *deleteServerlessPrivateEndpointOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deleteServerlessPrivateEndpointOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.DeleteServerlessPrivateEndpointApiParams{
		GroupId:      opts.groupId,
		InstanceName: opts.instanceName,
		EndpointId:   opts.endpointId,
	}
	resp, _, err := opts.client.ServerlessPrivateEndpointsApi.DeleteServerlessPrivateEndpointWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func deleteServerlessPrivateEndpointBuilder() *cobra.Command {
	opts := deleteServerlessPrivateEndpointOpts{}
	cmd := &cobra.Command{
		Use:   "deleteServerlessPrivateEndpoint",
		Short: "Remove One Private Endpoint for One Serverless Instance",
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
	cmd.Flags().StringVar(&opts.instanceName, "instanceName", "", `Human-readable label that identifies the serverless instance from which the tenant endpoint will be removed.`)
	cmd.Flags().StringVar(&opts.endpointId, "endpointId", "", `Unique 24-hexadecimal digit string that identifies the tenant endpoint which will be removed.`)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("instanceName")
	_ = cmd.MarkFlagRequired("endpointId")
	return cmd
}

type getServerlessPrivateEndpointOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	instanceName string
	endpointId   string
}

func (opts *getServerlessPrivateEndpointOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getServerlessPrivateEndpointOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetServerlessPrivateEndpointApiParams{
		GroupId:      opts.groupId,
		InstanceName: opts.instanceName,
		EndpointId:   opts.endpointId,
	}
	resp, _, err := opts.client.ServerlessPrivateEndpointsApi.GetServerlessPrivateEndpointWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getServerlessPrivateEndpointBuilder() *cobra.Command {
	opts := getServerlessPrivateEndpointOpts{}
	cmd := &cobra.Command{
		Use:   "getServerlessPrivateEndpoint",
		Short: "Return One Private Endpoint for One Serverless Instance",
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
	cmd.Flags().StringVar(&opts.instanceName, "instanceName", "", `Human-readable label that identifies the serverless instance associated with the tenant endpoint.`)
	cmd.Flags().StringVar(&opts.endpointId, "endpointId", "", `Unique 24-hexadecimal digit string that identifies the tenant endpoint.`)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("instanceName")
	_ = cmd.MarkFlagRequired("endpointId")
	return cmd
}

type listServerlessPrivateEndpointsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	instanceName string
}

func (opts *listServerlessPrivateEndpointsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listServerlessPrivateEndpointsOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.ListServerlessPrivateEndpointsApiParams{
		GroupId:      opts.groupId,
		InstanceName: opts.instanceName,
	}
	resp, _, err := opts.client.ServerlessPrivateEndpointsApi.ListServerlessPrivateEndpointsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func listServerlessPrivateEndpointsBuilder() *cobra.Command {
	opts := listServerlessPrivateEndpointsOpts{}
	cmd := &cobra.Command{
		Use:   "listServerlessPrivateEndpoints",
		Short: "Return All Private Endpoints for One Serverless Instance",
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
	cmd.Flags().StringVar(&opts.instanceName, "instanceName", "", `Human-readable label that identifies the serverless instance associated with the tenant endpoint.`)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("instanceName")
	return cmd
}

type updateServerlessPrivateEndpointOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	instanceName string
	endpointId   string

	filename string
	fs       afero.Fs
}

func (opts *updateServerlessPrivateEndpointOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updateServerlessPrivateEndpointOpts) readData() (*admin.ServerlessTenantEndpointUpdate, error) {
	var out *admin.ServerlessTenantEndpointUpdate

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

func (opts *updateServerlessPrivateEndpointOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.UpdateServerlessPrivateEndpointApiParams{
		GroupId:      opts.groupId,
		InstanceName: opts.instanceName,
		EndpointId:   opts.endpointId,

		ServerlessTenantEndpointUpdate: data,
	}
	resp, _, err := opts.client.ServerlessPrivateEndpointsApi.UpdateServerlessPrivateEndpointWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func updateServerlessPrivateEndpointBuilder() *cobra.Command {
	opts := updateServerlessPrivateEndpointOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateServerlessPrivateEndpoint",
		Short: "Update One Private Endpoint for One Serverless Instance",
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
	cmd.Flags().StringVar(&opts.instanceName, "instanceName", "", `Human-readable label that identifies the serverless instance associated with the tenant endpoint that will be updated.`)
	cmd.Flags().StringVar(&opts.endpointId, "endpointId", "", `Unique 24-hexadecimal digit string that identifies the tenant endpoint which will be updated.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("instanceName")
	_ = cmd.MarkFlagRequired("endpointId")
	return cmd
}

func serverlessPrivateEndpointsBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serverlessPrivateEndpoints",
		Short: `Returns, adds, edits, and removes private endpoints for serverless instances. To learn more, see the Atlas Administration API tab on the following tutorial.`,
	}
	cmd.AddCommand(
		createServerlessPrivateEndpointBuilder(),
		deleteServerlessPrivateEndpointBuilder(),
		getServerlessPrivateEndpointBuilder(),
		listServerlessPrivateEndpointsBuilder(),
		updateServerlessPrivateEndpointBuilder(),
	)
	return cmd
}
