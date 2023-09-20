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

type createUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient

	filename string
	fs       afero.Fs
}

func (opts *createUserOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createUserOpts) readData() (*admin.CloudAppUser, error) {
	var out *admin.CloudAppUser

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

func (opts *createUserOpts) Run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	params := &admin.CreateUserApiParams{

		CloudAppUser: data,
	}
	resp, _, err := opts.client.MongoDBCloudUsersApi.CreateUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func createUserBuilder() *cobra.Command {
	opts := createUserOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createUser",
		Short: "Create One MongoDB Cloud User",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

type getUserOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client *admin.APIClient
	userId string
}

func (opts *getUserOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getUserOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetUserApiParams{
		UserId: opts.userId,
	}
	resp, _, err := opts.client.MongoDBCloudUsersApi.GetUserWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getUserBuilder() *cobra.Command {
	opts := getUserOpts{}
	cmd := &cobra.Command{
		Use:   "getUser",
		Short: "Return One MongoDB Cloud User using Its ID",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.userId, "userId", "", `Unique 24-hexadecimal digit string that identifies this user.`)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("userId")
	return cmd
}

type getUserByUsernameOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client   *admin.APIClient
	userName string
}

func (opts *getUserByUsernameOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getUserByUsernameOpts) Run(ctx context.Context, w io.Writer) error {
	params := &admin.GetUserByUsernameApiParams{
		UserName: opts.userName,
	}
	resp, _, err := opts.client.MongoDBCloudUsersApi.GetUserByUsernameWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return jsonwriter.Print(w, resp)
}

func getUserByUsernameBuilder() *cobra.Command {
	opts := getUserByUsernameOpts{}
	cmd := &cobra.Command{
		Use:   "getUserByUsername",
		Short: "Return One MongoDB Cloud User using Their Username",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initClient(),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.userName, "userName", "", `Email address that belongs to the MongoDB Cloud user account. You cannot modify this address after creating the user.`)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired("userName")
	return cmd
}

func mongoDBCloudUsersBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mongoDBCloudUsers",
		Short: `Returns, adds, and edits MongoDB Cloud users.`,
	}
	cmd.AddCommand(
		createUserBuilder(),
		getUserBuilder(),
		getUserByUsernameBuilder(),
	)
	return cmd
}
