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

package main

import (
	"io"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"

	"embed"

	"github.com/spf13/cobra"
)

//go:embed "templates/md.gotmpl"
var tmplFS embed.FS
var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("templates").Funcs(template.FuncMap{
		"replace": func(old, new, s string) string {
			return strings.ReplaceAll(s, old, new)
		},
		"replaceRegex": func(pattern, replacement, input string) string {
			return regexp.MustCompile(pattern).ReplaceAllString(input, replacement)
		},
		"alias": func(cmd *cobra.Command) []string {
			return alias(cmd, false)
		},
	}).ParseFS(tmplFS, "templates/md.gotmpl")).Lookup("md.gotmpl")
}

func applyTemplate(cmd *cobra.Command, w io.Writer) error {
	return tmpl.Execute(w, cmd)
}

func genMdDocs(cmd *cobra.Command, dir string) error {
	return walk(cmd, func(cmd *cobra.Command) error {
		if !cmd.IsAvailableCommand() || cmd.IsAdditionalHelpTopicCommand() {
			return nil
		}
		fileName := strings.ReplaceAll(cmd.CommandPath(), " ", "_") + ".md"
		fileName = path.Join(dir, fileName)
		f, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer f.Close()
		return applyTemplate(cmd, f)
	})
}
