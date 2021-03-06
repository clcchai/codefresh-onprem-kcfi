/*
Copyright The Codefresh Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (

	"io"
	"fmt"
	"github.com/spf13/cobra"

	"helm.sh/helm/v3/cmd/helm/require"
	"github.com/codefresh-io/kcfi/pkg/embeded/version"
)



func cfVersionCmd(out io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "version",
		Short: "shows kcfi version",
		Long:  "",
		Args:  require.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Fprintf(out, "version:  %s\nrevision: %s\n", version.Version, version.GitRevision)
			return nil
		},
	}

	origHelpFunc := cmd.HelpFunc()
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		hideKubeFlags(cmd)
		hideHelmCommonFlags(cmd)
		origHelpFunc(cmd, args)
	})
	return cmd
}