/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"Caesar/api"
	"Caesar/app/reset"

	"github.com/spf13/cobra"
)

var (
	document string
	fileName string
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "reset path json of hits is 0",
	Long: api.Banner + `

将路径json字典的hits重置为0
example:
     caesar reset -d ~/path/  
     caesar reset -f ~/paths.json
`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(document) > 0 && len(fileName) == 0 {
			reset.SetupHitsOfZeroInDocument(document)
			return

		} else if len(document) == 0 && len(fileName) > 0 {
			reset.SetupHitsOfZeroInFile(fileName)
			return

		} else {
			if err := cmd.Help(); err != nil {
				println(err.Error())
			}
			return

		}
	},
}

func init() {

	resetCmd.Flags().StringVarP(&document, "document", "d", "", "read document to reset hits 0")
	resetCmd.Flags().StringVarP(&fileName, "filename", "f", "", "read file to reset hits 0")
	resetCmd.Flags().SortFlags = false
	rootCmd.AddCommand(resetCmd)

}
