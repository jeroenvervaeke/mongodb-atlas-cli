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

type createCustomZoneMappingOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	groupId     string
	clusterName string

	filename string
	fs       afero.Fs
}

func (opts *createCustomZoneMappingOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createCustomZoneMappingOpts) readData() (*admin.GeoSharding, error) {
	var out *admin.GeoSharding

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

func (opts *createCustomZoneMappingOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.CreateCustomZoneMappingApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,

		GeoSharding: data,
	}
	resp, _, err := opts.client.GlobalClustersApi.CreateCustomZoneMappingWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func createCustomZoneMappingBuilder() *cobra.Command {
	opts := createCustomZoneMappingOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createCustomZoneMapping",
		Short: "Add One Entry to One Custom Zone Mapping",
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies this advanced cluster.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type createManagedNamespaceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	groupId     string
	clusterName string

	filename string
	fs       afero.Fs
}

func (opts *createManagedNamespaceOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createManagedNamespaceOpts) readData() (*admin.ManagedNamespace, error) {
	var out *admin.ManagedNamespace

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

func (opts *createManagedNamespaceOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.CreateManagedNamespaceApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,

		ManagedNamespace: data,
	}
	resp, _, err := opts.client.GlobalClustersApi.CreateManagedNamespaceWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func createManagedNamespaceBuilder() *cobra.Command {
	opts := createManagedNamespaceOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createManagedNamespace",
		Short: "Create One Managed Namespace in One Global Multi-Cloud Cluster",
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies this advanced cluster.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type deleteAllCustomZoneMappingsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	groupId     string
	clusterName string
}

func (opts *deleteAllCustomZoneMappingsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deleteAllCustomZoneMappingsOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.DeleteAllCustomZoneMappingsApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.GlobalClustersApi.DeleteAllCustomZoneMappingsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func deleteAllCustomZoneMappingsBuilder() *cobra.Command {
	opts := deleteAllCustomZoneMappingsOpts{}
	cmd := &cobra.Command{
		Use:   "deleteAllCustomZoneMappings",
		Short: "Remove All Custom Zone Mappings from One Global Multi-Cloud Cluster",
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies this advanced cluster.`)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type deleteManagedNamespaceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	clusterName string
	groupId     string
	db          string
	collection  string
}

func (opts *deleteManagedNamespaceOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deleteManagedNamespaceOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.DeleteManagedNamespaceApiParams{
		ClusterName: opts.clusterName,
		GroupId:     opts.groupId,
		Db:          &opts.db,
		Collection:  &opts.collection,
	}
	resp, _, err := opts.client.GlobalClustersApi.DeleteManagedNamespaceWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func deleteManagedNamespaceBuilder() *cobra.Command {
	opts := deleteManagedNamespaceOpts{}
	cmd := &cobra.Command{
		Use:   "deleteManagedNamespace",
		Short: "Remove One Managed Namespace from One Global Multi-Cloud Cluster",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies this advanced cluster.`)
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.db, "db", "", `Human-readable label that identifies the database that contains the collection.`)
	cmd.Flags().StringVar(&opts.collection, "collection", "", `Human-readable label that identifies the collection associated with the managed namespace.`)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("clusterName")
	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type getManagedNamespaceOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	groupId     string
	clusterName string
}

func (opts *getManagedNamespaceOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getManagedNamespaceOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetManagedNamespaceApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.GlobalClustersApi.GetManagedNamespaceWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getManagedNamespaceBuilder() *cobra.Command {
	opts := getManagedNamespaceOpts{}
	cmd := &cobra.Command{
		Use:   "getManagedNamespace",
		Short: "Return One Managed Namespace in One Global Multi-Cloud Cluster",
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies this advanced cluster.`)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

func globalClustersBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use: "globalClusters",
		Short: `Returns, adds, and removes Global Cluster managed namespaces and custom zone mappings. Each collection in a Global Cluster is associated with a managed namespace. When you create a managed namespace for a Global Cluster, MongoDB Cloud creates an empty collection for that namespace. Creating a managed namespace doesn&#39;t populate a collection with data. Similarly, deleting a managed namespace doesn&#39;t delete the associated collection.
MongoDB Cloud shards the empty collection using the required location field and a custom shard key. For example, if your custom shard key is &#x60;city&#x60;, the compound shard key is &#x60;location, city&#x60;. Each Global Cluster is also associated with one or more Global Writes Zones. When a user creates a Global Cluster, MongoDB Cloud automatically maps each location code to the closest geographical zone. Custom zone mappings allow administrators to override these automatic mappings. For example, a use case might require mapping a location code to a geographically distant zone. Administrators can manage custom zone mappings with the APIs below and the **Global Cluster Configuration** pane when you create or modify your Global Cluster.`,
	}
	cmd.AddCommand(
		createCustomZoneMappingBuilder(),
		createManagedNamespaceBuilder(),
		deleteAllCustomZoneMappingsBuilder(),
		deleteManagedNamespaceBuilder(),
		getManagedNamespaceBuilder(),
	)
	return cmd
}
