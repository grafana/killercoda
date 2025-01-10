# Prerequisites

In this section we will cover the prerequisites you need to have in place in order to build and run the Sandbox Transformer.

## Install Go

The Sandbox Transformer is written in Go, so you will need to have Go installed on your machine. You can download Go from the official website [here](https://golang.org/dl/). In this case we will install the Ubuntu package:

1. Download the Go archive:

   ```bash
   wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz
   ```{{exec}}

1. Remove old versions of Go and extract the archive:

   ```bash
   rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz
   ```{{exec}}

1. Add the Go binary install location to your `PATH` environment variable:

   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```{{exec}}

1. Verify the installation:

   ```bash
   go version
   ```{{exec}}

## Clone the repository

First, you need to clone the repository to your local machine. You can do this by running the following command:

```bash
git clone https://github.com/grafana/killercoda.git && cd killercoda
```{{exec}}

It's best practise to create a new branch for each new course you create. You can do this by running the following command:

```bash
git checkout -b my-new-course
```{{exec}}
