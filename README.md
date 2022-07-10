# cyderes project

### What is this?
This repo is my attempt at completing my assignment for (hopefully) working with Cyderes.

### How to use
This site has one (1) API endpoint `api/v1/query` which takes a query param `req` which expects an ipv4 address, domain name, or SHA256 hash. 

`api/v1/query?req=google.com` or `api/v1/query?req=142.250.65.174` or `api/v1/query?req=cf4b367e49bf0b22041c6f065f4aa19f3cfe39c8d5abc0617343d1a66c6a26f5`

This will return virus total data as well as some net/http related data.

The URL has been provided via email but if anyone were to stumble on this and want to try it the terraform should do the trick for you.


### Interesting things learned
I've never personally written code to be handled for lambda, so writing my code to work locally and then having to upend (almost) the entire thing to work with lambda was a little stressful, because of the deadline, but worth learning.

Apparently when you develop on Linux (arch btw), for your binary to be deployed in Lambda correctly you have to disable CGO when building. This was interesting because they mention it for [windows](https://github.com/aws/aws-lambda-go#for-developers-on-windows) but not for Linux, except in [this](https://github.com/aws/aws-lambda-go/issues/340#issue-748665352) issue.

I thought about running the inputs through Shodan but they charge $45 just for api access which was news to me.

Never used Github actions before, but it was the same as almost every other CI provider ever.