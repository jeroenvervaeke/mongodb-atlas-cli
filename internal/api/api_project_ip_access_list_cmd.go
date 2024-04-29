// Copyright 2024 MongoDB Inc
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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20231115012/admin"
)

type createProjectIpAccessListOpts struct {
	client  *admin.APIClient
	groupId string

	includeCount bool
	itemsPerPage int
	pageNum      int
	filename     string
	fs           afero.Fs
	format       string
	tmpl         *template.Template
	resp         *admin.PaginatedNetworkAccess
}

func (opts *createProjectIpAccessListOpts) preRun() (err error) {
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
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *createProjectIpAccessListOpts) readData(r io.Reader) (*[]admin.NetworkPermissionEntry, error) {
	var out *[]admin.NetworkPermissionEntry

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

func (opts *createProjectIpAccessListOpts) run(ctx context.Context, r io.Reader) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.CreateProjectIpAccessListApiParams{
		GroupId: opts.groupId,

		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,

		NetworkPermissionEntry: data,
	}

	var err error
	opts.resp, _, err = opts.client.ProjectIPAccessListApi.CreateProjectIpAccessListWithParams(ctx, params).Execute()
	return err
}

func (opts *createProjectIpAccessListOpts) postRun(_ context.Context, w io.Writer) error {

	prettyJSON, errJson := json.MarshalIndent(opts.resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err := fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err := json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	return opts.tmpl.Execute(w, parsedJSON)
}

func createProjectIpAccessListBuilder() *cobra.Command {
	opts := createProjectIpAccessListOpts{
		fs: afero.NewOsFs(),
	}
	const use = "createProjectIpAccessList"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Add Entries to Project IP Access List",
		Aliases: cli.GenerateAliases(use),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.postRun(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)
	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type deleteProjectIpAccessListOpts struct {
	client     *admin.APIClient
	groupId    string
	entryValue string
	format     string
	tmpl       *template.Template
	resp       map[string]interface{}
}

func (opts *deleteProjectIpAccessListOpts) preRun() (err error) {
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
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *deleteProjectIpAccessListOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.DeleteProjectIpAccessListApiParams{
		GroupId:    opts.groupId,
		EntryValue: opts.entryValue,
	}

	var err error
	opts.resp, _, err = opts.client.ProjectIPAccessListApi.DeleteProjectIpAccessListWithParams(ctx, params).Execute()
	return err
}

func (opts *deleteProjectIpAccessListOpts) postRun(_ context.Context, w io.Writer) error {

	prettyJSON, errJson := json.MarshalIndent(opts.resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err := fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err := json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	return opts.tmpl.Execute(w, parsedJSON)
}

func deleteProjectIpAccessListBuilder() *cobra.Command {
	opts := deleteProjectIpAccessListOpts{}
	const use = "deleteProjectIpAccessList"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Remove One Entry from One Project IP Access List",
		Aliases: cli.GenerateAliases(use),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.postRun(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.entryValue, "entryValue", "", `Access list entry that you want to remove from the project&#39;s IP access list. This value can use one of the following: one AWS security group ID, one IP address, or one CIDR block of addresses. For CIDR blocks that use a subnet mask, replace the forward slash (&#x60;/&#x60;) with its URL-encoded value (&#x60;%2F&#x60;). When you remove an entry from the IP access list, existing connections from the removed address or addresses may remain open for a variable amount of time. The amount of time it takes MongoDB Cloud to close the connection depends upon several factors, including:

- how your application established the connection,
- how MongoDB Cloud or the driver using the address behaves, and
- which protocol (like TCP or UDP) the connection uses.`)

	_ = cmd.MarkFlagRequired("entryValue")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type getProjectIpAccessListStatusOpts struct {
	client     *admin.APIClient
	groupId    string
	entryValue string
	format     string
	tmpl       *template.Template
	resp       *admin.NetworkPermissionEntryStatus
}

func (opts *getProjectIpAccessListStatusOpts) preRun() (err error) {
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
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *getProjectIpAccessListStatusOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.GetProjectIpAccessListStatusApiParams{
		GroupId:    opts.groupId,
		EntryValue: opts.entryValue,
	}

	var err error
	opts.resp, _, err = opts.client.ProjectIPAccessListApi.GetProjectIpAccessListStatusWithParams(ctx, params).Execute()
	return err
}

func (opts *getProjectIpAccessListStatusOpts) postRun(_ context.Context, w io.Writer) error {

	prettyJSON, errJson := json.MarshalIndent(opts.resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err := fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err := json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	return opts.tmpl.Execute(w, parsedJSON)
}

func getProjectIpAccessListStatusBuilder() *cobra.Command {
	opts := getProjectIpAccessListStatusOpts{}
	const use = "getProjectIpAccessListStatus"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Return Status of One Project IP Access List Entry",
		Aliases: cli.GenerateAliases(use),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.postRun(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.entryValue, "entryValue", "", `Network address or cloud provider security construct that identifies which project access list entry to be verified.`)

	_ = cmd.MarkFlagRequired("entryValue")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type getProjectIpListOpts struct {
	client     *admin.APIClient
	groupId    string
	entryValue string
	format     string
	tmpl       *template.Template
	resp       *admin.NetworkPermissionEntry
}

func (opts *getProjectIpListOpts) preRun() (err error) {
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
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *getProjectIpListOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.GetProjectIpListApiParams{
		GroupId:    opts.groupId,
		EntryValue: opts.entryValue,
	}

	var err error
	opts.resp, _, err = opts.client.ProjectIPAccessListApi.GetProjectIpListWithParams(ctx, params).Execute()
	return err
}

func (opts *getProjectIpListOpts) postRun(_ context.Context, w io.Writer) error {

	prettyJSON, errJson := json.MarshalIndent(opts.resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err := fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err := json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	return opts.tmpl.Execute(w, parsedJSON)
}

func getProjectIpListBuilder() *cobra.Command {
	opts := getProjectIpListOpts{}
	const use = "getProjectIpList"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Return One Project IP Access List Entry",
		Aliases: cli.GenerateAliases(use),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.postRun(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.entryValue, "entryValue", "", `Access list entry that you want to return from the project&#39;s IP access list. This value can use one of the following: one AWS security group ID, one IP address, or one CIDR block of addresses. For CIDR blocks that use a subnet mask, replace the forward slash (&#x60;/&#x60;) with its URL-encoded value (&#x60;%2F&#x60;).`)

	_ = cmd.MarkFlagRequired("entryValue")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type listProjectIpAccessListsOpts struct {
	client       *admin.APIClient
	groupId      string
	includeCount bool
	itemsPerPage int
	pageNum      int
	format       string
	tmpl         *template.Template
	resp         *admin.PaginatedNetworkAccess
}

func (opts *listProjectIpAccessListsOpts) preRun() (err error) {
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
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *listProjectIpAccessListsOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.ListProjectIpAccessListsApiParams{
		GroupId:      opts.groupId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}

	var err error
	opts.resp, _, err = opts.client.ProjectIPAccessListApi.ListProjectIpAccessListsWithParams(ctx, params).Execute()
	return err
}

func (opts *listProjectIpAccessListsOpts) postRun(_ context.Context, w io.Writer) error {

	prettyJSON, errJson := json.MarshalIndent(opts.resp, "", " ")
	if errJson != nil {
		return errJson
	}

	if opts.format == "" {
		_, err := fmt.Fprintln(w, string(prettyJSON))
		return err
	}

	var parsedJSON interface{}
	if err := json.Unmarshal([]byte(prettyJSON), &parsedJSON); err != nil {
		return err
	}

	return opts.tmpl.Execute(w, parsedJSON)
}

func listProjectIpAccessListsBuilder() *cobra.Command {
	opts := listProjectIpAccessListsOpts{}
	const use = "listProjectIpAccessLists"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Return Project IP Access List",
		Aliases: cli.GenerateAliases(use),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.InOrStdin())
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return opts.postRun(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

func projectIPAccessListBuilder() *cobra.Command {
	const use = "projectIPAccessList"
	cmd := &cobra.Command{
		Use:     use,
		Short:   `Returns, adds, edits, and removes network access limits to database deployments in Atlas. This resource replaces the whitelist resource. Atlas removed whitelists in July 2021. Update your applications to use this new resource. This resource manages a project&#39;s IP Access List and supports creating temporary Access List entries that automatically expire within a user-configurable 7-day period.`,
		Aliases: cli.GenerateAliases(use),
	}
	cmd.AddCommand(
		createProjectIpAccessListBuilder(),
		deleteProjectIpAccessListBuilder(),
		getProjectIpAccessListStatusBuilder(),
		getProjectIpListBuilder(),
		listProjectIpAccessListsBuilder(),
	)
	return cmd
}
