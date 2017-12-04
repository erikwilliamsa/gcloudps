# gcloudps

[![GoDoc](https://godoc.org/github.com/erikwilliamsa/gcloudps?status.svg)](https://godoc.org/github.com/erikwilliamsa/gcloudps)
[![Build](https://travis-ci.org/erikwilliamsa/gcloudps.svg?branch=master)](https://travis-ci.org/erikwilliamsa/gcloudps)

## Google Cloud Pub/Sub CLI Utility.

Publish and subscribe to Google Cloud Pub/Sub from the CLI.

Install from source:

`go get github.com/erikwilliamsa/gcloudps`


### Usage:

* Log into your Gooogle Cloud project using the gcloud command 

**Subscribing** 

* `gcloudps sub -p <gcp project name> -t <topicname>  -s <subscriptionname>` 

Other flags: 

| Flag | Description|
-------|------------|
| --no-delete | Disables the auto delete of subscriptions on exit |
| --preview | Displays each messages data (body).  Will auto format JSON payloads | 



**Publishing** 

Publish a single customm message: 

`gcloudps pub -p <gcp project name> -t <topicname>  -m "my message!" ` 

Generate an publish a number of messages: 

`gcloudps pub generate -c 1000 -p <gcp project name> -t <topicname>` 
