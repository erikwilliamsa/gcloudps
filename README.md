# gcloudps

## Google Cloud Pub/Sub CLI Utility.

Publish and subscribe to Google Cloud Pub/Sub from the CLI.

Install from source:

`go get github.com/erikwilliamsa/gcloudps`


### Usage:

* Log into your Gooogle Cloud project using the gcloud command 

**Subscribing** 

* `gcloudps sub -p <gcp project name> -t <topicname>  -s <subscriptionname>` 

> **Note:** If you do not want to delete your subscritpion, use the --no-delete flag.

**Publishing** 

Publish a single customm message: 

`gcloudps pub -p <gcp project name> -t <topicname>  -m "my message!" ` 

Generate an publish a number of messages: 

`gcloudps pub generate -c 1000 -p <gcp project name> -t <topicname>` 



