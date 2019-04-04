# Data Stream Delivery
This repository contains the needed code to add support for data streams to the data marketplace. The project is written in [Go](https://golang.org/).
To run this component correctly, you should be familiar with the [Data marketplace](https://github.com/lgsvl/data-marketplace) components because there is a particular dependency between the components.


# Build prerequisites
  * Install [golang](https://golang.org/).
  * Install [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).
  * Configure go. GOPATH environment variable must be set correctly before starting the build process.

### Download and build source code

```bash
mkdir -p $HOME/workspace
export GOPATH=$HOME/workspace
mkdir -p $GOPATH/src/github.com/lgsvl
cd $GOPATH/src/github.com/lgsvl
git clone git@github.com:lgsvl/data-marketplace-stream-delivery.git
cd data-marketplace-stream-delivery
./scripts/build
```

### Docker Build

You can use the dockerfile to build a docker image:
```
docker build -t data-stream-delivery .
docker run -p 7778:7778 data-stream-delivery
```

### Kubernetes Deployment 
The [deployment](./deployment) folder contains the deployment and persistent volume claim manifests to deploy this component.
We assume that [Data marketplace Chaincode REST](https://github.com/lgsvl/data-marketplace-chaincode-rest) is already deployed.

To create a stream, you can do:
```
curl -X POST  "https://<address>:7778/v1/pullpushstream" -H "Content-Type: application/json" -H "authorization: user eyJraWQiOiJjNmdCQWhydDBPMmplOTI2RWVqaFwvaHdVXC9ha2dhc2JOT3puVnR0OXdsc0k9IiwiYWxnIjoiUlMyNTYifQ.eyJhdF9oYXNoIjoiWV9menhtX3EtQTBUU1pSNzd4QXU3USIsInN1YiI6ImI1MGE0YjYyLTRmMmQtNDI3NC1iNzljLTdhMzA4MmEwMTllOSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtd2VzdC0yLmFtYXpvbmF3cy5jb21cL3VzLXdlc3QtMl9zdUdKeWNaNXciLCJjb2duaXRvOnVzZXJuYW1lIjoiamltIiwibm9uY2UiOiJmb29iYXJiYXoiLCJhdWQiOiI3Z21ucXAyNzIzNGFha25xdDRkMmd0MWI1ciIsImV2ZW50X2lkIjoiMDMyYjU2NjItY2NjNy0xMWU4LTljYzQtOTc1OTIxN2EwNTlkIiwidG9rZW5fdXNlIjoiaWQiLCJhdXRoX3RpbWUiOjE1MzkyMDE1NzYsImV4cCI6MTUzOTIwNTE3NiwiaWF0IjoxNTM5MjAxNTc2LCJlbWFpbCI6ImppbUBjb21wYW55Mi5jb20ifQ.ccada-wPb9loOHLuKqnms_hIhoFB-jvD4IcrmT1Y72XjjpT-T_rmSK7ya8ZBK86S5O3GHYo8a6tPPNSoOxjLeFJa_6EW54ZLFUY4mrlqyl1kLOpq5JFNSRUGPith_DpWaM38NKgnmeTBEAhixhAcCtMn0u7LjHJ34zLNrPWk95tcTMRXXo40Pb5uPZENGsouHC_kVxdcbjbSMBrI0GgKRo-WROY1HLsS4fb2MXI4tKUevOFCTn1Rx6Z0Gdz1wA4TeAyRYiXTVg5K6t11IjQ9cq9sRIkAnOzCvyiNKXFQOiPh-Fm8iqQPBkbk5wF3JwHMXmnCA0und-DhF0MPEpg7Qg" --data '{"DataContractTypeID":"123","DataStreamProtocol":"KAFKA","DataStreamSourceURL":"http://stream.meetup.com/2/rsvps","DataStreamDestinationURL":"kafka-svc:9093","DataStreamTopic":"123"}'
```
In the normal scenario these data should be passed over from the portal of the data marketplace.
This request assumes that you have a running kafka on `kafka-svc:9093` and that you already created a `DataContractType`. It assumes also that you have a valid token.

# Testing data-stream-delivery

Run the tests:
```bash
./scripts/run_glide_up
./scripts/run_units.sh
```

# Queries Description

More details about the queries and payloads are [here](./Queries.md).
