// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"os"

	"github.com/spf13/cobra"
)

//
const (
	Insert = "insert"
	Remove = "remove"
)

var (
	output   string
	filename string
	str      string
	rowIndex int
	action   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "addItemTocsv",
	Long: `add the new rows item into the csv `,
	Run: func(cmd *cobra.Command, args []string) {
		do()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&filename, "file", "f", "", "want to insert csv file")
	rootCmd.Flags().StringVarP(&output, "output", "o", "output.csv", "output csv filename")
	rootCmd.Flags().StringVarP(&str, "string", "s", "", "want to insert string")
	rootCmd.Flags().IntVarP(&rowIndex, "rows", "r", 0, "want to insert rows index")
	rootCmd.Flags().StringVarP(&action, "action", "a", "insert", "which action do you want ( insert / remove ) default: insert")

}
