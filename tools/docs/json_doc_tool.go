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
	"bytes"
	"encoding/json"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

func genJsonDocs(cmd *cobra.Command, dir string) error {
	return walk(cmd, func(cmd *cobra.Command) error {
		if !cmd.IsAvailableCommand() || cmd.IsAdditionalHelpTopicCommand() {
			return nil
		}
		fileName := strings.ReplaceAll(cmd.CommandPath(), " ", "_") + ".json"
		fileName = path.Join(dir, fileName)
		f, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer f.Close()
		data := map[string]any{
			"CommandPathAliases":      alias(cmd, false),
			"CommandPath":             cmd.CommandPath(),
			"Name":                    cmd.Name(),
			"Use":                     cmd.Use,
			"UseLine":                 cmd.UseLine(),
			"Aliases":                 cmd.Aliases,
			"SuggestFor":              cmd.SuggestFor,
			"Short":                   cmd.Short,
			"GroupID":                 cmd.GroupID,
			"Long":                    cmd.Long,
			"Example":                 cmd.Example,
			"ArgAliases":              cmd.ArgAliases,
			"Deprecated":              cmd.Deprecated,
			"Usage":                   cmd.UsageString(),
			"Runnable":                cmd.Runnable(),
			"HasAvailableSubCommands": cmd.HasAvailableSubCommands(),
		}

		if cmd.HasAvailableFlags() {
			data["Flags"] = cmd.Flags().FlagUsages()
		}

		if cmd.HasAvailableInheritedFlags() {
			data["InheritedFlags"] = cmd.InheritedFlags().FlagUsages()
		}

		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(data); err != nil {
			return err
		}
		_, err = f.Write(buffer.Bytes())
		return err
	})
}
