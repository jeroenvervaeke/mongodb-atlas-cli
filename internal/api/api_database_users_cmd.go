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

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type createDatabaseUserOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *createDatabaseUserOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *createDatabaseUserOpts) readData() (*admin.CloudDatabaseUser, error) {
	var out *admin.CloudDatabaseUser

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

func (opts *createDatabaseUserOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.CreateDatabaseUserApiParams{
		GroupId: opts.groupId,

		CloudDatabaseUser: data,
	}

	resp, _, err := opts.client.DatabaseUsersApi.CreateDatabaseUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func createDatabaseUserBuilder() *cobra.Command {
	opts := createDatabaseUserOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createDatabaseUser",
		Short: "Create One Database User in One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type deleteDatabaseUserOpts struct {
	client       *admin.APIClient
	groupId      string
	databaseName string
	username     string
}

func (opts *deleteDatabaseUserOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *deleteDatabaseUserOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.DeleteDatabaseUserApiParams{
		GroupId:      opts.groupId,
		DatabaseName: opts.databaseName,
		Username:     opts.username,
	}

	resp, _, err := opts.client.DatabaseUsersApi.DeleteDatabaseUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func deleteDatabaseUserBuilder() *cobra.Command {
	opts := deleteDatabaseUserOpts{}
	cmd := &cobra.Command{
		Use:   "deleteDatabaseUser",
		Short: "Remove One Database User from One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.databaseName, "databaseName", "", `Human-readable label that identifies the database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, or LDAP, this value should be &#x60;$external&#x60;. If the user authenticates with SCRAM-SHA or OIDC, this value should be &#x60;admin&#x60;.`)
	cmd.Flags().StringVar(&opts.username, "username", "", `Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:

| Authentication Method | Parameter Needed | Parameter Value | username Format |
|---|---|---|---|
| AWS IAM | awsType | ROLE | &lt;abbr title&#x3D;&quot;Amazon Resource Name&quot;&gt;ARN&lt;/abbr&gt; |
| AWS IAM | awsType | USER | &lt;abbr title&#x3D;&quot;Amazon Resource Name&quot;&gt;ARN&lt;/abbr&gt; |
| x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| OIDC | oidcAuthType | IDP_GROUP | Atlas OIDC IdP ID (found in federation settings), followed by a &#39;/&#39;, followed by the IdP group name |
| SCRAM-SHA | awsType, x509Type, ldapAuthType, oidcAuthType | NONE | Alphanumeric string |
`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("databaseName")
	_ = cmd.MarkFlagRequired("username")
	return cmd
}

type getDatabaseUserOpts struct {
	client       *admin.APIClient
	groupId      string
	databaseName string
	username     string
}

func (opts *getDatabaseUserOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getDatabaseUserOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.GetDatabaseUserApiParams{
		GroupId:      opts.groupId,
		DatabaseName: opts.databaseName,
		Username:     opts.username,
	}

	resp, _, err := opts.client.DatabaseUsersApi.GetDatabaseUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func getDatabaseUserBuilder() *cobra.Command {
	opts := getDatabaseUserOpts{}
	cmd := &cobra.Command{
		Use:   "getDatabaseUser",
		Short: "Return One Database User from One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.databaseName, "databaseName", "", `Human-readable label that identifies the database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, or LDAP, this value should be &#x60;$external&#x60;. If the user authenticates with SCRAM-SHA or OIDC, this value should be &#x60;admin&#x60;.`)
	cmd.Flags().StringVar(&opts.username, "username", "", `Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:

| Authentication Method | Parameter Needed | Parameter Value | username Format |
|---|---|---|---|
| AWS IAM | awsType | ROLE | &lt;abbr title&#x3D;&quot;Amazon Resource Name&quot;&gt;ARN&lt;/abbr&gt; |
| AWS IAM | awsType | USER | &lt;abbr title&#x3D;&quot;Amazon Resource Name&quot;&gt;ARN&lt;/abbr&gt; |
| x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| OIDC | oidcAuthType | IDP_GROUP | Atlas OIDC IdP ID (found in federation settings), followed by a &#39;/&#39;, followed by the IdP group name |
| SCRAM-SHA | awsType, x509Type, ldapAuthType, oidcAuthType | NONE | Alphanumeric string |
`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("databaseName")
	_ = cmd.MarkFlagRequired("username")
	return cmd
}

type listDatabaseUsersOpts struct {
	client       *admin.APIClient
	groupId      string
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listDatabaseUsersOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listDatabaseUsersOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.ListDatabaseUsersApiParams{
		GroupId:      opts.groupId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}

	resp, _, err := opts.client.DatabaseUsersApi.ListDatabaseUsersWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func listDatabaseUsersBuilder() *cobra.Command {
	opts := listDatabaseUsersOpts{}
	cmd := &cobra.Command{
		Use:   "listDatabaseUsers",
		Short: "Return All Database Users from One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type updateDatabaseUserOpts struct {
	client       *admin.APIClient
	groupId      string
	databaseName string
	username     string

	filename string
	fs       afero.Fs
}

func (opts *updateDatabaseUserOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *updateDatabaseUserOpts) readData() (*admin.CloudDatabaseUser, error) {
	var out *admin.CloudDatabaseUser

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

func (opts *updateDatabaseUserOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.UpdateDatabaseUserApiParams{
		GroupId:      opts.groupId,
		DatabaseName: opts.databaseName,
		Username:     opts.username,

		CloudDatabaseUser: data,
	}

	resp, _, err := opts.client.DatabaseUsersApi.UpdateDatabaseUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	prettyJSON, errJson := json.MarshalIndent(resp, "", " ")
	if errJson != nil {
		return errJson
	}

	_, err = fmt.Fprintln(w, string(prettyJSON))
	return err
}

func updateDatabaseUserBuilder() *cobra.Command {
	opts := updateDatabaseUserOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updateDatabaseUser",
		Short: "Update One Database User in One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.databaseName, "databaseName", "", `Human-readable label that identifies the database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, or LDAP, this value should be &#x60;$external&#x60;. If the user authenticates with SCRAM-SHA or OIDC, this value should be &#x60;admin&#x60;.`)
	cmd.Flags().StringVar(&opts.username, "username", "", `Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:

| Authentication Method | Parameter Needed | Parameter Value | username Format |
|---|---|---|---|
| AWS IAM | awsType | ROLE | &lt;abbr title&#x3D;&quot;Amazon Resource Name&quot;&gt;ARN&lt;/abbr&gt; |
| AWS IAM | awsType | USER | &lt;abbr title&#x3D;&quot;Amazon Resource Name&quot;&gt;ARN&lt;/abbr&gt; |
| x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name |
| OIDC | oidcAuthType | IDP_GROUP | Atlas OIDC IdP ID (found in federation settings), followed by a &#39;/&#39;, followed by the IdP group name |
| SCRAM-SHA | awsType, x509Type, ldapAuthType, oidcAuthType | NONE | Alphanumeric string |
`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("databaseName")
	_ = cmd.MarkFlagRequired("username")
	return cmd
}

func databaseUsersBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "databaseUsers",
		Short: `Returns, adds, edits, and removes database users.`,
	}
	cmd.AddCommand(
		createDatabaseUserBuilder(),
		deleteDatabaseUserBuilder(),
		getDatabaseUserBuilder(),
		listDatabaseUsersBuilder(),
		updateDatabaseUserBuilder(),
	)
	return cmd
}
