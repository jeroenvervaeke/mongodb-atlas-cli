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

type createCostExplorerQueryProcessOpts struct {
	client *admin.APIClient
	orgId  string

	filename string
	fs       afero.Fs
	format   string
	tmpl     *template.Template
	resp     *admin.CostExplorerFilterResponse
}

func (opts *createCostExplorerQueryProcessOpts) preRun() (err error) {
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

func (opts *createCostExplorerQueryProcessOpts) readData(r io.Reader) (*admin.CostExplorerFilterRequestBody, error) {
	var out *admin.CostExplorerFilterRequestBody

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

func (opts *createCostExplorerQueryProcessOpts) run(ctx context.Context, r io.Reader) error {
	data, errData := opts.readData(r)
	if errData != nil {
		return errData
	}

	params := &admin.CreateCostExplorerQueryProcessApiParams{
		OrgId: opts.orgId,

		CostExplorerFilterRequestBody: data,
	}

	var err error
	opts.resp, _, err = opts.client.InvoicesApi.CreateCostExplorerQueryProcessWithParams(ctx, params).Execute()
	return err
}

func (opts *createCostExplorerQueryProcessOpts) postRun(_ context.Context, w io.Writer) error {

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

func createCostExplorerQueryProcessBuilder() *cobra.Command {
	opts := createCostExplorerQueryProcessOpts{
		fs: afero.NewOsFs(),
	}
	const use = "createCostExplorerQueryProcess"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Create Cost Explorer query process",
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

type createCostExplorerQueryProcess1Opts struct {
	client *admin.APIClient
	orgId  string
	token  string
	format string
	tmpl   *template.Template
	resp   string
}

func (opts *createCostExplorerQueryProcess1Opts) preRun() (err error) {
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

func (opts *createCostExplorerQueryProcess1Opts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.CreateCostExplorerQueryProcess1ApiParams{
		OrgId: opts.orgId,
		Token: opts.token,
	}

	var err error
	opts.resp, _, err = opts.client.InvoicesApi.CreateCostExplorerQueryProcess1WithParams(ctx, params).Execute()
	return err
}

func (opts *createCostExplorerQueryProcess1Opts) postRun(_ context.Context, w io.Writer) error {

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

func createCostExplorerQueryProcess1Builder() *cobra.Command {
	opts := createCostExplorerQueryProcess1Opts{}
	const use = "createCostExplorerQueryProcess1"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Return results from a given Cost Explorer query, or notify that the results are not ready yet.",
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
	cmd.Flags().StringVar(&opts.token, "token", "", `Unique 64 digit string that identifies the Cost Explorer query.`)

	_ = cmd.MarkFlagRequired("token")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type downloadInvoiceCSVOpts struct {
	client    *admin.APIClient
	orgId     string
	invoiceId string
	format    string
	tmpl      *template.Template
	resp      string
}

func (opts *downloadInvoiceCSVOpts) preRun() (err error) {
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

func (opts *downloadInvoiceCSVOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.DownloadInvoiceCSVApiParams{
		OrgId:     opts.orgId,
		InvoiceId: opts.invoiceId,
	}

	var err error
	opts.resp, _, err = opts.client.InvoicesApi.DownloadInvoiceCSVWithParams(ctx, params).Execute()
	return err
}

func (opts *downloadInvoiceCSVOpts) postRun(_ context.Context, w io.Writer) error {

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

func downloadInvoiceCSVBuilder() *cobra.Command {
	opts := downloadInvoiceCSVOpts{}
	const use = "downloadInvoiceCSV"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Return One Organization Invoice as CSV",
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
	cmd.Flags().StringVar(&opts.invoiceId, "invoiceId", "", `Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.`)

	_ = cmd.MarkFlagRequired("invoiceId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type getInvoiceOpts struct {
	client    *admin.APIClient
	orgId     string
	invoiceId string
	format    string
	tmpl      *template.Template
	resp      string
}

func (opts *getInvoiceOpts) preRun() (err error) {
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

func (opts *getInvoiceOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.GetInvoiceApiParams{
		OrgId:     opts.orgId,
		InvoiceId: opts.invoiceId,
	}

	var err error
	opts.resp, _, err = opts.client.InvoicesApi.GetInvoiceWithParams(ctx, params).Execute()
	return err
}

func (opts *getInvoiceOpts) postRun(_ context.Context, w io.Writer) error {

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

func getInvoiceBuilder() *cobra.Command {
	opts := getInvoiceOpts{}
	const use = "getInvoice"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Return One Organization Invoice",
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
	cmd.Flags().StringVar(&opts.invoiceId, "invoiceId", "", `Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.`)

	_ = cmd.MarkFlagRequired("invoiceId")
	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type listInvoicesOpts struct {
	client       *admin.APIClient
	orgId        string
	includeCount bool
	itemsPerPage int
	pageNum      int
	format       string
	tmpl         *template.Template
	resp         *admin.PaginatedApiInvoice
}

func (opts *listInvoicesOpts) preRun() (err error) {
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

func (opts *listInvoicesOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.ListInvoicesApiParams{
		OrgId:        opts.orgId,
		IncludeCount: &opts.includeCount,
		ItemsPerPage: &opts.itemsPerPage,
		PageNum:      &opts.pageNum,
	}

	var err error
	opts.resp, _, err = opts.client.InvoicesApi.ListInvoicesWithParams(ctx, params).Execute()
	return err
}

func (opts *listInvoicesOpts) postRun(_ context.Context, w io.Writer) error {

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

func listInvoicesBuilder() *cobra.Command {
	opts := listInvoicesOpts{}
	const use = "listInvoices"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Return All Invoices for One Organization",
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
	cmd.Flags().BoolVar(&opts.includeCount, "includeCount", true, `Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.`)
	cmd.Flags().IntVar(&opts.itemsPerPage, "itemsPerPage", 100, `Number of items that the response returns per page.`)
	cmd.Flags().IntVar(&opts.pageNum, "pageNum", 1, `Number of the page that displays the current set of the total objects that the response returns.`)

	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

type listPendingInvoicesOpts struct {
	client *admin.APIClient
	orgId  string
	format string
	tmpl   *template.Template
	resp   *admin.PaginatedApiInvoice
}

func (opts *listPendingInvoicesOpts) preRun() (err error) {
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

func (opts *listPendingInvoicesOpts) run(ctx context.Context, _ io.Reader) error {

	params := &admin.ListPendingInvoicesApiParams{
		OrgId: opts.orgId,
	}

	var err error
	opts.resp, _, err = opts.client.InvoicesApi.ListPendingInvoicesWithParams(ctx, params).Execute()
	return err
}

func (opts *listPendingInvoicesOpts) postRun(_ context.Context, w io.Writer) error {

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

func listPendingInvoicesBuilder() *cobra.Command {
	opts := listPendingInvoicesOpts{}
	const use = "listPendingInvoices"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Return All Pending Invoices for One Organization",
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

	cmd.Flags().StringVar(&opts.format, "format", "", "Format of the output")
	return cmd
}

func invoicesBuilder() *cobra.Command {
	const use = "invoices"
	cmd := &cobra.Command{
		Use:     use,
		Short:   `Returns invoices.`,
		Aliases: cli.GenerateAliases(use),
	}
	cmd.AddCommand(
		createCostExplorerQueryProcessBuilder(),
		createCostExplorerQueryProcess1Builder(),
		downloadInvoiceCSVBuilder(),
		getInvoiceBuilder(),
		listInvoicesBuilder(),
		listPendingInvoicesBuilder(),
	)
	return cmd
}
