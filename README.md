# cyderes project

### What is this?
This repo is my attempt at completing my assignment for (hopefully) working with Cyderes.

### How to use
This site has one (1) API endpoint `api/v1/query` which takes a query param `req` which expects an ipv4 address, domain name, or SHA256 hash. 

`api/v1/query?req=google.com` or `api/v1/query?req=142.250.65.174` or `api/v1/query?req=cf4b367e49bf0b22041c6f065f4aa19f3cfe39c8d5abc0617343d1a66c6a26f5`

This will return virus total data regarding the input's reputation and last submitted date.