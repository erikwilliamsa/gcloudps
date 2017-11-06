# gcloudps

## Google Cloud Pub/Sub CLI Utility.

Subscribes and counts the number of messages.

Install from source:

`go get github.com/erikwilliamsa/gcloudps`


### Usage:

* Log into your Gooogle Cloud project using the gcloud command 

**Subscribing** 

* `gcloudps sub -p <gcp project name> -t <topicname>  -s <subscriptionname>` 

> **Note:** Current version Auto deletes subscriptions and Auto-acks messages
Google Cloud Pub/Sub CLI Utility

**Publishing** 

Publish a single customm message: 

`gcloudps pub -p <gcp project name> -t <topicname>  -m "my message!" ` 

Generate an publish a number of messages: 

`gcloudps pub generate -c 1000 -p <gcp project name> -t <topicname>` 



