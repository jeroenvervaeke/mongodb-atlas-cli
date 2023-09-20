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

type createPrivateEndpointOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client            *admin.APIClient
	groupId           string
	cloudProvider     string
	endpointServiceId string

	filename string
	fs       afero.Fs
}

func (opts *createPrivateEndpointOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createPrivateEndpointOpts) readData() (*admin.CreateEndpointRequest, error) {
	var out *admin.CreateEndpointRequest

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

func (opts *createPrivateEndpointOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.CreatePrivateEndpointApiParams{
		GroupId:           opts.groupId,
		CloudProvider:     opts.cloudProvider,
		EndpointServiceId: opts.endpointServiceId,

		CreateEndpointRequest: data,
	}
	resp, _, err := opts.client.PrivateEndpointServicesApi.CreatePrivateEndpointWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func createPrivateEndpointBuilder() *cobra.Command {
	opts := createPrivateEndpointOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createPrivateEndpoint",
		Short: "Create One Private Endpoint for One Provider",
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
	cmd.Flags().StringVar(&opts.cloudProvider, "cloudProvider", "&quot;AWS&quot;", `Cloud service provider that manages this private endpoint.`)
	cmd.Flags().StringVar(&opts.endpointServiceId, "endpointServiceId", "", `Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to create a private endpoint.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("cloudProvider")
	_ = cmd.MarkFlagRequired("endpointServiceId")
	return cmd
}

type createPrivateEndpointServiceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *createPrivateEndpointServiceOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createPrivateEndpointServiceOpts) readData() (*admin.CloudProviderEndpointServiceRequest, error) {
	var out *admin.CloudProviderEndpointServiceRequest

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

func (opts *createPrivateEndpointServiceOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.CreatePrivateEndpointServiceApiParams{
		GroupId: opts.groupId,

		CloudProviderEndpointServiceRequest: data,
	}
	resp, _, err := opts.client.PrivateEndpointServicesApi.CreatePrivateEndpointServiceWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func createPrivateEndpointServiceBuilder() *cobra.Command {
	opts := createPrivateEndpointServiceOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createPrivateEndpointService",
		Short: "Create One Private Endpoint Service for One Provider",
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

type deletePrivateEndpointOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client            *admin.APIClient
	groupId           string
	cloudProvider     string
	endpointId        string
	endpointServiceId string
}

func (opts *deletePrivateEndpointOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deletePrivateEndpointOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.DeletePrivateEndpointApiParams{
		GroupId:           opts.groupId,
		CloudProvider:     opts.cloudProvider,
		EndpointId:        opts.endpointId,
		EndpointServiceId: opts.endpointServiceId,
	}
	resp, _, err := opts.client.PrivateEndpointServicesApi.DeletePrivateEndpointWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func deletePrivateEndpointBuilder() *cobra.Command {
	opts := deletePrivateEndpointOpts{}
	cmd := &cobra.Command{
		Use:   "deletePrivateEndpoint",
		Short: "Remove One Private Endpoint for One Provider",
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
	cmd.Flags().StringVar(&opts.cloudProvider, "cloudProvider", "&quot;AWS&quot;", `Cloud service provider that manages this private endpoint.`)
	cmd.Flags().StringVar(&opts.endpointId, "endpointId", "", `Unique string that identifies the private endpoint you want to delete. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints.`)
	cmd.Flags().StringVar(&opts.endpointServiceId, "endpointServiceId", "", `Unique 24-hexadecimal digit string that identifies the private endpoint service from which you want to delete a private endpoint.`)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("cloudProvider")
	_ = cmd.MarkFlagRequired("endpointId")
	_ = cmd.MarkFlagRequired("endpointServiceId")
	return cmd
}

type deletePrivateEndpointServiceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client            *admin.APIClient
	groupId           string
	cloudProvider     string
	endpointServiceId string
}

func (opts *deletePrivateEndpointServiceOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deletePrivateEndpointServiceOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.DeletePrivateEndpointServiceApiParams{
		GroupId:           opts.groupId,
		CloudProvider:     opts.cloudProvider,
		EndpointServiceId: opts.endpointServiceId,
	}
	resp, _, err := opts.client.PrivateEndpointServicesApi.DeletePrivateEndpointServiceWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func deletePrivateEndpointServiceBuilder() *cobra.Command {
	opts := deletePrivateEndpointServiceOpts{}
	cmd := &cobra.Command{
		Use:   "deletePrivateEndpointService",
		Short: "Remove One Private Endpoint Service for One Provider",
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
	cmd.Flags().StringVar(&opts.cloudProvider, "cloudProvider", "&quot;AWS&quot;", `Cloud service provider that manages this private endpoint service.`)
	cmd.Flags().StringVar(&opts.endpointServiceId, "endpointServiceId", "", `Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to delete.`)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("cloudProvider")
	_ = cmd.MarkFlagRequired("endpointServiceId")
	return cmd
}

type getPrivateEndpointOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client            *admin.APIClient
	groupId           string
	cloudProvider     string
	endpointId        string
	endpointServiceId string
}

func (opts *getPrivateEndpointOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getPrivateEndpointOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetPrivateEndpointApiParams{
		GroupId:           opts.groupId,
		CloudProvider:     opts.cloudProvider,
		EndpointId:        opts.endpointId,
		EndpointServiceId: opts.endpointServiceId,
	}
	resp, _, err := opts.client.PrivateEndpointServicesApi.GetPrivateEndpointWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getPrivateEndpointBuilder() *cobra.Command {
	opts := getPrivateEndpointOpts{}
	cmd := &cobra.Command{
		Use:   "getPrivateEndpoint",
		Short: "Return One Private Endpoint for One Provider",
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
	cmd.Flags().StringVar(&opts.cloudProvider, "cloudProvider", "&quot;AWS&quot;", `Cloud service provider that manages this private endpoint.`)
	cmd.Flags().StringVar(&opts.endpointId, "endpointId", "", `Unique string that identifies the private endpoint you want to return. The format of the **endpointId** parameter differs for AWS and Azure. You must URL encode the **endpointId** for Azure private endpoints.`)
	cmd.Flags().StringVar(&opts.endpointServiceId, "endpointServiceId", "", `Unique 24-hexadecimal digit string that identifies the private endpoint service for which you want to return a private endpoint.`)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("cloudProvider")
	_ = cmd.MarkFlagRequired("endpointId")
	_ = cmd.MarkFlagRequired("endpointServiceId")
	return cmd
}

type getPrivateEndpointServiceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client            *admin.APIClient
	groupId           string
	cloudProvider     string
	endpointServiceId string
}

func (opts *getPrivateEndpointServiceOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getPrivateEndpointServiceOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetPrivateEndpointServiceApiParams{
		GroupId:           opts.groupId,
		CloudProvider:     opts.cloudProvider,
		EndpointServiceId: opts.endpointServiceId,
	}
	resp, _, err := opts.client.PrivateEndpointServicesApi.GetPrivateEndpointServiceWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getPrivateEndpointServiceBuilder() *cobra.Command {
	opts := getPrivateEndpointServiceOpts{}
	cmd := &cobra.Command{
		Use:   "getPrivateEndpointService",
		Short: "Return One Private Endpoint Service for One Provider",
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
	cmd.Flags().StringVar(&opts.cloudProvider, "cloudProvider", "&quot;AWS&quot;", `Cloud service provider that manages this private endpoint service.`)
	cmd.Flags().StringVar(&opts.endpointServiceId, "endpointServiceId", "", `Unique 24-hexadecimal digit string that identifies the private endpoint service that you want to return.`)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("cloudProvider")
	_ = cmd.MarkFlagRequired("endpointServiceId")
	return cmd
}

type getRegionalizedPrivateEndpointSettingOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string
}

func (opts *getRegionalizedPrivateEndpointSettingOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getRegionalizedPrivateEndpointSettingOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetRegionalizedPrivateEndpointSettingApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.PrivateEndpointServicesApi.GetRegionalizedPrivateEndpointSettingWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getRegionalizedPrivateEndpointSettingBuilder() *cobra.Command {
	opts := getRegionalizedPrivateEndpointSettingOpts{}
	cmd := &cobra.Command{
		Use:   "getRegionalizedPrivateEndpointSetting",
		Short: "Return Regionalized Private Endpoint Status",
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

type listPrivateEndpointServicesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client        *admin.APIClient
	groupId       string
	cloudProvider string
}

func (opts *listPrivateEndpointServicesOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listPrivateEndpointServicesOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.ListPrivateEndpointServicesApiParams{
		GroupId:       opts.groupId,
		CloudProvider: opts.cloudProvider,
	}
	resp, _, err := opts.client.PrivateEndpointServicesApi.ListPrivateEndpointServicesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func listPrivateEndpointServicesBuilder() *cobra.Command {
	opts := listPrivateEndpointServicesOpts{}
	cmd := &cobra.Command{
		Use:   "listPrivateEndpointServices",
		Short: "Return All Private Endpoint Services for One Provider",
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
	cmd.Flags().StringVar(&opts.cloudProvider, "cloudProvider", "&quot;AWS&quot;", `Cloud service provider that manages this private endpoint service.`)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("cloudProvider")
	return cmd
}

type toggleRegionalizedPrivateEndpointSettingOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *toggleRegionalizedPrivateEndpointSettingOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *toggleRegionalizedPrivateEndpointSettingOpts) readData() (*admin.ProjectSettingItem, error) {
	var out *admin.ProjectSettingItem

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

func (opts *toggleRegionalizedPrivateEndpointSettingOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.ToggleRegionalizedPrivateEndpointSettingApiParams{
		GroupId: opts.groupId,

		ProjectSettingItem: data,
	}
	resp, _, err := opts.client.PrivateEndpointServicesApi.ToggleRegionalizedPrivateEndpointSettingWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func toggleRegionalizedPrivateEndpointSettingBuilder() *cobra.Command {
	opts := toggleRegionalizedPrivateEndpointSettingOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "toggleRegionalizedPrivateEndpointSetting",
		Short: "Toggle Regionalized Private Endpoint Status",
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

func privateEndpointServicesBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "privateEndpointServices",
		Short: `Returns, adds, edits, and removes private endpoint services.`,
	}
	cmd.AddCommand(
		createPrivateEndpointBuilder(),
		createPrivateEndpointServiceBuilder(),
		deletePrivateEndpointBuilder(),
		deletePrivateEndpointServiceBuilder(),
		getPrivateEndpointBuilder(),
		getPrivateEndpointServiceBuilder(),
		getRegionalizedPrivateEndpointSettingBuilder(),
		listPrivateEndpointServicesBuilder(),
		toggleRegionalizedPrivateEndpointSettingBuilder(),
	)
	return cmd
}
