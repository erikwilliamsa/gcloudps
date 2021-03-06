// Copyright © 2017 Erik Williams erikwilliamsa@gmail.com
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
	"context"
	"fmt"
	"log"
	"os"

	ps "cloud.google.com/go/pubsub"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	// ProjectName Name of the GCP Project
	ProjectName string
	// TopicName Name of the Topic to be used during execution
	TopicName string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gcloudps",
	Short: "Google Cloud Pub/Sub Cli",
	Long: `Google Cloud Pub/Sub CLI:
	Allows allows publishing and subscribing of message form a GCP topic`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gcloudps.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.PersistentFlags().StringVarP(&ProjectName, "project", "p", "", "The Google cloud project where the client needs to connect")
	RootCmd.PersistentFlags().StringVarP(&TopicName, "topic", "t", "", "Topic for the client to use in the publish or subscribe.")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gcloudps" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gcloudps")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initClient() (context.Context, *ps.Client, *ps.Topic) {
	ctx := context.Background()
	client, err := ps.NewClient(ctx, ProjectName)
	if err != nil {
		log.Fatal(err)
	}
	return ctx, client, client.Topic(TopicName)
}
