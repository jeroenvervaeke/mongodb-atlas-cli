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

type CreateDatabaseUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	
}

func (opts *CreateDatabaseUserOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *CreateDatabaseUserOpts) Run(ctx context.Context) error {
	params := &admin.CreateDatabaseUserApiParams{
		GroupId: opts.groupId,
		
	}
	resp, _, err := opts.client.DatabaseUsersApi.CreateDatabaseUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func CreateDatabaseUserBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := CreateDatabaseUserOpts{}
	cmd := &cobra.Command{
		Use:     "createDatabaseUser",
		// Aliases: []string{"?"},
		Short:   "Create One Database User in One Project",
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
type DeleteDatabaseUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	databaseName string
	username string
}

func (opts *DeleteDatabaseUserOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *DeleteDatabaseUserOpts) Run(ctx context.Context) error {
	params := &admin.DeleteDatabaseUserApiParams{
		GroupId: opts.groupId,
		DatabaseName: opts.databaseName,
		Username: opts.username,
	}
	resp, _, err := opts.client.DatabaseUsersApi.DeleteDatabaseUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func DeleteDatabaseUserBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := DeleteDatabaseUserOpts{}
	cmd := &cobra.Command{
		Use:     "deleteDatabaseUser",
		// Aliases: []string{"?"},
		Short:   "Remove One Database User from One Project",
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
	cmd.Flags().StringVar(&opts.databaseName, "databaseName", "", "Human-readable label that identifies the database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, or LDAP, this value should be &#x60;$external&#x60;. If the user authenticates with SCRAM-SHA, this value should be &#x60;admin&#x60;.")
	cmd.Flags().StringVar(&opts.username, "username", "", "Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:  | Authentication Method | Parameter Needed | Parameter Value | username Format | |---|---|---|---| | AWS IAM | awsType | ROLE | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | AWS IAM | awsType | USER | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | SCRAM-SHA | awsType, x509Type, ldapAuthType | NONE | Alphanumeric string | ")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("databaseName")
	_ = cmd.MarkFlagRequired("username")
	return cmd
}
type GetDatabaseUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	databaseName string
	username string
}

func (opts *GetDatabaseUserOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *GetDatabaseUserOpts) Run(ctx context.Context) error {
	params := &admin.GetDatabaseUserApiParams{
		GroupId: opts.groupId,
		DatabaseName: opts.databaseName,
		Username: opts.username,
	}
	resp, _, err := opts.client.DatabaseUsersApi.GetDatabaseUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func GetDatabaseUserBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := GetDatabaseUserOpts{}
	cmd := &cobra.Command{
		Use:     "getDatabaseUser",
		// Aliases: []string{"?"},
		Short:   "Return One Database User from One Project",
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
	cmd.Flags().StringVar(&opts.databaseName, "databaseName", "", "Human-readable label that identifies the database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, or LDAP, this value should be &#x60;$external&#x60;. If the user authenticates with SCRAM-SHA, this value should be &#x60;admin&#x60;.")
	cmd.Flags().StringVar(&opts.username, "username", "", "Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:  | Authentication Method | Parameter Needed | Parameter Value | username Format | |---|---|---|---| | AWS IAM | awsType | ROLE | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | AWS IAM | awsType | USER | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | SCRAM-SHA | awsType, x509Type, ldapAuthType | NONE | Alphanumeric string | ")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("databaseName")
	_ = cmd.MarkFlagRequired("username")
	return cmd
}
type ListDatabaseUsersOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	includeCount bool
	itemsPerPage int
	pageNum int
}

func (opts *ListDatabaseUsersOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *ListDatabaseUsersOpts) Run(ctx context.Context) error {
	params := &admin.ListDatabaseUsersApiParams{
		GroupId: opts.groupId,
		IncludeCount: opts.includeCount,
		ItemsPerPage: opts.itemsPerPage,
		PageNum: opts.pageNum,
	}
	resp, _, err := opts.client.DatabaseUsersApi.ListDatabaseUsersWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func ListDatabaseUsersBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := ListDatabaseUsersOpts{}
	cmd := &cobra.Command{
		Use:     "listDatabaseUsers",
		// Aliases: []string{"?"},
		Short:   "Return All Database Users from One Project",
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
type UpdateDatabaseUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	groupId string
	databaseName string
	username string
	
}

func (opts *UpdateDatabaseUserOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = NewClientWithAuth()
		return err
	}
}

func (opts *UpdateDatabaseUserOpts) Run(ctx context.Context) error {
	params := &admin.UpdateDatabaseUserApiParams{
		GroupId: opts.groupId,
		DatabaseName: opts.databaseName,
		Username: opts.username,
		
	}
	resp, _, err := opts.client.DatabaseUsersApi.UpdateDatabaseUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func UpdateDatabaseUserBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := UpdateDatabaseUserOpts{}
	cmd := &cobra.Command{
		Use:     "updateDatabaseUser",
		// Aliases: []string{"?"},
		Short:   "Update One Database User in One Project",
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
	cmd.Flags().StringVar(&opts.databaseName, "databaseName", "", "Human-readable label that identifies the database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, or LDAP, this value should be &#x60;$external&#x60;. If the user authenticates with SCRAM-SHA, this value should be &#x60;admin&#x60;.")
	cmd.Flags().StringVar(&opts.username, "username", "", "Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:  | Authentication Method | Parameter Needed | Parameter Value | username Format | |---|---|---|---| | AWS IAM | awsType | ROLE | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | AWS IAM | awsType | USER | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | SCRAM-SHA | awsType, x509Type, ldapAuthType | NONE | Alphanumeric string | ")
	

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("databaseName")
	_ = cmd.MarkFlagRequired("username")
	return cmd
}

func DatabaseUsersBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "databaseUsers",
		Short:   "Returns, adds, edits, and removes database users.",
	}
	cmd.AddCommand(
		CreateDatabaseUserBuilder(),
		DeleteDatabaseUserBuilder(),
		GetDatabaseUserBuilder(),
		ListDatabaseUsersBuilder(),
		UpdateDatabaseUserBuilder(),
	)
	return cmd
}

