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

type createPeeringConnectionOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *createPeeringConnectionOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *createPeeringConnectionOpts) readData() (*admin.BaseNetworkPeeringConnectionSettings, error) {
	var out *admin.BaseNetworkPeeringConnectionSettings

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

func (opts *createPeeringConnectionOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.CreatePeeringConnectionApiParams{
		GroupId: opts.groupId,

		BaseNetworkPeeringConnectionSettings: data,
	}

	resp, _, err := opts.client.NetworkPeeringApi.CreatePeeringConnectionWithParams(ctx, params).Execute()
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

func createPeeringConnectionBuilder() *cobra.Command {
	opts := createPeeringConnectionOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createPeeringConnection",
		Short: "Create One New Network Peering Connection",
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

type createPeeringContainerOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *createPeeringContainerOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *createPeeringContainerOpts) readData() (*admin.CloudProviderContainer, error) {
	var out *admin.CloudProviderContainer

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

func (opts *createPeeringContainerOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.CreatePeeringContainerApiParams{
		GroupId: opts.groupId,

		CloudProviderContainer: data,
	}

	resp, _, err := opts.client.NetworkPeeringApi.CreatePeeringContainerWithParams(ctx, params).Execute()
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

func createPeeringContainerBuilder() *cobra.Command {
	opts := createPeeringContainerOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "createPeeringContainer",
		Short: "Create One New Network Peering Container",
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

type deletePeeringConnectionOpts struct {
	client  *admin.APIClient
	groupId string
	peerId  string
}

func (opts *deletePeeringConnectionOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *deletePeeringConnectionOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.DeletePeeringConnectionApiParams{
		GroupId: opts.groupId,
		PeerId:  opts.peerId,
	}

	resp, _, err := opts.client.NetworkPeeringApi.DeletePeeringConnectionWithParams(ctx, params).Execute()
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

func deletePeeringConnectionBuilder() *cobra.Command {
	opts := deletePeeringConnectionOpts{}
	cmd := &cobra.Command{
		Use:   "deletePeeringConnection",
		Short: "Remove One Existing Network Peering Connection",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.peerId, "peerId", "", `Unique 24-hexadecimal digit string that identifies the network peering connection that you want to delete.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("peerId")
	return cmd
}

type deletePeeringContainerOpts struct {
	client      *admin.APIClient
	groupId     string
	containerId string
}

func (opts *deletePeeringContainerOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *deletePeeringContainerOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.DeletePeeringContainerApiParams{
		GroupId:     opts.groupId,
		ContainerId: opts.containerId,
	}

	resp, _, err := opts.client.NetworkPeeringApi.DeletePeeringContainerWithParams(ctx, params).Execute()
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

func deletePeeringContainerBuilder() *cobra.Command {
	opts := deletePeeringContainerOpts{}
	cmd := &cobra.Command{
		Use:   "deletePeeringContainer",
		Short: "Remove One Network Peering Container",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.containerId, "containerId", "", `Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("containerId")
	return cmd
}

type disablePeeringOpts struct {
	client  *admin.APIClient
	groupId string

	filename string
	fs       afero.Fs
}

func (opts *disablePeeringOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *disablePeeringOpts) readData() (*admin.PrivateIPMode, error) {
	var out *admin.PrivateIPMode

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

func (opts *disablePeeringOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.DisablePeeringApiParams{
		GroupId: opts.groupId,

		PrivateIPMode: data,
	}

	resp, _, err := opts.client.NetworkPeeringApi.DisablePeeringWithParams(ctx, params).Execute()
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

func disablePeeringBuilder() *cobra.Command {
	opts := disablePeeringOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "disablePeering",
		Short: "Disable Connect via Peering Only Mode for One Project",
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

type getPeeringConnectionOpts struct {
	client  *admin.APIClient
	groupId string
	peerId  string
}

func (opts *getPeeringConnectionOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getPeeringConnectionOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.GetPeeringConnectionApiParams{
		GroupId: opts.groupId,
		PeerId:  opts.peerId,
	}

	resp, _, err := opts.client.NetworkPeeringApi.GetPeeringConnectionWithParams(ctx, params).Execute()
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

func getPeeringConnectionBuilder() *cobra.Command {
	opts := getPeeringConnectionOpts{}
	cmd := &cobra.Command{
		Use:   "getPeeringConnection",
		Short: "Return One Network Peering Connection in One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.peerId, "peerId", "", `Unique 24-hexadecimal digit string that identifies the network peering connection that you want to retrieve.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("peerId")
	return cmd
}

type getPeeringContainerOpts struct {
	client      *admin.APIClient
	groupId     string
	containerId string
}

func (opts *getPeeringContainerOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *getPeeringContainerOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.GetPeeringContainerApiParams{
		GroupId:     opts.groupId,
		ContainerId: opts.containerId,
	}

	resp, _, err := opts.client.NetworkPeeringApi.GetPeeringContainerWithParams(ctx, params).Execute()
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

func getPeeringContainerBuilder() *cobra.Command {
	opts := getPeeringContainerOpts{}
	cmd := &cobra.Command{
		Use:   "getPeeringContainer",
		Short: "Return One Network Peering Container",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.containerId, "containerId", "", `Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("containerId")
	return cmd
}

type listPeeringConnectionsOpts struct {
	client       *admin.APIClient
	groupId      string
	includeCount bool
	itemsPerPage int
	pageNum      int
	providerName string
}

func (opts *listPeeringConnectionsOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listPeeringConnectionsOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.ListPeeringConnectionsApiParams{
		GroupId:      opts.groupId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
		ProviderName: &opts.providerName,
	}

	resp, _, err := opts.client.NetworkPeeringApi.ListPeeringConnectionsWithParams(ctx, params).Execute()
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

func listPeeringConnectionsBuilder() *cobra.Command {
	opts := listPeeringConnectionsOpts{}
	cmd := &cobra.Command{
		Use:   "listPeeringConnections",
		Short: "Return All Network Peering Connections in One Project",
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
	cmd.Flags().StringVar(&opts.providerName, "providerName", "&quot;AWS&quot;", `Cloud service provider to use for this VPC peering connection.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

type listPeeringContainerByCloudProviderOpts struct {
	client       *admin.APIClient
	groupId      string
	providerName string
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listPeeringContainerByCloudProviderOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listPeeringContainerByCloudProviderOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.ListPeeringContainerByCloudProviderApiParams{
		GroupId:      opts.groupId,
		ProviderName: &opts.providerName,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}

	resp, _, err := opts.client.NetworkPeeringApi.ListPeeringContainerByCloudProviderWithParams(ctx, params).Execute()
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

func listPeeringContainerByCloudProviderBuilder() *cobra.Command {
	opts := listPeeringContainerByCloudProviderOpts{}
	cmd := &cobra.Command{
		Use:   "listPeeringContainerByCloudProvider",
		Short: "Return All Network Peering Containers in One Project for One Cloud Provider",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.providerName, "providerName", "&quot;AWS&quot;", `Cloud service provider that serves the desired network peering containers.`)
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("providerName")
	return cmd
}

type listPeeringContainersOpts struct {
	client       *admin.APIClient
	groupId      string
	includeCount bool
	itemsPerPage int
	pageNum      int
}

func (opts *listPeeringContainersOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *listPeeringContainersOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.ListPeeringContainersApiParams{
		GroupId:      opts.groupId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}

	resp, _, err := opts.client.NetworkPeeringApi.ListPeeringContainersWithParams(ctx, params).Execute()
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

func listPeeringContainersBuilder() *cobra.Command {
	opts := listPeeringContainersOpts{}
	cmd := &cobra.Command{
		Use:   "listPeeringContainers",
		Short: "Return All Network Peering Containers in One Project",
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

type updatePeeringConnectionOpts struct {
	client  *admin.APIClient
	groupId string
	peerId  string

	filename string
	fs       afero.Fs
}

func (opts *updatePeeringConnectionOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *updatePeeringConnectionOpts) readData() (*admin.BaseNetworkPeeringConnectionSettings, error) {
	var out *admin.BaseNetworkPeeringConnectionSettings

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

func (opts *updatePeeringConnectionOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.UpdatePeeringConnectionApiParams{
		GroupId: opts.groupId,
		PeerId:  opts.peerId,

		BaseNetworkPeeringConnectionSettings: data,
	}

	resp, _, err := opts.client.NetworkPeeringApi.UpdatePeeringConnectionWithParams(ctx, params).Execute()
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

func updatePeeringConnectionBuilder() *cobra.Command {
	opts := updatePeeringConnectionOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updatePeeringConnection",
		Short: "Update One New Network Peering Connection",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.peerId, "peerId", "", `Unique 24-hexadecimal digit string that identifies the network peering connection that you want to update.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("peerId")
	return cmd
}

type updatePeeringContainerOpts struct {
	client      *admin.APIClient
	groupId     string
	containerId string

	filename string
	fs       afero.Fs
}

func (opts *updatePeeringContainerOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *updatePeeringContainerOpts) readData() (*admin.CloudProviderContainer, error) {
	var out *admin.CloudProviderContainer

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

func (opts *updatePeeringContainerOpts) run(ctx context.Context, w io.Writer) error {
	data, errData := opts.readData()
	if errData != nil {
		return errData
	}
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.UpdatePeeringContainerApiParams{
		GroupId:     opts.groupId,
		ContainerId: opts.containerId,

		CloudProviderContainer: data,
	}

	resp, _, err := opts.client.NetworkPeeringApi.UpdatePeeringContainerWithParams(ctx, params).Execute()
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

func updatePeeringContainerBuilder() *cobra.Command {
	opts := updatePeeringContainerOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "updatePeeringContainer",
		Short: "Update One Network Peering Container",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)
	cmd.Flags().StringVar(&opts.containerId, "containerId", "", `Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that you want to remove.`)

	cmd.Flags().StringVarP(&opts.filename, "file", "f", "", "Path to an optional JSON configuration file if not passed stdin is expected")

	_ = cmd.MarkFlagRequired("groupId")
	_ = cmd.MarkFlagRequired("containerId")
	return cmd
}

type verifyConnectViaPeeringOnlyModeForOneProjectOpts struct {
	client  *admin.APIClient
	groupId string
}

func (opts *verifyConnectViaPeeringOnlyModeForOneProjectOpts) preRun() (err error) {
	opts.client, err = newClientWithAuth()
	return err
}

func (opts *verifyConnectViaPeeringOnlyModeForOneProjectOpts) run(ctx context.Context, w io.Writer) error {
	if opts.groupId == "" {
		opts.groupId = config.ProjectID()
	}

	params := &admin.VerifyConnectViaPeeringOnlyModeForOneProjectApiParams{
		GroupId: opts.groupId,
	}

	resp, _, err := opts.client.NetworkPeeringApi.VerifyConnectViaPeeringOnlyModeForOneProjectWithParams(ctx, params).Execute()
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

func verifyConnectViaPeeringOnlyModeForOneProjectBuilder() *cobra.Command {
	opts := verifyConnectViaPeeringOnlyModeForOneProjectOpts{}
	cmd := &cobra.Command{
		Use:   "verifyConnectViaPeeringOnlyModeForOneProject",
		Short: "Verify Connect via Peering Only Mode for One Project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.preRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.run(cmd.Context(), cmd.OutOrStdout())
		},
	}
	cmd.Flags().StringVar(&opts.groupId, "projectId", "", `Unique 24-hexadecimal digit string that identifies your project.`)

	_ = cmd.MarkFlagRequired("groupId")
	return cmd
}

func networkPeeringBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use: "networkPeering",
		Short: `Returns, adds, edits, and removes network peering containers and peering connections.
When you deploy an M10+ dedicated cluster, Atlas creates a VPC for the selected provider and region or regions if no existing VPC or VPC peering connection exists for that provider and region. Atlas assigns the VPC a Classless Inter-Domain Routing (CIDR) block.`,
	}
	cmd.AddCommand(
		createPeeringConnectionBuilder(),
		createPeeringContainerBuilder(),
		deletePeeringConnectionBuilder(),
		deletePeeringContainerBuilder(),
		disablePeeringBuilder(),
		getPeeringConnectionBuilder(),
		getPeeringContainerBuilder(),
		listPeeringConnectionsBuilder(),
		listPeeringContainerByCloudProviderBuilder(),
		listPeeringContainersBuilder(),
		updatePeeringConnectionBuilder(),
		updatePeeringContainerBuilder(),
		verifyConnectViaPeeringOnlyModeForOneProjectBuilder(),
	)
	return cmd
}
