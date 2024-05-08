/*
Copyright 2020 the Velero contributors.

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

package version

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/sharework-cn/idp/pkg/buildinfo"
)

func NewCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "version",
		Short: "Print the idp version",
		Run: func(c *cobra.Command, args []string) {
			printVersion(os.Stdout)
		},
	}
	return c
}

func printVersion(w io.Writer) {
	fmt.Fprintln(w, "idp:")
	fmt.Fprintf(w, "\tVersion: %s\n", buildinfo.Version)
	fmt.Fprintf(w, "\tGit commit: %s\n", buildinfo.FormattedGitSHA())
}
