#!/bin/sh

set -euf
# shellcheck disable=SC3040
(set -o pipefail 2> /dev/null) && set -o pipefail


sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

ARCH="$(dpkg --print-architecture)"
VERSION_CODENAME="$(source /etc/os-release && echo "${VERSION_CODENAME}")"
readonly ARCH VERSION_CODENAME

printf 'deb [arch=%s signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu %s stable' "${ARCH}" "${VERSION_CODENAME}" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt-get update
sudo apt-get install -y docker-compose-plugin && clear && echo "Setup complete. You may now begin the tutorial."