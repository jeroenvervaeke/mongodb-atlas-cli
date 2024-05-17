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

package talk

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/spf13/cobra"
)

func prompt(err error) string {
	use := "atlas"
	for _, arg := range os.Args[1:] {
		if strings.Contains(arg, `"`) {
			arg = strings.ReplaceAll(arg, `"`, `\"`)
		}
		if strings.Contains(arg, `"`) || strings.Contains(arg, " ") {
			arg = `"` + arg + `"`
		}
		use += " " + arg
	}
	use = strings.TrimSpace(use)

	outcome := "it ran successfully"
	if err != nil {
		outcome = fmt.Sprintf("I got the error `Error: %v`", err)
	}

	return fmt.Sprintf("I ran the command `%s` and %v, what should I try next?", use, outcome)
}

func Hint(err error) error {
	fmt.Fprint(os.Stderr, "\nHint: ")

	cmd := execute("--skip-repl", "--prompt", prompt(err))
	return cmd.Run()
}

func Run() error {
	return execute().Run()
}

func execute(args ...string) *exec.Cmd {
	cmd := exec.Command("atlas-talk", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "talk",
		Aliases: []string{"assist"},
		Short:   "Interactive help.",
		Args:    require.NoArgs,
		RunE: func(_ *cobra.Command, args []string) error {
			return Run()
		},
	}

	return cmd
}
