# gcloudps

## Google Cloud Pub/Sub CLI Utility.

Subscribes and counts the number of messages.

Install from source:

`go get install github.com/erikwilliamsa/gcloudps`


Usage:

* Log into your Gooogle Cloud project using the gcloud command 
* `gcloudps sub -p <gcp project name> -t <topicname>  -s <subscriptionname>` 

**Note:** Current version Auto deletes subscriptions and Auto-acks messages
Google Cloud Pub/Sub CLI Utility
