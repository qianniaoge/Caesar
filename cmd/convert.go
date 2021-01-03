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
	"Caesar/app/convert"

	"github.com/spf13/cobra"
)

var (
	convertDocument string
	convertText     string
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert path txt to path json (path file must end of .txt)",
	Long: api.Banner + `.

将普通路径文件转换为Caesar能识别的json格式的路径文件
example:
     caesar convert -d ~/path/  
     caesar convert -f ~/asp.txt
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(convertText) == 0 && len(convertDocument) == 0 {
			if err := cmd.Help(); err != nil {
				println(err.Error())
			}
			return
		}

		if len(convertText) > 0 {
			convert.TextToJsonOfFile(convertText)
		} else if len(convertDocument) > 0 {
			convert.TextToJsonOfPath(convertDocument)
		}
	},
}

func init() {
	convertCmd.Flags().StringVarP(&convertDocument, "document", "d", "", "read txt to json, from document(only txt)")
	convertCmd.Flags().StringVarP(&convertText, "file", "f", "", "read txt to json, from file")

	rootCmd.AddCommand(convertCmd)

}
