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
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/admin"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
)

type CreateClusterOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	
}

func (opts *CreateClusterOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateClusterOpts) Run(ctx context.Context) error {
	params := &admin.CreateClusterApiParams{
		GroupId: opts.groupId,
		
	}
	resp, _, err := opts.client.MultiCloudClustersApi.CreateClusterWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreateClusterBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreateClusterOpts{}
	cmd := &cobra.Command{
		Use:     "createCluster",
		// Aliases: []string{"?"},
		Short:   "Create One Multi-Cloud Cluster from One Project",
		Annotations: map[string]string{
			"output":      template,
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}
type DeleteClusterOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	retainBackups bool
}

func (opts *DeleteClusterOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteClusterOpts) Run(ctx context.Context) error {
	params := &admin.DeleteClusterApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		RetainBackups: opts.retainBackups,
	}
	_, err := opts.client.MultiCloudClustersApi.DeleteClusterWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(nil)
}

func DeleteClusterBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeleteClusterOpts{}
	cmd := &cobra.Command{
		Use:     "deleteCluster",
		// Aliases: []string{"?"},
		Short:   "Remove One Multi-Cloud Cluster from One Project",
		Annotations: map[string]string{
			"output":      template,
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Human-readable label that identifies the cluster.")
	cmd.Flags().BoolVar(&opts.retainBackups, "retainBackups", false, "Flag that indicates whether to retain backup snapshots for the deleted dedicated cluster.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}
type GetClusterOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
}

func (opts *GetClusterOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetClusterOpts) Run(ctx context.Context) error {
	params := &admin.GetClusterApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.MultiCloudClustersApi.GetClusterWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetClusterBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetClusterOpts{}
	cmd := &cobra.Command{
		Use:     "getCluster",
		// Aliases: []string{"?"},
		Short:   "Return One Multi-Cloud Cluster from One Project",
		Annotations: map[string]string{
			"output":      template,
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Human-readable label that identifies this advanced cluster.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}
type ListClustersOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	includeCount bool
	itemsPerPage int
	pageNum int
}

func (opts *ListClustersOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListClustersOpts) Run(ctx context.Context) error {
	params := &admin.ListClustersApiParams{
		GroupId: opts.groupId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
	}
	resp, _, err := opts.client.MultiCloudClustersApi.ListClustersWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListClustersBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListClustersOpts{}
	cmd := &cobra.Command{
		Use:     "listClusters",
		// Aliases: []string{"?"},
		Short:   "Return All Multi-Cloud Clusters from One Project",
		Annotations: map[string]string{
			"output":      template,
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, "Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.")
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, "Number of items that the response returns per page.")
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, "Number of the page that displays the current set of the total objects that the response returns.")

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}
type TestFailoverOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
}

func (opts *TestFailoverOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *TestFailoverOpts) Run(ctx context.Context) error {
	params := &admin.TestFailoverApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
	}
	_, err := opts.client.MultiCloudClustersApi.TestFailoverWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(nil)
}

func TestFailoverBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := TestFailoverOpts{}
	cmd := &cobra.Command{
		Use:     "testFailover",
		// Aliases: []string{"?"},
		Short:   "Test Failover for One Multi-Cloud Cluster",
		Annotations: map[string]string{
			"output":      template,
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Human-readable label that identifies the cluster.")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}
type UpdateClusterOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	clusterName string
	
}

func (opts *UpdateClusterOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateClusterOpts) Run(ctx context.Context) error {
	params := &admin.UpdateClusterApiParams{
		GroupId: opts.groupId,
		ClusterName: opts.clusterName,
		
	}
	resp, _, err := opts.client.MultiCloudClustersApi.UpdateClusterWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdateClusterBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdateClusterOpts{}
	cmd := &cobra.Command{
		Use:     "updateCluster",
		// Aliases: []string{"?"},
		Short:   "Modify One Multi-Cloud Cluster from One Project",
		Annotations: map[string]string{
			"output":      template,
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
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", "Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.")
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", "Human-readable label that identifies the advanced cluster to modify.")
	

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

func MultiCloudClustersBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "multiCloudClusters",
		Short:   "Returns, adds, edits, or removes multi-cloud clusters. Changes to cluster configurations can affect costs.

The total number of nodes in clusters spanning across regions has a specific constraint on a per-project basis. MongoDB Cloud limits the total number of nodes in other regions in one project to a total of 40. This total excludes Google Cloud regions communicating with each other, shared-tier clusters, or serverless clusters. The total number of nodes between any two regions must meet this constraint. For example, if a project has nodes in clusters spread across three regions: 30 nodes in Region A, 10 nodes in Region B, and 5 nodes in Region C, you can add only 5 more nodes to Region C because if you exclude Region C, Region A + Region B &#x3D; 40. If you exclude Region B, Region A + Region C &#x3D; 35, &lt;&#x3D; 40. If you exclude Region A, Region B + Region C &#x3D; 15, &lt;&#x3D; 40. Each combination of regions with the added 5 nodes still meets the per-project constraint. Region A + B &#x3D; 40. Region A + C &#x3D; 40. Region B + C &#x3D; 20. You can&#39;t create a multi-region cluster in a project if it has one or more clusters spanning 40 or more nodes in other regions. Each project supports up to 25 database deployments.

If your MongoDB Cloud project contains a custom role that uses actions introduced in a specific MongoDB version, you must delete that role before you create clusters with an earlier MongoDB version. MongoDB Cloud clusters created after July 2020 use TLS version 1.2 by default. When you create a cluster, MongoDB Cloud creates a network container in the project for the cloud provider to which you deploy the cluster if one doesn&#39;t already exist.",
	}
	cmd.AddCommand(
		CreateClusterBuilder(),
		DeleteClusterBuilder(),
		GetClusterBuilder(),
		ListClustersBuilder(),
		TestFailoverBuilder(),
		UpdateClusterBuilder(),
	)
	return cmd
}

