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
	"encoding/base64"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	width  int
	height int
	data   string
	out    string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "b64ToImg",
	Short: "",
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

	rootCmd.Flags().IntVarP(&width, "width", "w", 100, "image width")
	rootCmd.Flags().StringVarP(&data, "data", "d", "", "image base64 context")
	rootCmd.Flags().StringVarP(&out, "out", "o", "", "image output path")

}

func do() {
	src, err := ioutil.ReadFile(data)
	if err != nil {
		panic(err)
	}

	x := strings.Split(string(src), ",")

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(x[1]))

	config, format, err := image.DecodeConfig(reader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Width:", config.Width, "Height:", config.Height, "Format:", format)

	dist, _ := base64.StdEncoding.DecodeString(string(x[1]))

	if out == "" {
		out = "tmp.png"
	}

	f, _ := os.OpenFile(out, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)

}
