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

type createOnlineArchiveOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	groupId     string
	clusterName string
}

func (opts *createOnlineArchiveOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createOnlineArchiveOpts) Run(ctx context.Context) error {
	params := &admin.CreateOnlineArchiveApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.OnlineArchiveApi.CreateOnlineArchiveWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func createOnlineArchiveBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := createOnlineArchiveOpts{}
	cmd := &cobra.Command{
		Use:   "createOnlineArchive",
		Short: "Create One Online Archive",
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster that contains the collection for which you want to create one online archive.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type deleteOnlineArchiveOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	groupId     string
	archiveId   string
	clusterName string
}

func (opts *deleteOnlineArchiveOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deleteOnlineArchiveOpts) Run(ctx context.Context) error {
	params := &admin.DeleteOnlineArchiveApiParams{
		GroupId:     opts.groupId,
		ArchiveId:   opts.archiveId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.OnlineArchiveApi.DeleteOnlineArchiveWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func deleteOnlineArchiveBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := deleteOnlineArchiveOpts{}
	cmd := &cobra.Command{
		Use:   "deleteOnlineArchive",
		Short: "Remove One Online Archive",
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
	cmd.Flags().StringVar(&opts.archiveId, "archiveId", "", `Unique 24-hexadecimal digit string that identifies the online archive to delete.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster that contains the collection from which you want to remove an online archive.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("archiveId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type downloadOnlineArchiveQueryLogsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	groupId     string
	clusterName string
	startDate   int64
	endDate     int64
	archiveOnly bool
}

func (opts *downloadOnlineArchiveQueryLogsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *downloadOnlineArchiveQueryLogsOpts) Run(ctx context.Context) error {
	params := &admin.DownloadOnlineArchiveQueryLogsApiParams{
		GroupId:     opts.groupId,
		ClusterName: opts.clusterName,
		StartDate:   &opts.startDate,
		EndDate:     &opts.endDate,
		ArchiveOnly: &opts.archiveOnly,
	}
	resp, _, err := opts.client.OnlineArchiveApi.DownloadOnlineArchiveQueryLogsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func downloadOnlineArchiveQueryLogsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := downloadOnlineArchiveQueryLogsOpts{}
	cmd := &cobra.Command{
		Use:   "downloadOnlineArchiveQueryLogs",
		Short: "Download Online Archive Query Logs",
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster that contains the collection for which you want to return the query logs from one online archive.`)
	cmd.Flags().Int64Var(&opts.startDate, "startDate", 0, `Date and time that specifies the starting point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).`)
	cmd.Flags().Int64Var(&opts.endDate, "endDate", 0, `Date and time that specifies the end point for the range of log messages to return. This resource expresses this value in the number of seconds that have elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time).`)
	cmd.Flags().BoolVar(&opts.archiveOnly, "archiveOnly", false, `Flag that indicates whether to download logs for queries against your online archive only or both your online archive and cluster.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type getOnlineArchiveOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	groupId     string
	archiveId   string
	clusterName string
}

func (opts *getOnlineArchiveOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getOnlineArchiveOpts) Run(ctx context.Context) error {
	params := &admin.GetOnlineArchiveApiParams{
		GroupId:     opts.groupId,
		ArchiveId:   opts.archiveId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.OnlineArchiveApi.GetOnlineArchiveWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func getOnlineArchiveBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := getOnlineArchiveOpts{}
	cmd := &cobra.Command{
		Use:   "getOnlineArchive",
		Short: "Return One Online Archive",
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
	cmd.Flags().StringVar(&opts.archiveId, "archiveId", "", `Unique 24-hexadecimal digit string that identifies the online archive to return.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("archiveId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type listOnlineArchivesOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	clusterName  string
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listOnlineArchivesOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listOnlineArchivesOpts) Run(ctx context.Context) error {
	params := &admin.ListOnlineArchivesApiParams{
		GroupId:      opts.groupId,
		ClusterName:  opts.clusterName,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}
	resp, _, err := opts.client.OnlineArchiveApi.ListOnlineArchivesWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func listOnlineArchivesBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := listOnlineArchivesOpts{}
	cmd := &cobra.Command{
		Use:   "listOnlineArchives",
		Short: "Return All Online Archives for One Cluster",
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
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster that contains the collection for which you want to return the online archives.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

type updateOnlineArchiveOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client      *admin.APIClient
	groupId     string
	archiveId   string
	clusterName string
}

func (opts *updateOnlineArchiveOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updateOnlineArchiveOpts) Run(ctx context.Context) error {
	params := &admin.UpdateOnlineArchiveApiParams{
		GroupId:     opts.groupId,
		ArchiveId:   opts.archiveId,
		ClusterName: opts.clusterName,
	}
	resp, _, err := opts.client.OnlineArchiveApi.UpdateOnlineArchiveWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func updateOnlineArchiveBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := updateOnlineArchiveOpts{}
	cmd := &cobra.Command{
		Use:   "updateOnlineArchive",
		Short: "Update One Online Archive",
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
	cmd.Flags().StringVar(&opts.archiveId, "archiveId", "", `Unique 24-hexadecimal digit string that identifies the online archive to update.`)
	cmd.Flags().StringVar(&opts.clusterName, "clusterName", "", `Human-readable label that identifies the cluster that contains the specified collection from which Application created the online archive.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("archiveId")
	_ = cmd.MarkFlagRequired("clusterName")
	return cmd
}

func onlineArchiveBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "onlineArchive",
		Short: `Returns, adds, edits, or removes an online archive.`,
	}
	cmd.AddCommand(
		createOnlineArchiveBuilder(),
		deleteOnlineArchiveBuilder(),
		downloadOnlineArchiveQueryLogsBuilder(),
		getOnlineArchiveBuilder(),
		listOnlineArchivesBuilder(),
		updateOnlineArchiveBuilder(),
	)
	return cmd
}
