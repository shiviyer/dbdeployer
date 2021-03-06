// DBDeployer - The MySQL Sandbox
// Copyright © 2006-2018 Giuseppe Maxia
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

package cmd

import (
	"fmt"
	"github.com/datacharmer/dbdeployer/common"
	"github.com/spf13/cobra"
	"io/ioutil"
	"path"
)

// Shows the MySQL versions available in $SANDBOX_BINARY
// (default $HOME/opt/mysql)
func ShowVersions(cmd *cobra.Command, args []string) {
	Basedir := GetAbsolutePathFromFlag(cmd, "sandbox-binary")
	files, err := ioutil.ReadDir(Basedir)
	common.ErrCheckExitf(err, 1, "error reading directory %s: %s", Basedir, err)
	var dirs []string
	for _, f := range files {
		fname := f.Name()
		fmode := f.Mode()
		//fmt.Printf("%#v\n", fmode)
		if fmode.IsDir() {
			//fmt.Println(fname)
			mysqld := path.Join(Basedir, fname, "bin", "mysqld")
			if common.FileExists(mysqld) {
				dirs = append(dirs, fname)
			}
		}
	}
	maxWidth := 80
	maxLen := 0
	for _, dir := range dirs {
		if len(dir) > maxLen {
			maxLen = len(dir)
		}
	}
	fmt.Printf("Basedir: %s\n", Basedir)
	columns := int(maxWidth / (maxLen + 2))
	mask := fmt.Sprintf("%%-%ds", maxLen+2)
	count := 0
	for _, dir := range dirs {
		fmt.Printf(mask, dir)
		count += 1
		if count > columns {
			count = 0
			fmt.Println("")
		}
	}
	fmt.Println("")
}

// versionsCmd represents the versions command
var versionsCmd = &cobra.Command{
	Use:     "versions",
	Aliases: []string{"available"},
	Short:   "List available versions",
	Long:    ``,
	Run:     ShowVersions,
}

func init() {
	rootCmd.AddCommand(versionsCmd)
}
