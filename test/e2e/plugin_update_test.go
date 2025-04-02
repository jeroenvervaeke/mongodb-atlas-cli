// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build e2e || (atlas && plugin && update)

package e2e_test

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPluginUpdate(t *testing.T) {
	_ = tempConfigFolder(t)

	g := newAtlasE2ETestGenerator(t, withSnapshot())
	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)
	runPluginUpdateTest(g, cliPath, "Update without specifying version", false, examplePluginRepository, "v1.0.38", "")
	runPluginUpdateTest(g, cliPath, "Update with specifying version", false, examplePluginName, "v1.0.38", "v2.0.3")
	runPluginUpdateTest(g, cliPath, "Update with specifying latest version", false, examplePluginName, "v1.0.38", "latest")
	runPluginUpdateTest(g, cliPath, "Update using --all flag", false, "--all", "v1.0.34", "")
	runPluginUpdateTest(g, cliPath, "Update with lower version", true, examplePluginName, "v1.0.36", "v1.0.34")
	runPluginUpdateTest(g, cliPath, "Update with same version", true, examplePluginRepository, "v1.0.36", "v1.0.36")
	runPluginUpdateTest(g, cliPath, "Update with too many arguments", true, examplePluginName+" --all", "v1.0.34", "v2.0.0")
	runPluginUpdateTest(g, cliPath, "Update without any values", true, "", "v1.0.34", "v2.0.0")
}

func runPluginUpdateTest(g *atlasE2ETestGenerator, cliPath string, testName string, requireError bool, pluginValue string, initialVersion string, updateVersion string) {
	g.t.Helper()
	installExamplePlugin(g.t, cliPath, initialVersion)

	g.Run(testName, func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		if updateVersion != "" && pluginValue != "--all" {
			pluginValue = fmt.Sprintf("%s@%s", pluginValue, updateVersion)
		}

		cmd := exec.Command(cliPath,
			"plugin",
			"update",
			pluginValue)
		resp, err := RunAndGetStdOut(cmd)
		if requireError {
			require.Error(t, err, string(resp))
		} else {
			require.NoError(t, err, string(resp))
		}
	})

	deleteAllPlugins(g.t)
}
