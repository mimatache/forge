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
	"bytes"
	"github.com/spf13/cobra"
	"log"
	"mimatache/github.com/forge/internal/manifest"
	"mimatache/github.com/forge/internal/parse"
	"mimatache/github.com/forge/internal/shell"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute the given forgeries and all pre forgeries",
	Long: `This will execute all the given forgeries and any of the forgeries defined in the pre. 
A 'pre' forgery is any forge command that is `,
	Run: func(cmd *cobra.Command, args []string) {

		var execOut bytes.Buffer
		var errOut bytes.Buffer

		executor := shell.Executor{
			Ctx: cmd.Context(),
			Out: &execOut,
			Err: &errOut,
		}
		forgeries, err := parse.GetForgeries(forgeFile, parse.Read)
		if err != nil {
			log.Fatalf("could not read forge file %s: %s", forgeFile, err)
		}
		for _, v := range forgeries {
			v.Execute(executor.Execute)
		}


	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}

func makeMapOfForgeries(forgery ,in, out map[manifest.ForgeryName]manifest.Forgery ) {

}