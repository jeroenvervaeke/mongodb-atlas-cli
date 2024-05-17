// Copyright 2022 MongoDB Inc
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
	"log"
	"os"

	"github.com/mongodb-labs/cobra2snooty"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/root"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func setDisableAutoGenTag(cmd *cobra.Command) {
	cmd.DisableAutoGenTag = true
	for _, cmd := range cmd.Commands() {
		setDisableAutoGenTag(cmd)
	}
}

func run() error {
	atlasBuilder := root.Builder()
	atlasBuilder.InitDefaultCompletionCmd()

	setDisableAutoGenTag(atlasBuilder)

	docs := map[string]func(cmd *cobra.Command, dir string) error{
		"./docs/command":         cobra2snooty.GenTreeDocs,
		"./docs/command-cobramd": doc.GenMarkdownTree,
		"./docs/command-md":      genMdDocs,
		"./docs/command-json":    genJsonDocs,
	}

	for dir, f := range docs {
		const docsPermissions = 0766
		if err := os.RemoveAll(dir); err != nil {
			return err
		}
		if err := os.MkdirAll(dir, docsPermissions); err != nil {
			return err
		}
		if err := f(atlasBuilder, dir); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
