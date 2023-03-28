// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

	pubsub "github.com/erikwilliamsa/gcloudps/pubsub"
	"github.com/erikwilliamsa/gcloudps/workers"
	"github.com/spf13/cobra"
)

var generateCount int

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a number of generic messages and publisht them to the given topic.",
	Args: func(cmd *cobra.Command, args []string) error {
		flags := make(map[string]string)
		flags["count"] = string(generateCount)
		return CheckRequiredFlags(flags)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
		ctx, _, topic := initClient()
		pc := &pubsub.PublishClient{Context: ctx, Topic: topic, Batch: false}
		pw := workers.PublisheWorker{Publish: pc}
		pw.GenerateMessages(generateCount)

	},
}

func init() {
	pubCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntVarP(&generateCount, "count", "c", 0, "The number of messages to generate and publish.")

}
