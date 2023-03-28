// Copyright © 2017 NAME HERE Erik Williams erikwilliamsa@gmail.com
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
	"os/signal"
	"strings"

	"github.com/erikwilliamsa/gcloudps/formatters"

	"cloud.google.com/go/pubsub"
	ps "github.com/erikwilliamsa/gcloudps/pubsub"
	"github.com/erikwilliamsa/gcloudps/workers"
	"github.com/spf13/cobra"
)

var (
	subName       string
	dontdeletesub = false
	preview       = false
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Subscribe to a given subscription/topic",
	Args: func(cmd *cobra.Command, args []string) error {
		flags := make(map[string]string)
		flags["subname"] = subName
		return CheckRequiredFlags(flags)
	},
	Long: `
	Creates a subscriber that will continue to recieve messages while running.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Using project " + ProjectName)

		ctx, client, topic := initClient()

		fmt.Println("Creating subscription " + subName)
		subscription, err := client.CreateSubscription(ctx, subName,
			pubsub.SubscriptionConfig{Topic: topic})
		if err != nil {
			if strings.Contains(err.Error(), "Resource not found") {
				log.Fatal(err.Error())
			} else {
				fmt.Println(subName + " already created" + err.Error())
			}

		}

		cleanup(ctx, subscription)
		mch := workers.NewCountMessageHandler()

		if preview {
			mch.Preview = true
			mch.Formatter = &formatters.JSONFormatter{}
		}
		sc := ps.NewSubscriberClient(ctx, subscription, mch)

		workers.Subscribe(ctx, sc)

	},
}

func cleanup(ctx context.Context, s *pubsub.Subscription) {
	//Want to figure out how to do this
	// in its own method to be used elsewhere
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	fmt.Println("Press Ctrl+c to exit")
	go func() {
		for sig := range c {
			if sig != nil {
				fmt.Println("\nExiting")
				fmt.Println("Deleting the subscribtion")
				if !dontdeletesub {
					deleteSub(ctx, s)
					os.Exit(0)
				} else {

					fmt.Println("Not deleting the subscription. The subscription will continue to receive messages.")

					os.Exit(0)
				}

			}

		}
	}()

}

func deleteSub(ctx context.Context, s *pubsub.Subscription) {
	err := s.Delete(ctx)
	if err != nil {
		fmt.Println("Subscribtion was not deleted: " + err.Error())
	} else {
		fmt.Println("Subscribtion deleted")
	}
}
func init() {
	RootCmd.AddCommand(subCmd)
	subCmd.Flags().StringVarP(&subName, "subname", "s", "", "Name of the subscription to use")
	subCmd.Flags().BoolVar(&dontdeletesub, "no-delete", false, "Prevent deleting the subcription on exit.")
	subCmd.Flags().BoolVar(&preview, "preview", false, "Preview deleting the subcription on exit.")

}
