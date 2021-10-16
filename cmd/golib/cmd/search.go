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
	"strings"
	"text/tabwriter"

	"github.com/TheBoringDude/golibraries"
	"github.com/spf13/cobra"
)

var searchOptionsPlatforms = []string{}
var searchOptionsPage uint64
var searchOptionsPerPage uint64
var searchOptionsLanguages = []string{}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for names in package manager repositories",
	Long:  `Search for a project in package manager repositories.`,
	Run: func(cmd *cobra.Command, args []string) {
		q := strings.Join(args, "")
		if q == "" {
			fmt.Println("Nothing to search...")
			return
		}

		if ApiKey == "" {
			fmt.Println("ApiKey is empty!")
			return
		}

		golib := golibraries.New(ApiKey)

		search, err := golib.ProjectSearch(q, golibraries.ProjectSearchOptions{PerPage: int(searchOptionsPerPage), Page: int(searchOptionsPage), Platforms: searchOptionsPlatforms, Languages: searchOptionsLanguages})
		if err != nil {
			log.Fatalln(err)
		}

		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)

		if len(search) == 0 {
			fmt.Println("\n  No projects found with query:", q)
			return
		}

		fmt.Printf("\n Search results for `%s`:\n", q)
		fmt.Println()
		fmt.Fprintln(writer, "  ", "No.\t Project Name\t Project Description")

		for index, v := range search {
			fmt.Fprintln(writer, "  ", index+1, ")", "\t", v.Name, "\t", v.Description)
		}

		writer.Flush()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringSliceVarP(&searchOptionsPlatforms, "platforms", "p", []string{}, "Search only on specific platforms")
	searchCmd.Flags().StringSliceVarP(&searchOptionsLanguages, "languages", "l", []string{}, "Search only on specific platforms")
	searchCmd.Flags().Uint64Var(&searchOptionsPage, "page", 1, "Page of search result.")
	searchCmd.Flags().Uint64Var(&searchOptionsPerPage, "per-page", 10, "Number of search result items to return.")
}
