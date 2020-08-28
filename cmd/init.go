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
	"fmt"
	"github.com/mimatache/forge/internal/manifest"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Flag string

const (
	FORCE = "force"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates a default Forge file",
	Long:  `Creates a default forge file that can also be used to verify forge executions`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("Forge"); err == nil {
			force, err := cmd.Flags().GetBool(FORCE)
			if err != nil {
				log.Fatalf("Could not read flag %s", "force")
			}
			if force {
				fmt.Println("Removing old Forge files")
				err := os.Remove("Forge")
				if err != nil {
					log.Fatal("Could not remove Forge files")
				}
			} else {
				fmt.Println("Forge already exists in current directory")
				os.Exit(0)
			}
		}
		contents, err := yaml.Marshal(manifest.InitializeDemoForge())
		if err != nil {
			log.Fatalf("could not create Forge: %s", err)
		}
		log.Println("Creating forge file")
		err = ioutil.WriteFile("Forge", contents, 0644)
		if err != nil {
			log.Fatalf("could not create Forge file: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().Bool(FORCE, false, "Force forge file creation. This will remove any existing forge files")
}
