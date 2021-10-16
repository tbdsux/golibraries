/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"os"
	"strconv"

	"github.com/TheBoringDude/golibraries"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var srPlatform string
var srProject string

// sourcerankCmd represents the sourcerank command
var sourcerankCmd = &cobra.Command{
	Use:   "sourcerank",
	Short: "Get a project's SOURCERANK score",
	Long: `Get breakdown of a SOURCERANK score of a given project 
in the specific platform.`,
	Run: func(cmd *cobra.Command, args []string) {
		if ApiKey == "" {
			fmt.Println("ApiKey is empty!")
			return
		}

		golib := golibraries.New(ApiKey)

		// query project sourcerank
		sr, err := golib.ProjectSourceRank(srPlatform, srProject)
		if err != nil {
			log.Fatalln(err)
		}

		table := tablewriter.NewWriter(os.Stdout)

		// write output
		fmt.Printf("\n Project: %s", srProject)
		fmt.Printf("\n Platform: %s\n\n", srPlatform)

		table.SetBorder(false)
		table.AppendBulk([][]string{
			{"  Basic Info Present", strconv.Itoa(int(sr.BasicInfoPresent))},
			{"  Repository Present", strconv.Itoa(int(sr.RepositoryPresent))},
			{"  Readme Present", strconv.Itoa(int(sr.ReadmePresent))},
			{"  License Present", strconv.Itoa(int(sr.LicensePresent))},
			{"  Versions Present", strconv.Itoa(int(sr.VersionsPresent))},
			{"  Follows Semver", strconv.Itoa(int(sr.FollowsSemver))},
			{"  Recent Release", strconv.Itoa(int(sr.RecentRelease))},
			{"  Not Brand New", strconv.Itoa(int(sr.NotBrandNew))},
			{"  One Point Oh", strconv.Itoa(int(sr.OnePointOh))},
			{"  Dependent Projects", strconv.Itoa(int(sr.DependentProjects))},
			{"  Dependent Repositories", strconv.Itoa(int(sr.DependentRepositories))},
			{"  Stars", strconv.Itoa(int(sr.Stars))},
			{"  Contributors", strconv.Itoa(int(sr.Contributors))},
			{"  Subscribers", strconv.Itoa(int(sr.Subscribers))},
			{"  All Prereleases", strconv.Itoa(int(sr.AllPrereleases))},
			{"  Any Outdated Dependencies", strconv.Itoa(int(sr.AnyOutdatedDependencies))},
			{"  Is Deprecated", strconv.Itoa(int(sr.IsDeprecated))},
			{"  Is Unmaintained", strconv.Itoa(int(sr.IsUnmaintained))},
			{"  Is Removed", strconv.Itoa(int(sr.IsRemoved))}})

		table.Render()

	},
}

func init() {
	rootCmd.AddCommand(sourcerankCmd)

	sourcerankCmd.Flags().StringVarP(&srProject, "project", "p", "", "Project name")
	sourcerankCmd.Flags().StringVar(&srPlatform, "platform", "", "Project platform")

	sourcerankCmd.MarkFlagRequired("platform")
	sourcerankCmd.MarkFlagRequired("project")
}
