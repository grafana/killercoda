# Download and configure Pyroscope

1. Download Pyroscope.

   You can use Docker or download a binary to install Pyroscope.

   - To install with Docker, run the following command:

     ```bash
     docker pull grafana/pyroscope:latest
     ```{{exec}}

   - To use a local binary:

     Download the appropriate [release asset](https://github.com/grafana/pyroscope/releases/latest) for your operating system and architecture and make it executable.

     For example, for Linux with the AMD64 architecture:

     ```bash
     # Download Pyroscope v1.0.0 and unpack it to the current folder
     curl -fL https://github.com/grafana/pyroscope/releases/download/v1.0.0/pyroscope_1.0.0_linux_amd64.tar.gz | tar xvz
     ```{{exec}}
