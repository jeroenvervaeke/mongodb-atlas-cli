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

type createThirdPartyIntegrationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client          *admin.APIClient
	integrationType string
	groupId         string

	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *createThirdPartyIntegrationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *createThirdPartyIntegrationOpts) Run(ctx context.Context) error {
	params := &admin.CreateThirdPartyIntegrationApiParams{
		IntegrationType: opts.integrationType,
		GroupId:         opts.groupId,

		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}
	resp, _, err := opts.client.ThirdPartyIntegrationsApi.CreateThirdPartyIntegrationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func createThirdPartyIntegrationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := createThirdPartyIntegrationOpts{}
	cmd := &cobra.Command{
		Use:   "createThirdPartyIntegration",
		Short: "Configure One Third-Party Service Integration",
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
	cmd.Flags().StringVar(&opts.integrationType, "integrationType", "", `Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.`)
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("integrationType")
	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type deleteThirdPartyIntegrationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client          *admin.APIClient
	integrationType string
	groupId         string
}

func (opts *deleteThirdPartyIntegrationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *deleteThirdPartyIntegrationOpts) Run(ctx context.Context) error {
	params := &admin.DeleteThirdPartyIntegrationApiParams{
		IntegrationType: opts.integrationType,
		GroupId:         opts.groupId,
	}
	resp, _, err := opts.client.ThirdPartyIntegrationsApi.DeleteThirdPartyIntegrationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func deleteThirdPartyIntegrationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := deleteThirdPartyIntegrationOpts{}
	cmd := &cobra.Command{
		Use:   "deleteThirdPartyIntegration",
		Short: "Remove One Third-Party Service Integration",
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
	cmd.Flags().StringVar(&opts.integrationType, "integrationType", "", `Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.`)
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	_ = cmd.MarkFlagRequired("integrationType")
	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type getThirdPartyIntegrationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client          *admin.APIClient
	groupId         string
	integrationType string
}

func (opts *getThirdPartyIntegrationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *getThirdPartyIntegrationOpts) Run(ctx context.Context) error {
	params := &admin.GetThirdPartyIntegrationApiParams{
		GroupId:         opts.groupId,
		IntegrationType: opts.integrationType,
	}
	resp, _, err := opts.client.ThirdPartyIntegrationsApi.GetThirdPartyIntegrationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func getThirdPartyIntegrationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := getThirdPartyIntegrationOpts{}
	cmd := &cobra.Command{
		Use:   "getThirdPartyIntegration",
		Short: "Return One Third-Party Service Integration",
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
	cmd.Flags().StringVar(&opts.integrationType, "integrationType", "", `Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("integrationType")
	return cmd
}

type listThirdPartyIntegrationsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client       *admin.APIClient
	groupId      string
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listThirdPartyIntegrationsOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *listThirdPartyIntegrationsOpts) Run(ctx context.Context) error {
	params := &admin.ListThirdPartyIntegrationsApiParams{
		GroupId:      opts.groupId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}
	resp, _, err := opts.client.ThirdPartyIntegrationsApi.ListThirdPartyIntegrationsWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func listThirdPartyIntegrationsBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := listThirdPartyIntegrationsOpts{}
	cmd := &cobra.Command{
		Use:   "listThirdPartyIntegrations",
		Short: "Return All Active Third-Party Service Integrations",
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

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type updateThirdPartyIntegrationOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	client          *admin.APIClient
	integrationType string
	groupId         string

	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *updateThirdPartyIntegrationOpts) initClient() func() error {
	return func() error {
		var err error
		opts.client, err = newClientWithAuth()
		return err
	}
}

func (opts *updateThirdPartyIntegrationOpts) Run(ctx context.Context) error {
	params := &admin.UpdateThirdPartyIntegrationApiParams{
		IntegrationType: opts.integrationType,
		GroupId:         opts.groupId,

		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}
	resp, _, err := opts.client.ThirdPartyIntegrationsApi.UpdateThirdPartyIntegrationWithParams(ctx, params).Execute()
	if err != nil {
		return err
	}

	return opts.Print(resp)
}

func updateThirdPartyIntegrationBuilder() *cobra.Command {
	const template = "<<some template>>"

	opts := updateThirdPartyIntegrationOpts{}
	cmd := &cobra.Command{
		Use:   "updateThirdPartyIntegration",
		Short: "Update One Third-Party Service Integration",
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
	cmd.Flags().StringVar(&opts.integrationType, "integrationType", "", `Human-readable label that identifies the service which you want to integrate with MongoDB Cloud.`)
	cmd.Flags().StringVar(&opts.groupId, "groupId", "", `Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.

**NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.`)

	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("integrationType")
	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

func thirdPartyIntegrationsBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use: "thirdPartyIntegrations",
		Short: `Returns, adds, edits, and removes third-party service integration configurations. MongoDB Cloud sends alerts to each third-party service that you configure.

**IMPORTANT**: Each project can only have one configuration per integrationType.`,
	}
	cmd.AddCommand(
		createThirdPartyIntegrationBuilder(),
		deleteThirdPartyIntegrationBuilder(),
		getThirdPartyIntegrationBuilder(),
		listThirdPartyIntegrationsBuilder(),
		updateThirdPartyIntegrationBuilder(),
	)
	return cmd
}
