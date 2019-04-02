# Data Stream Delivery
This repository contains the needed code to add support for data streams to the data marketplace. The project is written in [Go](https://golang.org/).

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
# Testing data-stream-delivery

Run the tests:
```bash
./scripts/run_glide_up
./scripts/run_units.sh
```
