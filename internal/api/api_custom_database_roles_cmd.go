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
	"strings"
	"text/template"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type createCustomDatabaseRoleOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
	resp     *admin.UserCustomDBRole
}

func (opts *createCustomDatabaseRoleOpts) preRun() (err error) {
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

func (opts *createCustomDatabaseRoleOpts) readData(r io.Reader) (*admin.UserCustomDBRole, error) {
	var out *admin.UserCustomDBRole

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

func (opts *createCustomDatabaseRoleOpts) run(ctx context.Context, r io.Reader) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.CreateCustomDatabaseRoleApiParams{
		GroupId: opts.groupId,

		UserCustomDBRole: data,
	}

	var err error
	opts.resp, _, err = opts.client.CustomDatabaseRolesApi.CreateCustomDatabaseRoleWithParams(ctx, params).Execute()
	return err
}

func (opts *createCustomDatabaseRoleOpts) postRun(_ context.Context, w io.Writer) error {

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

func createCustomDatabaseRoleBuilder() *cobra.Command {
	opts := createCustomDatabaseRoleOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createCustomDatabaseRole",
		Short: "Create One Custom Role",
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

type deleteCustomDatabaseRoleOpts struct {
	client   *admin.APIClient
	groupId  string
	roleName string
}

func (opts *deleteCustomDatabaseRoleOpts) preRun() (err error) {
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

func (opts *deleteCustomDatabaseRoleOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.DeleteCustomDatabaseRoleApiParams{
		GroupId:  opts.groupId,
		RoleName: opts.roleName,
	}

	var err error
	_, err = opts.client.CustomDatabaseRolesApi.DeleteCustomDatabaseRoleWithParams(ctx, params).Execute()
	return err
}

func (opts *deleteCustomDatabaseRoleOpts) postRun(_ context.Context, _ io.Writer) error {

	return nil
}

func deleteCustomDatabaseRoleBuilder() *cobra.Command {
	opts := deleteCustomDatabaseRoleOpts{}
	cmd := &cobra.Command{
		Use:   "deleteCustomDatabaseRole",
		Short: "Remove One Custom Role from One Project",
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
	cmd.Flags().StringVar(&opts.roleName, "roleName", "", `Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.`)

	_ = cmd.MarkFlagRequired("roleName")
	return cmd
}

type getCustomDatabaseRoleOpts struct {
	client   *admin.APIClient
	groupId  string
	roleName string
	format   string
	tmpl     *template.Template
	resp     *admin.UserCustomDBRole
}

func (opts *getCustomDatabaseRoleOpts) preRun() (err error) {
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

func (opts *getCustomDatabaseRoleOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.GetCustomDatabaseRoleApiParams{
		GroupId:  opts.groupId,
		RoleName: opts.roleName,
	}

	var err error
	opts.resp, _, err = opts.client.CustomDatabaseRolesApi.GetCustomDatabaseRoleWithParams(ctx, params).Execute()
	return err
}

func (opts *getCustomDatabaseRoleOpts) postRun(_ context.Context, w io.Writer) error {

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

func getCustomDatabaseRoleBuilder() *cobra.Command {
	opts := getCustomDatabaseRoleOpts{}
	cmd := &cobra.Command{
		Use:   "getCustomDatabaseRole",
		Short: "Return One Custom Role in One Project",
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
	cmd.Flags().StringVar(&opts.roleName, "roleName", "", `Human-readable label that identifies the role for the request. This name must be unique for this custom role in this project.`)

	_ = cmd.MarkFlagRequired("roleName")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type listCustomDatabaseRolesOpts struct {
	client  *admin.APIClient
	groupId string
	format  string
	tmpl    *template.Template
	resp    []admin.UserCustomDBRole
}

func (opts *listCustomDatabaseRolesOpts) preRun() (err error) {
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

func (opts *listCustomDatabaseRolesOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.ListCustomDatabaseRolesApiParams{
		GroupId: opts.groupId,
	}

	var err error
	opts.resp, _, err = opts.client.CustomDatabaseRolesApi.ListCustomDatabaseRolesWithParams(ctx, params).Execute()
	return err
}

func (opts *listCustomDatabaseRolesOpts) postRun(_ context.Context, w io.Writer) error {

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

func listCustomDatabaseRolesBuilder() *cobra.Command {
	opts := listCustomDatabaseRolesOpts{}
	cmd := &cobra.Command{
		Use:   "listCustomDatabaseRoles",
		Short: "Return All Custom Roles in One Project",
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

	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type updateCustomDatabaseRoleOpts struct {
	client   *admin.APIClient
	groupId  string
	roleName string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
	resp     *admin.UserCustomDBRole
}

func (opts *updateCustomDatabaseRoleOpts) preRun() (err error) {
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

func (opts *updateCustomDatabaseRoleOpts) readData(r io.Reader) (*admin.UpdateCustomDBRole, error) {
	var out *admin.UpdateCustomDBRole

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

func (opts *updateCustomDatabaseRoleOpts) run(ctx context.Context, r io.Reader) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.UpdateCustomDatabaseRoleApiParams{
		GroupId:  opts.groupId,
		RoleName: opts.roleName,

		UpdateCustomDBRole: data,
	}

	var err error
	opts.resp, _, err = opts.client.CustomDatabaseRolesApi.UpdateCustomDatabaseRoleWithParams(ctx, params).Execute()
	return err
}

func (opts *updateCustomDatabaseRoleOpts) postRun(_ context.Context, w io.Writer) error {

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

func updateCustomDatabaseRoleBuilder() *cobra.Command {
	opts := updateCustomDatabaseRoleOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateCustomDatabaseRole",
		Short: "Update One Custom Role in One Project",
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
	cmd.Flags().StringVar(&opts.roleName, "roleName", "", `Human-readable label that identifies the role for the request. This name must beunique for this custom role in this project.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("roleName")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

func customDatabaseRolesBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "customDatabaseRoles",
		Short: `Returns, adds, edits, and removes custom database user privilege roles. Use custom roles to specify custom sets of actions that the MongoDB Cloud built-in roles can&#39;t describe. You define custom roles at the project level, for all clusters in the project. This resource supports a subset of MongoDB privilege actions. You can create a subset of custom role actions. To create a wider list of custom role actions, use the MongoDB Cloud user interface. Custom roles must include actions that all project&#39;s clusters support, and that are compatible with each MongoDB version that your project&#39;s clusters use. For example, if your project has MongoDB 4.2 clusters, you can&#39;t create custom roles that use actions introduced in MongoDB 4.4.`,
	}
	cmd.AddCommand(
		createCustomDatabaseRoleBuilder(),
		deleteCustomDatabaseRoleBuilder(),
		getCustomDatabaseRoleBuilder(),
		listCustomDatabaseRolesBuilder(),
		updateCustomDatabaseRoleBuilder(),
	)
	return cmd
}
