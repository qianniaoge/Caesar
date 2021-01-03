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
	"Caesar/app/audit"
	"Caesar/internal/library/boot"
	"Caesar/internal/relation"
	"Caesar/pkg/record"

	"github.com/spf13/cobra"
)

var (
	target  string
	text    string
	flag    string
	threads int
	delay   int
)

// auditCmd represents the audit command
var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "sensitive files found",
	Long: api.Banner + `

敏感文件扫描功能.
example: 
   caesar audit --target-address=http://127.0.0.1 --flag=common,php -g 1
   caesar audit --target-address=target.txt --flag=common,php
   caesar audit --read=requests.txt --flag=common,php
`,

	Run: func(cmd *cobra.Command, args []string) {

		if len(flag) == 0 && len(target) == 0 && len(text) == 0 {
			if err := cmd.Help(); err != nil {
				println(err.Error())
			}
			return
		}

		if len(target) == 0 && len(text) == 0 {
			println("You must enter target or read")
			return
		}

		if len(flag) == 0 {
			println("You must enter flag")
			return
		}

		// 开启日志记录器
		record.Logs(relation.Paths.Result+"/console.log", relation.Engine.Silence)

		if len(target) > 0 && len(flag) > 0 {

			// 读取url地址
			audit.StartSensitiveFoundFromAddress(target, flag, threads, delay)
			return

		}

		if len(text) > 0 && len(flag) > 0 {
			// 读取url地址
			audit.StartSensitiveFoundFromText(text, flag, threads, delay)
			return

		}

	},
}

func init() {
	auditCmd.Flags().StringVarP(&target, "target-address", "t", "", "scan target (type: txt file or address string)")
	auditCmd.Flags().StringVarP(&text, "read", "r", "", "read file to request")
	auditCmd.Flags().StringVarP(&flag, "flag", "f", "", "dict type (example: common or common,asp) , All dict is: "+boot.GetFlag("assets/directory"))
	auditCmd.Flags().IntVarP(&threads, "threads", "g", 3, "The threads num")
	auditCmd.Flags().IntVarP(&delay, "delay", "d", 0, "time delay")
	auditCmd.Flags().SortFlags = false
	rootCmd.AddCommand(auditCmd)

}
