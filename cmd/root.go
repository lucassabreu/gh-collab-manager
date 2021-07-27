/*
Copyright © 2021 lucassabreu

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
	"os"

	"github.com/lucassabreu/gh-collab-manager/internal"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var (
	cfgFile string

	version   string
	commit    string
	buildDate string
)

const GH_TOKEN_KEY = "github-token"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "gh-collab-manager",
	Short:         "Add and remove users from repositories",
	SilenceErrors: true,
	SilenceUsage:  true,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {

		if v, _ := cmd.Flags().GetBool("version"); v {
			cmd.Printf(
				"Version: %s, Commit: %s, Build At: %s\n",
				version, commit, buildDate,
			)
			return nil
		}

		stringRepos, _ := cmd.Flags().GetStringArray("repo")
		toadd, _ := cmd.Flags().GetStringArray("add")
		toRemove, _ := cmd.Flags().GetStringArray("remove")

		repos, err := internal.MapStringItoRepository(stringRepos)
		if err != nil {
			return err
		}

		return internal.Execute(
			viper.GetString(GH_TOKEN_KEY),
			repos,
			toadd, toRemove,
		)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(v, c, d string) {
	version = v
	commit = c
	buildDate = d

	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gh-collab-manager.yaml)")

	rootCmd.Flags().StringP(GH_TOKEN_KEY, "t", "", "github token to access the api (defaults to $GITHUB_TOKEN)")
	_ = viper.BindPFlag(GH_TOKEN_KEY, rootCmd.Flags().Lookup(GH_TOKEN_KEY))
	_ = viper.BindEnv(GH_TOKEN_KEY, "GITHUB_TOKEN")

	rootCmd.Flags().StringArrayP("repo", "R", []string{}, "repositories to add/remove collaborators")
	rootCmd.Flags().StringArrayP("add", "a", []string{}, "handle of collaborators to invite to the repositories")
	rootCmd.Flags().StringArrayP("remove", "r", []string{}, "handle of collaborators to remove from the repositories")

	rootCmd.Flags().BoolP("version", "v", false, "shows cli version")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".gh-collab-manager" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gh-collab-manager")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
