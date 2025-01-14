# Prerequisites

In this section we will cover the prerequisites you need to have in place in order to download and run the Sandbox Transformer.

## Download the transformer

The Sandbox Transformer is written in Go and is distributed as a binary. You may also build the transformer from source if you prefer.

1. Download the Transformer binary from the [releases page](https://github.com/grafana/killercoda/releases):

   ```bash
   wget https://github.com/grafana/killercoda/releases/download/v0.1.5/transformer-linux-amd64 -O transformer
   ```{{exec}}

1. Make the binary executable:

   ```bash
   chmod +x transformer
   ```{{exec}}

## Clone the repository

You will also need to clone the repository to your local machine. You can do this by running the following command:

```bash
git clone https://github.com/grafana/killercoda.git && cd killercoda
```{{exec}}

Its best practise to create a new branch for each new course you create. You can do this by running the following command:

```bash
git checkout -b my-new-course
```{{exec}}
