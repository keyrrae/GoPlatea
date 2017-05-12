#!/bin/sh

###########################
# Docker SETUP
###########################
apt-get update
apt-get install -y \
    linux-image-extra-$(uname -r) \
    linux-image-extra-virtual
apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
apt-get install -y docker-ce

add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

#ln -sf /usr/bin/docker.io /usr/local/bin/docker
#sed -i '$acomplete -F _docker docker' /etc/bash_completion.d/docker.io

echo "Docker Setup complete"

###########################
# Golang setup
###########################
apt-get update
apt-get install -y golang-go
echo "Golang setup Complete"

###########################
# Start Docker
###########################
chmod 777 UpdateDocker.sh

service docker.io restart
./updatedocker.sh
