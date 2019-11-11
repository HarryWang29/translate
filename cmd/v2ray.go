/*
Copyright © 2019 HarryWang <wrz890829@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"translate/translate"
)

// VmessCmd represents the Vmess command
//todo 优化终端提示
var VmessCmd = &cobra.Command{
	Use:   "vmess",
	Short: "订阅为vmess协议",
	Long:  `订阅为vmess协议，默认使用神机规则`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("args is error")
		}
		err := translate.Run(cmd.CalledAs(), args[0], ruleName, subLinks)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(VmessCmd)
	VmessCmd.PersistentFlags().StringSliceVar(&subLinks, "subLink", []string{}, "订阅链接")
	//VmessCmd.PersistentFlags().StringVar(&ruleName, "ruleName", "", "订阅链接")
	ruleName = "ConnersHua" //目前先只支持神机，其他后期再进行支持

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// VmessCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// VmessCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
