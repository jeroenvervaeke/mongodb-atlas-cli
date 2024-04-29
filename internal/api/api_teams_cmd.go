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

type addAllTeamsToProjectOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
	resp     *admin.PaginatedTeamRole
}

func (opts *addAllTeamsToProjectOpts) preRun() (err error) {
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

func (opts *addAllTeamsToProjectOpts) readData(r io.Reader) (*[]admin.TeamRole, error) {
	var out *[]admin.TeamRole

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

func (opts *addAllTeamsToProjectOpts) run(ctx context.Context, r io.Reader) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.AddAllTeamsToProjectApiParams{
		GroupId: opts.groupId,

		TeamRole: data,
	}

	var err error
	opts.resp, _, err = opts.client.TeamsApi.AddAllTeamsToProjectWithParams(ctx, params).Execute()
	return err
}

func (opts *addAllTeamsToProjectOpts) postRun(_ context.Context, w io.Writer) error {

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

func addAllTeamsToProjectBuilder() *cobra.Command {
	opts := addAllTeamsToProjectOpts{
		fs: afero.NewOsFs(),
	}
	const use = "addAllTeamsToProject"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Add One or More Teams to One Project",
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

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type addTeamUserOpts struct {
	client *admin.APIClient
	orgId  string
	teamId string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
	resp     *admin.PaginatedApiAppUser
}

func (opts *addTeamUserOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	if opts.format != "" {
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *addTeamUserOpts) readData(r io.Reader) (*[]admin.AddUserToTeam, error) {
	var out *[]admin.AddUserToTeam

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

func (opts *addTeamUserOpts) run(ctx context.Context, r io.Reader) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.AddTeamUserApiParams{
		OrgId:  opts.orgId,
		TeamId: opts.teamId,

		AddUserToTeam: data,
	}

	var err error
	opts.resp, _, err = opts.client.TeamsApi.AddTeamUserWithParams(ctx, params).Execute()
	return err
}

func (opts *addTeamUserOpts) postRun(_ context.Context, w io.Writer) error {

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

func addTeamUserBuilder() *cobra.Command {
	opts := addTeamUserOpts{
		fs: afero.NewOsFs(),
	}
	const use = "addTeamUser"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Assign MongoDB Cloud Users from One Organization to One Team",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.teamId, "teamId", "", `Unique 24-hexadecimal character string that identifies the team to which you want to add MongoDB Cloud users.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("teamId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type createTeamOpts struct {
	client *admin.APIClient
	orgId  string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
	resp     *admin.Team
}

func (opts *createTeamOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	if opts.format != "" {
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *createTeamOpts) readData(r io.Reader) (*admin.Team, error) {
	var out *admin.Team

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

func (opts *createTeamOpts) run(ctx context.Context, r io.Reader) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.CreateTeamApiParams{
		OrgId: opts.orgId,

		Team: data,
	}

	var err error
	opts.resp, _, err = opts.client.TeamsApi.CreateTeamWithParams(ctx, params).Execute()
	return err
}

func (opts *createTeamOpts) postRun(_ context.Context, w io.Writer) error {

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

func createTeamBuilder() *cobra.Command {
	opts := createTeamOpts{
		fs: afero.NewOsFs(),
	}
	const use = "createTeam"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Create One Team in One Organization",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type deleteTeamOpts struct {
	client *admin.APIClient
	orgId  string
	teamId string
	format string
	tmpl   *template.Template
	resp   map[string]interface{}
}

func (opts *deleteTeamOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	if opts.format != "" {
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *deleteTeamOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.DeleteTeamApiParams{
		OrgId:  opts.orgId,
		TeamId: opts.teamId,
	}

	var err error
	opts.resp, _, err = opts.client.TeamsApi.DeleteTeamWithParams(ctx, params).Execute()
	return err
}

func (opts *deleteTeamOpts) postRun(_ context.Context, w io.Writer) error {

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

func deleteTeamBuilder() *cobra.Command {
	opts := deleteTeamOpts{}
	const use = "deleteTeam"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Remove One Team from One Organization",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.teamId, "teamId", "", `Unique 24-hexadecimal digit string that identifies the team that you want to delete.`)

	_ = cmd.MarkFlagRequired("teamId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type getTeamByIdOpts struct {
	client *admin.APIClient
	orgId  string
	teamId string
	format string
	tmpl   *template.Template
	resp   *admin.TeamResponse
}

func (opts *getTeamByIdOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	if opts.format != "" {
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *getTeamByIdOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.GetTeamByIdApiParams{
		OrgId:  opts.orgId,
		TeamId: opts.teamId,
	}

	var err error
	opts.resp, _, err = opts.client.TeamsApi.GetTeamByIdWithParams(ctx, params).Execute()
	return err
}

func (opts *getTeamByIdOpts) postRun(_ context.Context, w io.Writer) error {

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

func getTeamByIdBuilder() *cobra.Command {
	opts := getTeamByIdOpts{}
	const use = "getTeamById"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Return One Team using its ID",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.teamId, "teamId", "", `Unique 24-hexadecimal digit string that identifies the team whose information you want to return.`)

	_ = cmd.MarkFlagRequired("teamId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type getTeamByNameOpts struct {
	client   *admin.APIClient
	orgId    string
	teamName string
	format   string
	tmpl     *template.Template
	resp     *admin.TeamResponse
}

func (opts *getTeamByNameOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	if opts.format != "" {
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *getTeamByNameOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.GetTeamByNameApiParams{
		OrgId:    opts.orgId,
		TeamName: opts.teamName,
	}

	var err error
	opts.resp, _, err = opts.client.TeamsApi.GetTeamByNameWithParams(ctx, params).Execute()
	return err
}

func (opts *getTeamByNameOpts) postRun(_ context.Context, w io.Writer) error {

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

func getTeamByNameBuilder() *cobra.Command {
	opts := getTeamByNameOpts{}
	const use = "getTeamByName"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Return One Team using its Name",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.teamName, "teamName", "", `Name of the team whose information you want to return.`)

	_ = cmd.MarkFlagRequired("teamName")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type listOrganizationTeamsOpts struct {
	client       *admin.APIClient
	orgId        string
	itemsPerPage int
	includeCount bool
	pageNum      int
	format       string
	tmpl         *template.Template
	resp         *admin.PaginatedTeam
}

func (opts *listOrganizationTeamsOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	if opts.format != "" {
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *listOrganizationTeamsOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.ListOrganizationTeamsApiParams{
		OrgId:        opts.orgId,
		ItemsPerPage: &opts.itemsPerPage,
		IncludeCount: &opts.includeCount,
		PageNum:      &opts.pageNum,
	}

	var err error
	opts.resp, _, err = opts.client.TeamsApi.ListOrganizationTeamsWithParams(ctx, params).Execute()
	return err
}

func (opts *listOrganizationTeamsOpts) postRun(_ context.Context, w io.Writer) error {

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

func listOrganizationTeamsBuilder() *cobra.Command {
	opts := listOrganizationTeamsOpts{}
	const use = "listOrganizationTeams"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Return All Teams in One Organization",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type listProjectTeamsOpts struct {
	client       *admin.APIClient
	groupId      string
	includeCount bool
	itemsPerPage int
	pageNum      int
	format       string
	tmpl         *template.Template
	resp         *admin.PaginatedTeamRole
}

func (opts *listProjectTeamsOpts) preRun() (err error) {
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

func (opts *listProjectTeamsOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.ListProjectTeamsApiParams{
		GroupId:      opts.groupId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}

	var err error
	opts.resp, _, err = opts.client.TeamsApi.ListProjectTeamsWithParams(ctx, params).Execute()
	return err
}

func (opts *listProjectTeamsOpts) postRun(_ context.Context, w io.Writer) error {

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

func listProjectTeamsBuilder() *cobra.Command {
	opts := listProjectTeamsOpts{}
	const use = "listProjectTeams"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Return All Teams in One Project",
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

type listTeamUsersOpts struct {
	client       *admin.APIClient
	orgId        string
	teamId       string
	itemsPerPage int
	pageNum      int
	format       string
	tmpl         *template.Template
	resp         *admin.PaginatedApiAppUser
}

func (opts *listTeamUsersOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	if opts.format != "" {
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *listTeamUsersOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.ListTeamUsersApiParams{
		OrgId:        opts.orgId,
		TeamId:       opts.teamId,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}

	var err error
	opts.resp, _, err = opts.client.TeamsApi.ListTeamUsersWithParams(ctx, params).Execute()
	return err
}

func (opts *listTeamUsersOpts) postRun(_ context.Context, w io.Writer) error {

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

func listTeamUsersBuilder() *cobra.Command {
	opts := listTeamUsersOpts{}
	const use = "listTeamUsers"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Return All MongoDB Cloud Users Assigned to One Team",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.teamId, "teamId", "", `Unique 24-hexadecimal digit string that identifies the team whose application users you want to return.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("teamId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type removeProjectTeamOpts struct {
	client  *admin.APIClient
	groupId string
	teamId  string
}

func (opts *removeProjectTeamOpts) preRun() (err error) {
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

	return nil
}

func (opts *removeProjectTeamOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.RemoveProjectTeamApiParams{
		GroupId: opts.groupId,
		TeamId:  opts.teamId,
	}

	var err error
	_, err = opts.client.TeamsApi.RemoveProjectTeamWithParams(ctx, params).Execute()
	return err
}

func (opts *removeProjectTeamOpts) postRun(_ context.Context, _ io.Writer) error {

	return nil
}

func removeProjectTeamBuilder() *cobra.Command {
	opts := removeProjectTeamOpts{}
	const use = "removeProjectTeam"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Remove One Team from One Project",
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
	cmd.Flags().StringVar(&opts.teamId, "teamId", "", `Unique 24-hexadecimal digit string that identifies the team that you want to remove from the specified project.`)

	_ = cmd.MarkFlagRequired("teamId")
	return cmd
}

type removeTeamUserOpts struct {
	client *admin.APIClient
	orgId  string
	teamId string
	userId string
}

func (opts *removeTeamUserOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	return nil
}

func (opts *removeTeamUserOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.RemoveTeamUserApiParams{
		OrgId:  opts.orgId,
		TeamId: opts.teamId,
		UserId: opts.userId,
	}

	var err error
	_, err = opts.client.TeamsApi.RemoveTeamUserWithParams(ctx, params).Execute()
	return err
}

func (opts *removeTeamUserOpts) postRun(_ context.Context, _ io.Writer) error {

	return nil
}

func removeTeamUserBuilder() *cobra.Command {
	opts := removeTeamUserOpts{}
	const use = "removeTeamUser"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Remove One MongoDB Cloud User from One Team",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.teamId, "teamId", "", `Unique 24-hexadecimal digit string that identifies the team from which you want to remove one database application user.`)
	cmd.Flags().StringVar(&opts.userId, "userId", "", `Unique 24-hexadecimal digit string that identifies MongoDB Cloud user that you want to remove from the specified team.`)

	_ = cmd.MarkFlagRequired("teamId")
	_ = cmd.MarkFlagRequired("userId")
	return cmd
}

type renameTeamOpts struct {
	client *admin.APIClient
	orgId  string
	teamId string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
	resp     *admin.TeamResponse
}

func (opts *renameTeamOpts) preRun() (err error) {
	if opts.client, err = newClientWithAuth(config.UserAgent, config.Default()); err != nil {
		return err
	}

	if opts.orgId == "" {
		opts.orgId = config.OrgID()
	}
	if opts.orgId == "" {
		return errors.New(`required flag(s) "orgId" not set`)
	}
	b, errDecode := hex.DecodeString(opts.orgId)
	if errDecode != nil || len(b) != 12 {
		return fmt.Errorf("the provided value '%s' is not a valid ID", opts.orgId)
	}

	if opts.format != "" {
		if opts.tmpl, err = template.New("").Parse(strings.ReplaceAll(opts.format, "\\n", "\n") + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (opts *renameTeamOpts) readData(r io.Reader) (*admin.TeamUpdate, error) {
	var out *admin.TeamUpdate

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

func (opts *renameTeamOpts) run(ctx context.Context, r io.Reader) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.RenameTeamApiParams{
		OrgId:  opts.orgId,
		TeamId: opts.teamId,

		TeamUpdate: data,
	}

	var err error
	opts.resp, _, err = opts.client.TeamsApi.RenameTeamWithParams(ctx, params).Execute()
	return err
}

func (opts *renameTeamOpts) postRun(_ context.Context, w io.Writer) error {

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

func renameTeamBuilder() *cobra.Command {
	opts := renameTeamOpts{
		fs: afero.NewOsFs(),
	}
	const use = "renameTeam"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Rename One Team",
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
	cmd.Flags().StringVar(&opts.orgId, "orgId", "", `Unique 24-hexadecimal digit string that identifies the organization`)
	cmd.Flags().StringVar(&opts.teamId, "teamId", "", `Unique 24-hexadecimal digit string that identifies the team that you want to rename.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("teamId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type updateTeamRolesOpts struct {
	client  *admin.APIClient
	groupId string
	teamId  string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
	resp     *admin.PaginatedTeamRole
}

func (opts *updateTeamRolesOpts) preRun() (err error) {
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

func (opts *updateTeamRolesOpts) readData(r io.Reader) (*admin.TeamRole, error) {
	var out *admin.TeamRole

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

func (opts *updateTeamRolesOpts) run(ctx context.Context, r io.Reader) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.UpdateTeamRolesApiParams{
		GroupId: opts.groupId,
		TeamId:  opts.teamId,

		TeamRole: data,
	}

	var err error
	opts.resp, _, err = opts.client.TeamsApi.UpdateTeamRolesWithParams(ctx, params).Execute()
	return err
}

func (opts *updateTeamRolesOpts) postRun(_ context.Context, w io.Writer) error {

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

func updateTeamRolesBuilder() *cobra.Command {
	opts := updateTeamRolesOpts{
		fs: afero.NewOsFs(),
	}
	const use = "updateTeamRoles"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Update Team Roles in One Project",
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
	cmd.Flags().StringVar(&opts.teamId, "teamId", "", `Unique 24-hexadecimal digit string that identifies the team for which you want to update roles.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("teamId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

func teamsBuilder() *cobra.Command {
	const use = "teams"
	cmd := &cobra.Command{
		Use:     use,
		Short:   `Returns, adds, edits, or removes teams.`,
		Aliases: cli.GenerateAliases(use),
	}
	cmd.AddCommand(
		addAllTeamsToProjectBuilder(),
		addTeamUserBuilder(),
		createTeamBuilder(),
		deleteTeamBuilder(),
		getTeamByIdBuilder(),
		getTeamByNameBuilder(),
		listOrganizationTeamsBuilder(),
		listProjectTeamsBuilder(),
		listTeamUsersBuilder(),
		removeProjectTeamBuilder(),
		removeTeamUserBuilder(),
		renameTeamBuilder(),
		updateTeamRolesBuilder(),
	)
	return cmd
}
