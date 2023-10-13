#!/bin/bash

docker kill $(docker ps -q)

#Remove containers
docker rm  $(docker ps -aq)

#Remove images
docker rmi -f $(docker images | grep "dev\|none\|test-vp\|peer[0-9]" | awk '{print $3}')

#Prune volumes
echo y | docker volume prune

#Prune network
echo y | docker network prune

#Remove certificates;
# rm -rf crypto-config/;

#Remove channel artifacts;
# rm -rf channel-artifacts/*;
