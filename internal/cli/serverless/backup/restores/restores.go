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

package restores

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	const use = "restores"
	cmd := &cobra.Command{
		Use:        use,
		Short:      "Manage cloud backup restore jobs for your project.",
		Aliases:    cli.GenerateAliases(use),
		Deprecated: "please use the 'atlas backup restores' command instead. For the migration guide and timeline, visit: https://dochub.mongodb.org/core/flex-migration",
	}

	cmd.AddCommand(
		ListBuilder(),
		DescribeBuilder(),
		WatchBuilder(),
		CreateBuilder(),
	)

	return cmd
}
