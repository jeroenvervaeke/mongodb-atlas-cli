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
	"strings"

	"github.com/spf13/cobra"
)

func splitUse(cmd *cobra.Command) (string, string) {
	useSegments := strings.SplitN(cmd.Use, " ", 2)
	suffix := ""
	if len(useSegments) > 1 {
		suffix = useSegments[1]
	}
	if !cmd.DisableFlagsInUseLine && cmd.HasAvailableFlags() {
		suffix += " [flags]"
		suffix = strings.TrimSpace(suffix)
	}
	return useSegments[0], suffix
}

func alias(cmd *cobra.Command, includeUse bool) []string {
	path := [][]string{}
	currentCmd := cmd
	for currentCmd != nil {
		aliases := append([]string{currentCmd.Name()}, currentCmd.Aliases...)

		path = append([][]string{aliases}, path...)
		currentCmd = currentCmd.Parent()
	}
	if includeUse {
		_, suffix := splitUse(cmd)
		path = append(path, []string{suffix})
	}
	for len(path) > 1 {
		newPaths := []string{}
		for _, topLevel := range path[0] {
			for _, secondLevel := range path[1] {
				newPaths = append(newPaths, topLevel+" "+secondLevel)
			}
		}
		path = path[2:]
		path = append([][]string{newPaths}, path...)
	}
	return path[0]
}

func walk(cmd *cobra.Command, f func(cmd *cobra.Command) error) error {
	if err := f(cmd); err != nil {
		return err
	}
	for _, childCmd := range cmd.Commands() {
		if err := walk(childCmd, f); err != nil {
			return err
		}
	}
	return nil
}
