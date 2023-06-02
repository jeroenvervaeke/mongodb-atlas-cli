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

package generated

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/admin"
)

type getClusterAdvancedConfigurationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	groupId     string
	clusterName string
}

func (opts *getClusterAdvancedConfigurationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getClusterAdvancedConfigurationOpts) Run(ctx context.Context) error {
	params := &admin.GetClusterAdvancedConfigurationApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.ClustersApi.GetClusterAdvancedConfigurationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func getClusterAdvancedConfigurationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := getClusterAdvancedConfigurationOpts{}
	cmd := &cobra.Command{
		Use:   "getClusterAdvancedConfiguration",
		Short: "Return One Advanced Configuration Options for One Cluster",
		Annotations: map[string]string{
			"output": template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type getClusterStatusOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	groupId     string
	clusterName string
}

func (opts *getClusterStatusOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getClusterStatusOpts) Run(ctx context.Context) error {
	params := &admin.GetClusterStatusApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.ClustersApi.GetClusterStatusWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func getClusterStatusBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := getClusterStatusOpts{}
	cmd := &cobra.Command{
		Use:   "getClusterStatus",
		Short: "Return Status of All Cluster Operations",
		Annotations: map[string]string{
			"output": template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type getSampleDatasetLoadStatusOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client          *admin.APIClient
	groupId         string
	sampleDatasetId string
}

func (opts *getSampleDatasetLoadStatusOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getSampleDatasetLoadStatusOpts) Run(ctx context.Context) error {
	params := &admin.GetSampleDatasetLoadStatusApiParams{
		GroupId:         opts.groupId,
		SampleDatasetId: opts.sampleDatasetId,
	}
	resp, _, err := opts.client.ClustersApi.GetSampleDatasetLoadStatusWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func getSampleDatasetLoadStatusBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := getSampleDatasetLoadStatusOpts{}
	cmd := &cobra.Command{
		Use:   "getSampleDatasetLoadStatus",
		Short: "Check Status of Cluster Sample Dataset Request",
		Annotations: map[string]string{
			"output": template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.sampleDatasetId, "sampleDatasetId", "", `Unique 24-hexadecimal digit string that identifies the loaded sample dataset.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("sampleDatasetId")
	return cmd
}

type listCloudProviderRegionsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	includeCount bool
	itemsPerPage int
	pageNum      int
	providers    []string
	tier         string
}

func (opts *listCloudProviderRegionsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listCloudProviderRegionsOpts) Run(ctx context.Context) error {
	params := &admin.ListCloudProviderRegionsApiParams{
		GroupId:      opts.groupId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
		Providers:    &opts.providers,
		Tier:         &opts.tier,
	}
	resp, _, err := opts.client.ClustersApi.ListCloudProviderRegionsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func listCloudProviderRegionsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := listCloudProviderRegionsOpts{}
	cmd := &cobra.Command{
		Use:   "listCloudProviderRegions",
		Short: "Return All Cloud Provider Regions",
		Annotations: map[string]string{
			"output": template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)
	cmd.Flags().StringSliceVar(&opts.providers, "providers", nil, `Cloud providers whose regions to retrieve. When you specify multiple providers, the response can return only tiers and regions that support multi-cloud clusters.`)
	cmd.Flags().StringVar(&opts.tier, "tier", "", `Cluster tier for which to retrieve the regions.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type listClustersForAllProjectsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listClustersForAllProjectsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listClustersForAllProjectsOpts) Run(ctx context.Context) error {
	params := &admin.ListClustersForAllProjectsApiParams{
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}
	resp, _, err := opts.client.ClustersApi.ListClustersForAllProjectsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func listClustersForAllProjectsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := listClustersForAllProjectsOpts{}
	cmd := &cobra.Command{
		Use:   "listClustersForAllProjects",
		Short: "Return All Authorized Clusters in All Projects",
		Annotations: map[string]string{
			"output": template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	return cmd
}

type loadSampleDatasetOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string
	name    string
}

func (opts *loadSampleDatasetOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *loadSampleDatasetOpts) Run(ctx context.Context) error {
	params := &admin.LoadSampleDatasetApiParams{
		GroupId: opts.groupId,
		Name:    opts.name,
	}
	resp, _, err := opts.client.ClustersApi.LoadSampleDatasetWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func loadSampleDatasetBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := loadSampleDatasetOpts{}
	cmd := &cobra.Command{
		Use:   "loadSampleDataset",
		Short: "Load Sample Dataset Request into Cluster",
		Annotations: map[string]string{
			"output": template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.name, "name", "", `Human-readable label that identifies the cluster into which you load the sample dataset.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("name")
	return cmd
}

type updateClusterAdvancedConfigurationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	groupId     string
	clusterName string
}

func (opts *updateClusterAdvancedConfigurationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updateClusterAdvancedConfigurationOpts) Run(ctx context.Context) error {
	params := &admin.UpdateClusterAdvancedConfigurationApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.ClustersApi.UpdateClusterAdvancedConfigurationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func updateClusterAdvancedConfigurationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := updateClusterAdvancedConfigurationOpts{}
	cmd := &cobra.Command{
		Use:   "updateClusterAdvancedConfiguration",
		Short: "Update Advanced Configuration Options for One Cluster",
		Annotations: map[string]string{
			"output": template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type upgradeSharedClusterOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string
}

func (opts *upgradeSharedClusterOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *upgradeSharedClusterOpts) Run(ctx context.Context) error {
	params := &admin.UpgradeSharedClusterApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.ClustersApi.UpgradeSharedClusterWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func upgradeSharedClusterBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := upgradeSharedClusterOpts{}
	cmd := &cobra.Command{
		Use:   "upgradeSharedCluster",
		Short: "Upgrade One Shared-tier Cluster",
		Annotations: map[string]string{
			"output": template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type upgradeSharedClusterToServerlessOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client  *admin.APIClient
	groupId string
}

func (opts *upgradeSharedClusterToServerlessOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *upgradeSharedClusterToServerlessOpts) Run(ctx context.Context) error {
	params := &admin.UpgradeSharedClusterToServerlessApiParams{
		GroupId: opts.groupId,
	}
	resp, _, err := opts.client.ClustersApi.UpgradeSharedClusterToServerlessWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func upgradeSharedClusterToServerlessBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := upgradeSharedClusterToServerlessOpts{}
	cmd := &cobra.Command{
		Use:   "upgradeSharedClusterToServerless",
		Short: "Upgrades One Shared-Tier Cluster to the Serverless Instance",
		Annotations: map[string]string{
			"output": template,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
				opts.InitOutput(cmd.OutOrStdout(), template),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

func clustersBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clusters",
		Short: `Returns, adds, edits, and removes database deployments. Changes to cluster configurations can affect costs. This resource requires your Project ID.`,
	}
	cmd.AddCommand(
		getClusterAdvancedConfigurationBuilder(),
		getClusterStatusBuilder(),
		getSampleDatasetLoadStatusBuilder(),
		listCloudProviderRegionsBuilder(),
		listClustersForAllProjectsBuilder(),
		loadSampleDatasetBuilder(),
		updateClusterAdvancedConfigurationBuilder(),
		upgradeSharedClusterBuilder(),
		upgradeSharedClusterToServerlessBuilder(),
	)
	return cmd
}
