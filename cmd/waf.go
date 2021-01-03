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
	"Caesar/app/waf"

	"github.com/spf13/cobra"
)

var (
	wafTarget  string
	wafThreads int
	wafDelay   int
)

// wafCmd represents the waf command
var wafCmd = &cobra.Command{
	Use:   "waf",
	Short: "WAF check tool",
	Long: api.Banner + `

检查目标是否受WAF保护
example:
   caesar waf --target-address=target.txt -g 5
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(wafTarget) == 0 {
			if err := cmd.Help(); err != nil {
				println(err.Error())
			}
			return
		}
		waf.StartWafCheck(wafTarget, wafThreads, wafDelay)
	},
}

func init() {

	wafCmd.Flags().StringVarP(&wafTarget, "target-address", "t", "", "scan target (type: txt file or address string)")
	wafCmd.Flags().IntVarP(&wafThreads, "threads", "g", 3, "The threads num")
	wafCmd.Flags().IntVarP(&wafDelay, "delay", "d", 0, "time delay")
	wafCmd.Flags().SortFlags = false
	rootCmd.AddCommand(wafCmd)

}
