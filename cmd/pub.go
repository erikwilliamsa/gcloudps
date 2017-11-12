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

	ps "cloud.google.com/go/pubsub"
	"github.com/spf13/cobra"
)

var message string

// pubCmd represents the pub command
var pubCmd = &cobra.Command{
	Use:   "pub",
	Short: "Publish Messages to a Topic",
	Args: func(cmd *cobra.Command, args []string) error {
		flags := make(map[string]string)
		flags["message"] = message
		return CheckRequiredFlags(flags)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Publishing to %s\n", TopicName)
		ctx, _, topic := initClient()
		msg := &ps.Message{Data: []byte(message)}
		r := topic.Publish(ctx, msg)
		id, err := r.Get(ctx)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Sent message with id " + id)
		}
	},
}

func init() {
	RootCmd.AddCommand(pubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	pubCmd.Flags().StringVarP(&message, "messageData", "m", "", "Message to be sent")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pubCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
